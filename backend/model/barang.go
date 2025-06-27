package model

// Barang represents a product or item in the system
type Barang struct {
	KodeBarang string  `gorm:"primarykey" json:"kode_barang"` // Product code
	Nama       string  `json:"nama"`                          // Product name
	Kategori   string  `json:"kategori"`                      // Category of the product
	HargaJual  float64 `json:"harga_jual"`                    // Selling price of the product
	HargaBeli  float64 `json:"harga_beli"`                    // Cost price of the product
	Stok       uint    `json:"stok"`                          // Stock quantity of the product
	Model    `json:"model"` // Embeds common fields like CreatedAt, UpdatedAt, etc.
}

// Details provides a detailed view of a product, including its history
type Details struct {
	KodeBarang string        `gorm:"primarykey" json:"kode_barang"` // Product code
	Nama       string        `json:"nama"`                          // Product name
	HargaPokok float64       `json:"harga_pokok"`                   // Cost price of the product
	HargaJual  float64       `json:"harga_jual"`                    // Selling price of the product
	TipeBarang string        `json:"tipe_barang"`                   // Type/category of the product
	Stok       uint          `json:"stok"`                          // Stock quantity of the product
	Model                    // Embeds common fields like CreatedAt, UpdatedAt, etc.                     // Person who created the record
	Histori    []HistoriASKM `gorm:"foreignKey:ID_Barang" json:"histori_stok"` // History of stock movements
}

// CreateB represents the data structure used for creating a new product
type CreateB struct {
	ID         uint64       `gorm:"primarykey" json:"id"`                     // Unique identifier for the product
	KodeBarang string       `json:"kode_barang"`                              // Product code
	Nama       string       `json:"nama_barang"`                              // Product name
	HargaPokok float64      `json:"harga_pokok"`                              // Cost price of the product
	HargaJual  float64      `json:"harga_jual"`                               // Selling price of the product
	TipeBarang string       `json:"tipe_barang"`                              // Type/category of the product
	Stok       uint         `json:"stok"`                                     // Stock quantity of the product
	CreatedBy  string       `json:"created_by"`                               // Person who created the record
	Histori    []HistoriASK `gorm:"foreignKey:ID_Barang" json:"histori_stok"` // History of stock movements
}
