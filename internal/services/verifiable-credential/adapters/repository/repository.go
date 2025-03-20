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

func (r *Repository) UpsertVerifiableCredential(ctx context.Context, model *models.VerifiableCredential) (*models.VerifiableCredential, error) {
	if err := r.BeginWithCtx(ctx).Save(&model).Error; err != nil {
		return nil, errors.WithStack(err)
	}

	return model, nil
}

func (r *Repository) ListVerifiableCredentials(ctx context.Context) ([]models.VerifiableCredential, error) {
	elems := []models.VerifiableCredential{}
	if err := r.BeginWithCtx(ctx).Find(&elems).Error; err != nil {
		return nil, errors.WithStack(err)
	}

	return elems, nil
}
