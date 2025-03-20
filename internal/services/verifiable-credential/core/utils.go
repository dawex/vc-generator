package core

import (
	"encoding/json"
	"time"

	"github.com/dawex/vc-generator/internal/common/db/models"
	"github.com/dawex/vc-generator/internal/common/utils"
	"github.com/trustbloc/vc-go/verifiable"

	lddocloader "github.com/trustbloc/did-go/doc/ld/documentloader"
	ldtestutil "github.com/trustbloc/did-go/doc/ld/testutil"
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
