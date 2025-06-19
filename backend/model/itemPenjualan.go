package model

import (
	"time"

	"gorm.io/gorm"
)

// ItemPenjualan represents an item sold in a sale transaction
type ItemPenjualan struct {
	IDBarang    uint           `json:"id_barang"`               // ID of the item being sold
	IDPenjualan uint           `json:"id_penjualan"`            // ID of the sale transaction (foreign key)
	Jumlah      uint           `json:"jumlah"`                  // Quantity of the item sold
	SubTotal    float64        `json:"sub_total"`               // Subtotal amount for this item (quantity * item price)
	CreatedAt   time.Time      `json:"created_at"`              // Timestamp when the record was created
	UpdatedAt   time.Time      `json:"updated_at"`              // Timestamp when the record was last updated
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"deleted_at"` // Timestamp when the record was deleted (soft deletion)
}
