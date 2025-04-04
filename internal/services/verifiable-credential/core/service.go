package core

import (
	"context"
	"crypto/ed25519"
	"encoding/hex"
	"encoding/json"
	"time"

	"github.com/dawex/vc-generator/internal/common/config"
	"github.com/dawex/vc-generator/internal/common/db/models"
	"github.com/dawex/vc-generator/internal/common/entities"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"

	jsonld "github.com/trustbloc/did-go/doc/ld/processor"
	utiltime "github.com/trustbloc/did-go/doc/util/time"

	"github.com/trustbloc/kms-go/spi/kms"
	"github.com/trustbloc/vc-go/proof/creator"
	"github.com/trustbloc/vc-go/proof/jwtproofs/eddsa"
	"github.com/trustbloc/vc-go/proof/ldproofs/ed25519signature2018"
	"github.com/trustbloc/vc-go/proof/ldproofs/ed25519signature2020"
	"github.com/trustbloc/vc-go/proof/ldproofs/jsonwebsignature2020"
	"github.com/trustbloc/vc-go/verifiable"

	compliancelogs_ports "github.com/dawex/vc-generator/internal/services/compliance-logs/ports"
	negotiationcontracts_ports "github.com/dawex/vc-generator/internal/services/negotiation-contracts/ports"
	verifiable_credential_ports "github.com/dawex/vc-generator/internal/services/verifiable-credential/ports"
)

type Service struct {
	negotiationcontracts_repository  negotiationcontracts_ports.Repository
	compliancelogs_repository        compliancelogs_ports.Repository
	verifiable_credential_repository verifiable_credential_ports.Repository
	app_config                       config.Config
	publicKey                        ed25519.PublicKey
	privateKey                       ed25519.PrivateKey
	signer                           *creator.ProofCreator
}

func New(app_config config.Config, verifiable_credential_repository verifiable_credential_ports.Repository, compliancelogs_repository compliancelogs_ports.Repository, negotiationcontracts_repository negotiationcontracts_ports.Repository) *Service {
	// Generate an Ed25519 key pair for the issuer from configured Seed
	privateKey := ed25519.NewKeyFromSeed([]byte(app_config.Security.Seed))
	publicKey := privateKey.Public().(ed25519.PublicKey)

	log.Info().Msgf("publicKey (hex encoded) : %v", hex.EncodeToString(publicKey))
	log.Info().Msgf("privateKey (hex encoded) : %v", hex.EncodeToString(privateKey))

	// Init Signer
	signer := creator.New(
		creator.WithLDProofType(jsonwebsignature2020.New(), entities.NewEd25519Signer(privateKey)),
		creator.WithLDProofType(ed25519signature2018.New(), entities.NewEd25519Signer(privateKey)),
		creator.WithLDProofType(ed25519signature2020.New(), entities.NewEd25519Signer(privateKey)),
		creator.WithJWTAlg(eddsa.New(), entities.NewEd25519Signer(privateKey)))

	return &Service{
		negotiationcontracts_repository:  negotiationcontracts_repository,
		compliancelogs_repository:        compliancelogs_repository,
		verifiable_credential_repository: verifiable_credential_repository,
		app_config:                       app_config,
		publicKey:                        publicKey,
		privateKey:                       privateKey,
		signer:                           signer,
	}
}

func (s *Service) GetPublicKey(ctx context.Context) (*ed25519.PublicKey, error) {
	return &s.publicKey, nil
}

func (s *Service) ListVerifiableCredentials(ctx context.Context) ([]models.VerifiableCredential, error) {
	return s.verifiable_credential_repository.ListVerifiableCredentials(ctx)
}

func (s *Service) SignVerifiableCredential(ctx context.Context, contractId string, executionId string) (*models.VerifiableCredential, error) {
	// Fetch saved Contract linked with contractId
	negotiationContract, err := s.negotiationcontracts_repository.GetNegotiationContract(ctx, contractId)
	if err != nil {
		return nil, err
	}

	// Fetch all ComplianceLogs linked with contractId and executionId
	compliancelogs, err := s.compliancelogs_repository.ListComplianceLogs(ctx, contractId, executionId)
	if err != nil {
		return nil, err
	}

	if len(compliancelogs) == 0 {
		return nil, errors.New("no compliancelogs registered")
	}

	// Construct credential subject
	subjects := []verifiable.Subject{}
	var complianceAudit []map[string]interface{}
	for _, compliancelog := range compliancelogs {

		var complianceLogsValues []compliancelogs_ports.ComplianceAuditLog

		complianceLogsValue, err := compliancelog.ComplianceLogs.Value()
		if err != nil {
			return nil, err
		}
		if err := json.Unmarshal(complianceLogsValue.([]byte), &complianceLogsValues); err != nil {
			return nil, err
		}

		complianceAudit = append(complianceAudit, map[string]interface{}{
			"monitoringEvent": map[string]interface{}{
				"source":    compliancelog.Source,
				"timestamp": compliancelog.Timestamp,
				"metric":    compliancelog.Metric,
				"value":     compliancelog.Value,
				"log":       compliancelog.Log,
			},
			"complianceLogs": complianceLogsValues,
		})
	}

	negotiationContractEntity, _ := modelToEntity(negotiationContract)
	subjects = append(subjects, verifiable.Subject{
		ID: compliancelogs[0].ContractID,
		CustomFields: verifiable.CustomFields{
			"executionId":     compliancelogs[0].ExecutionID,
			"contract":        negotiationContractEntity,
			"complianceAudit": complianceAudit,
		},
	})

	// Init Credential Content
	issued := utiltime.NewTime(time.Now().UTC())
	credentialContent := verifiable.CredentialContents{
		Context: []string{"https://www.w3.org/2018/credentials/v1"},
		Types:   []string{"VerifiableCredential"},
		Subject: subjects,
		Issuer: &verifiable.Issuer{
			ID: s.app_config.Issuer.ID,
			CustomFields: verifiable.CustomFields{
				"name": s.app_config.Issuer.Name,
			},
		},
		Issued: issued,
	}

	// Create the VerifiableCredential from CredentialContent
	verifiableCredential, err := verifiable.CreateCredential(credentialContent, nil)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	// Sign VerifiableCredential
	err = verifiableCredential.AddLinkedDataProof(&verifiable.LinkedDataProofContext{
		Created:                 nil,
		SignatureType:           "Ed25519Signature2018",
		KeyType:                 kms.ED25519Type,
		ProofCreator:            s.signer,
		SignatureRepresentation: verifiable.SignatureJWS,
		VerificationMethod:      "assertionMethod",
	}, jsonld.WithDocumentLoader(getJSONLDDocumentLoader()))
	if err != nil {
		return nil, errors.WithStack(err)
	}

	// Map verifiableCredential to model
	modelToSave, err := mapVerifiableCredentialToModel(verifiableCredential)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return s.verifiable_credential_repository.UpsertVerifiableCredential(ctx, modelToSave)
}
