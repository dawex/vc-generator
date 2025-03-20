package repository

import (
	"context"

	"github.com/dawex/vc-generator/internal/common/db/models"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

type Repository struct {
	db *gorm.DB
}

func New(db *gorm.DB) *Repository {
	return &Repository{
		db: db,
	}
}

func (r *Repository) BeginWithCtx(ctx context.Context) *gorm.DB {
	return r.db.WithContext(ctx)
}

func (r *Repository) UpsertEvent(ctx context.Context, model *models.Event) (*models.Event, error) {
	if err := r.BeginWithCtx(ctx).Save(&model).Error; err != nil {
		return nil, errors.WithStack(err)
	}

	return model, nil
}

func (r *Repository) ListEvents(ctx context.Context, contractId string, executionId string) ([]models.Event, error) {
	elems := []models.Event{}
	if err := r.BeginWithCtx(ctx).Where("contract_id = ?", contractId).Where("execution_id = ?", executionId).Find(&elems).Error; err != nil {
		return nil, errors.WithStack(err)
	}

	return elems, nil
}
