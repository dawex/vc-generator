package core

import (
	"context"
	"crypto/ed25519"
	"encoding/hex"
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

	event_ports "github.com/dawex/vc-generator/internal/services/event/ports"
	verifiable_credential_ports "github.com/dawex/vc-generator/internal/services/verifiable-credential/ports"
)

type Service struct {
	event_repository                 event_ports.Repository
	verifiable_credential_repository verifiable_credential_ports.Repository
	app_config                       config.Config
	publicKey                        ed25519.PublicKey
	privateKey                       ed25519.PrivateKey
	signer                           *creator.ProofCreator
}

func New(app_config config.Config, verifiable_credential_repository verifiable_credential_ports.Repository, event_repository event_ports.Repository) *Service {
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
		event_repository:                 event_repository,
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
	// Fetch all Events linked with contractId and executionId
	events, err := s.event_repository.ListEvents(ctx, contractId, executionId)
	if err != nil {
		return nil, err
	}

	if len(events) == 0 {
		return nil, errors.New("no events registered")
	}

	// Map []event to verifiable.Subject
	subjects := []verifiable.Subject{}
	var monitoringEvents []map[string]interface{}
	for _, event := range events {
		monitoringEvents = append(monitoringEvents, map[string]interface{}{
			"source":    event.Source,
			"timestamp": event.Timestamp,
			"metric":    event.Metric,
			"value":     event.Value,
			"log":       event.Log,
		})
	}
	subjects = append(subjects, verifiable.Subject{
		ID: events[0].ContractID,
		CustomFields: verifiable.CustomFields{
			"executionId":      events[0].ExecutionID,
			"monitoringEvents": monitoringEvents,
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
