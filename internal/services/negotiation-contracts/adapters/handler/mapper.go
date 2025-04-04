package handler

import (
	"encoding/json"

	"github.com/dawex/vc-generator/internal/common/db/models"
	"github.com/dawex/vc-generator/internal/common/utils"
	"github.com/dawex/vc-generator/internal/services/negotiation-contracts/ports"
)

func modelToEntity(model *models.NegotiationContract) (*ports.NegotiationContract, error) {
	var entity ports.NegotiationContract

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

func entityToModel(entity *ports.NegotiationContract) (*models.NegotiationContract, error) {
	var model models.NegotiationContract

	model.ID = entity.Id
	model.ConsumerID = entity.ConsumerId
	model.CreatedAt = *entity.CreatedAt
	model.UpdatedAt = *entity.UpdatedAt
	model.Title = entity.Title
	model.DataProcessingWorkflowObject = entity.DataProcessingWorkflowObject
	model.NaturalLanguageDocument = entity.NaturalLanguageDocument
	model.NegotiationID = entity.NegotiationId
	model.Type = entity.Type
	model.ProducerID = entity.ProducerId

	policyB, err := json.Marshal(&entity.OdrlPolicy)
	if err != nil {
		return nil, err
	}
	model.OdrlPolicy = utils.JSON(policyB)

	resourceB, err := json.Marshal(&entity.ResourceDescriptionObject)
	if err != nil {
		return nil, err
	}
	model.ResourceDescriptionObject = utils.JSON(resourceB)

	return &model, nil
}
