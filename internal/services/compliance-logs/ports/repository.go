package ports

import (
	"context"

	"github.com/dawex/vc-generator/internal/common/db/models"
)

type Repository interface {
	ListComplianceLogs(ctx context.Context, contractId string, executionId string) ([]models.ComplianceLog, error)
	UpsertComplianceLog(ctx context.Context, model *models.ComplianceLog) (*models.ComplianceLog, error)
}
