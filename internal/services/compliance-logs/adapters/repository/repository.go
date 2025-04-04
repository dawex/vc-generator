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

func (r *Repository) UpsertComplianceLog(ctx context.Context, model *models.ComplianceLog) (*models.ComplianceLog, error) {
	if err := r.BeginWithCtx(ctx).Save(&model).Error; err != nil {
		return nil, errors.WithStack(err)
	}

	return model, nil
}

func (r *Repository) ListComplianceLogs(ctx context.Context, contractId string, executionId string) ([]models.ComplianceLog, error) {
	elems := []models.ComplianceLog{}
	if err := r.BeginWithCtx(ctx).Where("contract_id = ?", contractId).Where("execution_id = ?", executionId).Find(&elems).Error; err != nil {
		return nil, errors.WithStack(err)
	}

	return elems, nil
}
