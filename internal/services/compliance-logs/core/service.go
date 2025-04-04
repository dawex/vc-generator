package core

import (
	"context"

	"github.com/dawex/vc-generator/internal/common/db/models"
	"github.com/dawex/vc-generator/internal/services/compliance-logs/ports"
)

type Service struct {
	repository ports.Repository
}

func New(repository ports.Repository) *Service {
	return &Service{
		repository: repository,
	}
}

func (s *Service) ListComplianceLogs(ctx context.Context, contractId string, executionId string) ([]models.ComplianceLog, error) {
	return s.repository.ListComplianceLogs(ctx, contractId, executionId)
}

func (s *Service) SaveComplianceLog(ctx context.Context, model *models.ComplianceLog) (*models.ComplianceLog, error) {
	return s.repository.UpsertComplianceLog(ctx, model)
}
