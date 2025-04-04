package models

import (
	"time"

	"github.com/dawex/vc-generator/internal/common/utils"
	"github.com/google/uuid"
)

type ComplianceLog struct {
	CreatedAt time.Time `gorm:"not null"`

	ContractID     string     `gorm:"size:256;not null"`
	ExecutionID    string     `gorm:"size:256;not null"`
	Source         string     `gorm:"size:256;not null"`
	Timestamp      time.Time  `gorm:"not null"`
	Metric         string     `gorm:"size:256;not null"`
	Value          string     `gorm:"size:1024;not null"`
	Log            string     `gorm:"size:1024;not null"`
	Groups         string     `gorm:"size:1024;not null;column:log_group"`
	Result         *string    `gorm:"size:1024"`
	Params         utils.JSON `gorm:"type:jsonb"`
	ComplianceLogs utils.JSON `gorm:"type:jsonb"`

	ID uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4()"`
}

func (m *ComplianceLog) TableName() string {
	return "compliance_log"
}
