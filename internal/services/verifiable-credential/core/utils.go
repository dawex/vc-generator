package core

import (
	"encoding/json"
	"time"

	"github.com/dawex/vc-generator/internal/common/db/models"
	"github.com/dawex/vc-generator/internal/common/utils"
	"github.com/trustbloc/vc-go/verifiable"

	lddocloader "github.com/trustbloc/did-go/doc/ld/documentloader"
	ldtestutil "github.com/trustbloc/did-go/doc/ld/testutil"

	negotiationcontracts_ports "github.com/dawex/vc-generator/internal/services/negotiation-contracts/ports"
)

func getJSONLDDocumentLoader() *lddocloader.DocumentLoader {
	loader, err := ldtestutil.DocumentLoader()
	if err != nil {
		panic(err)
	}

	return loader
}

func mapVerifiableCredentialToModel(vc *verifiable.Credential) (*models.VerifiableCredential, error) {
	model := models.VerifiableCredential{}

	contextB, err := json.Marshal(vc.Contents().Context)
	if err != nil {
		return nil, err
	}
	model.Context = utils.JSON(contextB)

	typesB, err := json.Marshal(vc.Contents().Types)
	if err != nil {
		return nil, err
	}
	model.Type = utils.JSON(typesB)

	subjectB, err := json.Marshal(verifiable.SubjectToJSON(vc.Contents().Subject[0]))
	if err != nil {
		return nil, err
	}
	model.CredentialSubject = utils.JSON(subjectB)

	model.IssuanceDate = vc.Contents().Issued.Time
	model.IssuerID = vc.Contents().Issuer.ID
	model.IssuerName = vc.Contents().Issuer.CustomFields["name"].(string)

	time, err := time.Parse(time.RFC3339, vc.Proofs()[0]["created"].(string))
	if err != nil {
		return nil, err
	}
	model.ProofCreated = time
	model.ProofJws = vc.Proofs()[0]["jws"].(string)
	model.ProofPurpose = vc.Proofs()[0]["proofPurpose"].(string)
	model.ProofType = vc.Proofs()[0]["type"].(string)
	model.ProofVerificationMethod = vc.Proofs()[0]["verificationMethod"].(string)

	return &model, nil
}

func modelToEntity(model *models.NegotiationContract) (*negotiationcontracts_ports.NegotiationContract, error) {
	var entity negotiationcontracts_ports.NegotiationContract

	entity.Id = model.ID
	entity.Type = model.Type
	entity.ConsumerId = model.ConsumerID
	entity.ProducerId = model.ProducerID
	entity.DataProcessingWorkflowObject = model.DataProcessingWorkflowObject
	entity.NaturalLanguageDocument = model.NaturalLanguageDocument
	entity.Title = model.Title
	entity.NegotiationId = model.NegotiationID
	entity.CreatedAt = &model.CreatedAt
	entity.UpdatedAt = &model.UpdatedAt

	policyValue, err := model.OdrlPolicy.Value()
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(policyValue.([]byte), &entity.OdrlPolicy); err != nil {
		return nil, err
	}

	resourceValue, err := model.ResourceDescriptionObject.Value()
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(resourceValue.([]byte), &entity.ResourceDescriptionObject); err != nil {
		return nil, err
	}

	return &entity, nil
}
