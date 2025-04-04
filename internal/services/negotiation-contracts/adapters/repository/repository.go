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

func (r *Repository) UpsertNegotiationContract(ctx context.Context, model *models.NegotiationContract) (*models.NegotiationContract, error) {
	if err := r.BeginWithCtx(ctx).Save(&model).Error; err != nil {
		return nil, errors.WithStack(err)
	}

	return model, nil
}

func (r *Repository) ListNegotiationContracts(ctx context.Context) ([]models.NegotiationContract, error) {
	elems := []models.NegotiationContract{}
	if err := r.BeginWithCtx(ctx).Find(&elems).Error; err != nil {
		return nil, errors.WithStack(err)
	}

	return elems, nil
}

func (r *Repository) GetNegotiationContract(ctx context.Context, id string) (*models.NegotiationContract, error) {
	model := &models.NegotiationContract{}
	model.ID = id
	if err := r.BeginWithCtx(ctx).First(model).Error; err != nil {
		return nil, err
	}

	return model, nil
}
