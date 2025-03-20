package models

import (
	"time"

	"github.com/dawex/vc-generator/internal/common/utils"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type VerifiableCredential struct {
	CreatedAt time.Time `gorm:"not null"`

	IssuerID     string    `gorm:"size:256;not null"`
	IssuerName   string    `gorm:"size:256;not null"`
	IssuanceDate time.Time `gorm:"not null"`

	ProofCreated            time.Time `gorm:"not null"`
	ProofJws                string    `gorm:"size:1024;not null"`
	ProofPurpose            string    `gorm:"size:256;not null"`
	ProofType               string    `gorm:"size:256;not null"`
	ProofVerificationMethod string    `gorm:"size:256;not null"`

	Context           utils.JSON `gorm:"type:jsonb"`
	Type              utils.JSON `gorm:"type:jsonb"`
	CredentialSubject utils.JSON `gorm:"type:jsonb"`

	ID string `gorm:"primary"`
}

func (m *VerifiableCredential) TableName() string {
	return "verifiable_credential"
}

func (m *VerifiableCredential) BeforeCreate(tx *gorm.DB) (err error) {
	// Create VC unique ID
	m.ID = "urn:uuid:" + uuid.NewString()

	return
}
