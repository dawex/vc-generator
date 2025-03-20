package ports

import (
	"context"
	"crypto/ed25519"

	"github.com/dawex/vc-generator/internal/common/db/models"
)

type Service interface {
	GetPublicKey(ctx context.Context) (*ed25519.PublicKey, error)

	ListVerifiableCredentials(ctx context.Context) ([]models.VerifiableCredential, error)
	SignVerifiableCredential(ctx context.Context, contractId string, executionId string) (*models.VerifiableCredential, error)
}
