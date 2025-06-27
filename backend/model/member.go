package model

type Member struct {
	ID         string         `gorm:"primarykey" json:"id"` // Unique identifier for the member
	Nama       string         `json:"nama"`                 // Name of the member
	Pic        string         `json:"pic"`                  // Person in charge (PIC) of the member
	Perusahaan string         `json:"perusahaan"`           // Company name of the member
	Kategori   string         `json:"kategori"`             // Category of the member (e.g., gold, silver, etc.)
	Alamat     string         `json:"alamat"`               // Address of the member
	NoTelp     string         `json:"no_telp"`              // Phone number of the member
	Status     string         `json:"status"`               // Status of the member (e.g., active, inactive)
	Model      `json:"model"` // Embeds common fields like CreatedAt, UpdatedAt, etc.
}
