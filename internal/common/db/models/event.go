package models

import (
	"time"

	"github.com/google/uuid"
)

type Event struct {
	CreatedAt time.Time `gorm:"not null"`

	ContractID  string    `gorm:"size:256;not null"`
	ExecutionID string    `gorm:"size:256;not null"`
	Source      string    `gorm:"size:256;not null"`
	Timestamp   time.Time `gorm:"not null"`
	Metric      string    `gorm:"size:256;not null"`
	Value       string    `gorm:"size:1024;not null"`
	Log         string    `gorm:"size:1024;not null"`

	ID uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4()"`
}

func (m *Event) TableName() string {
	return "event"
}
