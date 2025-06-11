package model


// Histori represents a historical record of changes to an item, such as stock movements
type Histori struct {
	ID         uint   `gorm:"primarykey" json:"id"`     // Unique identifier for the history record
	ID_barang  uint   `json:"id_barang"`                // ID of the associated item (foreign key)
	Amount     int    `json:"amount"`                   // The amount of the change (positive or negative)
	Status     string `json:"status"`                   // Status of the record (e.g., "added", "removed")
	Keterangan string `json:"keterangan"`              // Additional notes or description
	Model                                 // Embedded Model struct for timestamps and soft deletion
}

// HistoriASK represents a historical record with no timestamps or soft deletion
type HistoriASK struct {
	Amount     int    `json:"amount"`     // The amount of the change
	Status     string `json:"status"`     // Status of the record
	Keterangan string `json:"keterangan"` // Additional notes or description
}

// HistoriASKM represents a historical record with timestamps and soft deletion, similar to Histori
type HistoriASKM struct {
	Amount     int    `json:"amount"`     // The amount of the change
	Status     string `json:"status"`     // Status of the record
	Keterangan string `json:"keterangan"` // Additional notes or description
	Model             // Embedded Model struct for timestamps and soft deletion
}
