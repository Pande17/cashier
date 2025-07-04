package model

import (
	"time"
)

// Invoice represents a sales transaction in the system
type Invoice struct {
	KodeInvoice     string     `gorm:"primarykey" json:"kode_invoice"`                      // Invoice code for the transaction
	MemberID        string     `gorm:"foreignKey:member_id;references:id" json:"member_id"` // ID of the buyer
	TanggalBeli     time.Time  `json:"tanggal_beli"`                                        // Date and time of the transaction
	JatuhTempo      *time.Time `json:"jatuh_tempo"`                                         // Due date for payment, can be null
	Status          string     `json:"status"`                                              // Status of the transaction (e.g., "paid", "unpaid")
	Ppn             *float64   `json:"ppn"`                                                 // Value-added tax applied to the transaction, can be null
	BiayaPengiriman *float64   `json:"biaya_pengiriman"`                                    // Shipping cost for the transaction
	Subtotal        float64    `json:"subtotal"`                                            // Total amount before discount
	DiskonTotal     float64    `json:"diskon_total"`                                        // Total discount applied to the transaction
	// Diskon          float64    `json:"diskon"`                                              // Discount amount applied
	Total float64 `json:"total"` // Final total amount after discount
	// Define the relationship with InvoiceItem (One-to-many)
	InvoiceItems []InvoiceItem  `gorm:"foreignKey:KodeInvoice;references:KodeInvoice" json:"invoice_items"` // List of items involved in the transaction
	Model        `json:"model"` // Embeds common fields like CreatedAt, UpdatedAt, etc.
}

// // CreateInv represents the data structure used for creating a new sales transaction
// type CreateInv struct {
// 	ID           uint64        `gorm:"primarykey" json:"id"` // Unique identifier for the sales transaction
// 	Kode_invoice string        `json:"kode_invoice"`         // Invoice code for the transaction
// 	Nama_pembeli string        `json:"nama_pembeli"`         // Name of the buyer
// 	Subtotal     float64       `json:"subtotal"`             // Total amount before discount
// 	Kode_diskon  string        `json:"kode_diskon"`          // Discount code applied to the transaction
// 	Diskon       float64       `json:"diskon"`               // Discount amount applied
// 	Total        float64       `json:"total"`                // Final total amount after discount
// 	InvoiceItem  []InvoiceItem `json:"invoice_item"`         // List of items involved in the transaction
// }
