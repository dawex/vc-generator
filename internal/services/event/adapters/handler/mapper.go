package handler

import (
	"github.com/dawex/vc-generator/internal/common/db/models"
	"github.com/dawex/vc-generator/internal/services/event/ports"
)

func modelToEntity(model *models.Event) *ports.Event {
	var entity ports.Event

	id := (ports.Id)(model.ID)
	entity.Id = &id
	entity.CreatedAt = &model.CreatedAt
	entity.ContractId = model.ContractID
	entity.ExecutionId = model.ExecutionID
	entity.MonitoringEvent.Log = model.Log
	entity.MonitoringEvent.Metric = ports.MonitoringEventMetric(model.Metric)
	entity.MonitoringEvent.Source = model.Source
	entity.MonitoringEvent.Value = model.Value
	entity.MonitoringEvent.Timestamp = model.Timestamp

	return &entity
}

func entityToModel(entity *ports.EventIn) *models.Event {
	var model models.Event

	model.ContractID = entity.ContractId
	model.ExecutionID = entity.ExecutionId
	model.Log = entity.MonitoringEvent.Log
	model.Metric = string(entity.MonitoringEvent.Metric)
	model.Source = entity.MonitoringEvent.Source
	model.Value = entity.MonitoringEvent.Value
	model.Timestamp = entity.MonitoringEvent.Timestamp

	return &model
}
