package ports

import (
	"context"

	"github.com/dawex/vc-generator/internal/common/db/models"
)

type Service interface {
	GetEvents(ctx context.Context, contractId string, executionId string) ([]models.Event, error)
	SaveEvent(ctx context.Context, model *models.Event) (*models.Event, error)
}
