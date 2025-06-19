package modelfunc

import (
	"cashier-machine/model" // Import the model package for the Histori model

	"gorm.io/gorm" // Import the GORM package for ORM functionality
)

// Define the Histori struct which embeds the model.Histori struct
type Histori struct {
	model.Histori
}

// Create inserts a new Histori record into the database
func (hs *Histori) Create(db *gorm.DB) error {
	err := db.Model(Histori{}).Create(&hs).Error // Attempt to create a new record
	if err != nil {                              // Check if there was an error
		return err // Return the error if creation fails
	}
	return nil // Return nil if creation is successful
}

// GetIDBarang retrieves all Histori records with a specific id_barang
func (hs *Histori) GetIDBarang(db *gorm.DB) ([]Histori, error) {
	res := []Histori{}                                                               // Initialize an empty slice of Histori
	err := db.Model(Histori{}).Where("id_barang = ?", hs.ID_barang).Find(&res).Error // Query for Histori records with the given id_barang
	if err != nil {                                                                  // Check if there was an error
		return []Histori{}, err // Return an empty slice and the error
	}
	return res, err // Return the retrieved records and the error (if any)
}
