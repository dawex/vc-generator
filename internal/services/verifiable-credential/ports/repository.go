package ports

import (
	"context"

	"github.com/dawex/vc-generator/internal/common/db/models"
)

type Repository interface {
	ListVerifiableCredentials(ctx context.Context) ([]models.VerifiableCredential, error)
	UpsertVerifiableCredential(ctx context.Context, model *models.VerifiableCredential) (*models.VerifiableCredential, error)
}
