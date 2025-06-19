package model

// Diskon represents a discount code and its details
type Diskon struct {
	ID         uint    `gorm:"primarykey" json:"id"` 	// Unique identifier for the discount code
	KodeDiskon string  `json:"kode_diskon"`          	// The code for the discount
	Amount     float64 `json:"amount"`               	// The discount amount or percentage
	Type       string  `json:"type"`                 	// The type of discount (e.g., percentage or fixed amount)
	Model              									// Embedded Model struct for timestamps and soft deletion
}
