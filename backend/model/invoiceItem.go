package model

import (
	"gorm.io/gorm"
)

// InvoiceItem represents an item sold in a sale transaction
type InvoiceItem struct {
	ID           string     `gorm:"primarykey" json:"id"` // Unique identifier for the invoice item
	KodeInvoice  string     `json:"kode_invoice"`         // ID of the sale transaction (foreign key)
	KodeBarang   string     `json:"kode_barang"`          // ID of the item being sold
	Quantity     uint       `json:"qty"`                  // Quantity of the item sold
	Harga        float64    `json:"harga"`                // Price of the item
	DiskonBarang float64    `json:"diskon_barang"`        // Discount applied to the item
	ItemTotal    float64    `json:"item_total"`           // Subtotal amount for this item (quantity * item price)
	Model        gorm.Model // Embedded Model struct for timestamps and soft deletion
}
