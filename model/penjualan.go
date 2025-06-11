package model

// Penjualan represents a sales transaction in the system
type Penjualan struct {
	ID           uint64  `gorm:"primarykey" json:"id"`              // Unique identifier for the sales transaction
	Kode_invoice string  `json:"kode_invoice"`                     // Invoice code for the transaction
	Nama_pembeli string  `json:"nama_pembeli"`                     // Name of the buyer
	Subtotal     float64 `json:"subtotal"`                         // Total amount before discount
	Kode_diskon  string  `json:"kode_diskon"`                      // Discount code applied to the transaction
	Diskon       float64 `json:"diskon"`                           // Discount amount applied
	Total        float64 `json:"total"`                            // Final total amount after discount
	Model        // Embeds common fields like CreatedAt, UpdatedAt, etc.
	Created_by   string  `json:"created_by"`                      // Person who created the transaction record
}

// CreateP represents the data structure used for creating a new sales transaction
type CreateP struct {
	ID             uint64          `gorm:"primarykey" json:"id"`             // Unique identifier for the sales transaction
	Kode_invoice   string          `json:"kode_invoice"`                    // Invoice code for the transaction
	Nama_pembeli   string          `json:"nama_pembeli"`                    // Name of the buyer
	Subtotal       float64         `json:"subtotal"`                        // Total amount before discount
	Kode_diskon    string          `json:"kode_diskon"`                     // Discount code applied to the transaction
	Diskon         float64         `json:"diskon"`                          // Discount amount applied
	Total          float64         `json:"total"`                           // Final total amount after discount
	Created_by     string          `json:"created_by"`                     // Person who created the transaction record
	ItemPenjualan  []ItemPenjualan `json:"item_penjualan"`                  // List of items involved in the transaction
}
