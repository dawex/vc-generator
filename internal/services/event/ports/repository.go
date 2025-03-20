package ports

import (
	"context"

	"github.com/dawex/vc-generator/internal/common/db/models"
)

type Repository interface {
	ListEvents(ctx context.Context, contractId string, executionId string) ([]models.Event, error)
	UpsertEvent(ctx context.Context, model *models.Event) (*models.Event, error)
}
