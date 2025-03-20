package core

import (
	"context"

	"github.com/dawex/vc-generator/internal/common/db/models"
	"github.com/dawex/vc-generator/internal/services/event/ports"
)

type Service struct {
	repository ports.Repository
}

func New(repository ports.Repository) *Service {
	return &Service{
		repository: repository,
	}
}

func (s *Service) GetEvents(ctx context.Context, contractId string, executionId string) ([]models.Event, error) {
	return s.repository.ListEvents(ctx, contractId, executionId)
}

func (s *Service) SaveEvent(ctx context.Context, model *models.Event) (*models.Event, error) {
	return s.repository.UpsertEvent(ctx, model)
}
