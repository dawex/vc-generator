package handler

import (
	"encoding/json"

	"github.com/dawex/vc-generator/internal/common/db/models"
	"github.com/dawex/vc-generator/internal/services/verifiable-credential/ports"
	"github.com/pkg/errors"
)

func modelToEntity(model *models.VerifiableCredential) (*ports.VcSigned, error) {
	var entity ports.VcSigned

	contextB, err := model.Context.Value()
	if err != nil {
		return nil, errors.WithStack(err)
	}
	var context []string
	if err := json.Unmarshal(contextB.([]byte), &context); err != nil {
		return nil, errors.WithStack(err)
	}
	entity.Context = context

	typesB, err := model.Type.Value()
	if err != nil {
		return nil, errors.WithStack(err)
	}
	var types []ports.Type
	if err := json.Unmarshal(typesB.([]byte), &types); err != nil {
		return nil, errors.WithStack(err)
	}
	entity.Type = types

	subjectB, err := model.CredentialSubject.Value()
	if err != nil {
		return nil, errors.WithStack(err)
	}
	var subject ports.CredentialSubject
	if err := json.Unmarshal(subjectB.([]byte), &subject); err != nil {
		return nil, errors.WithStack(err)
	}
	entity.CredentialSubject = subject

	entity.Id = model.ID
	entity.IssuanceDate = model.IssuanceDate
	entity.Issuer = ports.Issuer{
		Id:   model.IssuerID,
		Name: model.IssuerName,
	}
	entity.Proof = ports.Proof{
		Created:            model.ProofCreated,
		Jws:                model.ProofJws,
		ProofPurpose:       ports.ProofProofPurpose(model.ProofPurpose),
		Type:               ports.ProofType(model.ProofType),
		VerificationMethod: model.ProofVerificationMethod,
	}

	return &entity, nil
}
