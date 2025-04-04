package models

import (
	"time"

	"github.com/dawex/vc-generator/internal/common/utils"
)

type NegotiationContract struct {
	UpdatedAt time.Time `gorm:"not null"`
	CreatedAt time.Time `gorm:"not null"`

	Title                        *string `gorm:"size:1024"`
	Type                         string  `gorm:"size:1024;not null"`
	ConsumerID                   string  `gorm:"size:1024;not null"`
	ProducerID                   string  `gorm:"size:1024;not null"`
	DataProcessingWorkflowObject string  `gorm:"size:1024;not null"`
	NaturalLanguageDocument      string  `gorm:"size:1024;not null"`
	NegotiationID                *string `gorm:"size:1024"`

	OdrlPolicy                utils.JSON `gorm:"type:jsonb"`
	ResourceDescriptionObject utils.JSON `gorm:"type:jsonb"`

	ID string `gorm:"size:1024;not null;primary"`
}

func (m *NegotiationContract) TableName() string {
	return "negotiation_contract"
}
