package model

import (
	"time"

	"gorm.io/gorm"
)

type (
	TimestampedModel struct {
		CreatedAt time.Time
		UpdatedAt time.Time
		DeletedAt gorm.DeletedAt `gorm:"index"`
	}
)
