package model

import (
	"time"

	"gorm.io/gorm"
)

// Model is a base struct that includes common timestamp fields for database records
type Model struct {
	CreatedAt time.Time      `json:"created_at"`              // Timestamp when the record was created
	UpdatedAt time.Time      `json:"updated_at"`              // Timestamp when the record was last updated
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at"` // Timestamp when the record was soft deleted; index is used to speed up queries involving deletion
}
