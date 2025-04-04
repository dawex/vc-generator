package ports

import (
	"context"

	"github.com/dawex/vc-generator/internal/common/db/models"
)

type Service interface {
	ListNegotiationContracts(ctx context.Context) ([]models.NegotiationContract, error)
	SaveNegotiationContract(ctx context.Context, model *models.NegotiationContract) (*models.NegotiationContract, error)
}
