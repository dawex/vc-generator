package handler

import (
	"encoding/json"

	"github.com/dawex/vc-generator/internal/common/db/models"
	"github.com/dawex/vc-generator/internal/common/utils"
	"github.com/dawex/vc-generator/internal/services/compliance-logs/ports"
)

func modelToEntity(model *models.ComplianceLog) (*ports.ComplianceLog, error) {
	var entity ports.ComplianceLog

	id := (ports.Id)(model.ID)
	entity.Id = &id
	entity.CreatedAt = &model.CreatedAt
	entity.ContractId = model.ContractID
	entity.ExecutionId = model.ExecutionID
	entity.MonitoringEvent.Log = model.Log
	entity.MonitoringEvent.Metric = model.Metric
	entity.MonitoringEvent.Source = model.Source
	entity.MonitoringEvent.Value = model.Value
	entity.MonitoringEvent.Timestamp = model.Timestamp
	entity.MonitoringEvent.Groups = model.Groups
	entity.MonitoringEvent.Result = model.Result

	paramsValue, err := model.Params.Value()
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(paramsValue.([]byte), &entity.MonitoringEvent.Params); err != nil {
		return nil, err
	}

	complianceLogsValue, err := model.ComplianceLogs.Value()
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(complianceLogsValue.([]byte), &entity.ComplianceLogs); err != nil {
		return nil, err
	}

	return &entity, nil
}

func entityToModel(entity *ports.ComplianceLogIn) (*models.ComplianceLog, error) {
	var model models.ComplianceLog

	model.ContractID = entity.ContractId
	model.ExecutionID = entity.ExecutionId
	model.Log = entity.MonitoringEvent.Log
	model.Metric = string(entity.MonitoringEvent.Metric)
	model.Source = entity.MonitoringEvent.Source
	model.Value = entity.MonitoringEvent.Value
	model.Timestamp = entity.MonitoringEvent.Timestamp
	model.Result = entity.MonitoringEvent.Result
	model.Groups = entity.MonitoringEvent.Groups

	paramsB, err := json.Marshal(&entity.MonitoringEvent.Params)
	if err != nil {
		return nil, err
	}
	model.Params = utils.JSON(paramsB)

	logsB, err := json.Marshal(entity.ComplianceLogs)
	if err != nil {
		return nil, err
	}
	model.ComplianceLogs = utils.JSON(logsB)

	return &model, nil
}
