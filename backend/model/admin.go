package model

type Admin struct {
	ID       int64  `gorm:"primarykey" json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	Model    Model  // Embeds common fields like CreatedAt, UpdatedAt, etc.
}
