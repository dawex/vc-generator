package ports

import (
	"context"

	"github.com/dawex/vc-generator/internal/common/db/models"
)

type Service interface {
	ListComplianceLogs(ctx context.Context, contractId string, executionId string) ([]models.ComplianceLog, error)
	SaveComplianceLog(ctx context.Context, model *models.ComplianceLog) (*models.ComplianceLog, error)
}
