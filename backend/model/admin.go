package model

import "gorm.io/gorm"

type Admin struct {
	ID       int64      `gorm:"primarykey" json:"id"`
	Username string     `json:"username"`
	Password string     `json:"password"`
	Model    gorm.Model // Embeds common fields like CreatedAt, UpdatedAt, etc.
}
