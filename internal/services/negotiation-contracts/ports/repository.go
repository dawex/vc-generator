package ports

import (
	"context"

	"github.com/dawex/vc-generator/internal/common/db/models"
)

type Repository interface {
	ListNegotiationContracts(ctx context.Context) ([]models.NegotiationContract, error)
	UpsertNegotiationContract(ctx context.Context, model *models.NegotiationContract) (*models.NegotiationContract, error)
	GetNegotiationContract(ctx context.Context, contractID string) (*models.NegotiationContract, error)
}
