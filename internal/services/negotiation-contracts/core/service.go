package core

import (
	"context"

	"github.com/dawex/vc-generator/internal/common/db/models"
	"github.com/dawex/vc-generator/internal/services/negotiation-contracts/ports"
)

type Service struct {
	repository ports.Repository
}

func New(repository ports.Repository) *Service {
	return &Service{
		repository: repository,
	}
}

func (s *Service) ListNegotiationContracts(ctx context.Context) ([]models.NegotiationContract, error) {
	return s.repository.ListNegotiationContracts(ctx)
}

func (s *Service) SaveNegotiationContract(ctx context.Context, model *models.NegotiationContract) (*models.NegotiationContract, error) {
	return s.repository.UpsertNegotiationContract(ctx, model)
}
