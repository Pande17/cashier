package modelfunc

import (
	"projek/toko-retail/model" // Import the model package for the Penjualan model

	"gorm.io/gorm" // Import the GORM package for ORM functionality
)

// Define the Penjualan struct which embeds the model.Penjualan struct
type Penjualan struct {
	model.Penjualan
}

// CreatePenjualan inserts a new Penjualan record into the database
func (pj *Penjualan) CreatePenjualan(db *gorm.DB) error {
	err := db.Create(&pj).Error // Attempt to create a new Penjualan record
	if err != nil {             // Check if there was an error
		return err // Return the error if creation fails
	}
	return nil // Return nil if creation is successful
}

// GetAll retrieves all Penjualan records from the database
func (pj *Penjualan) GetAll(db *gorm.DB) ([]Penjualan, error) {
	res := []Penjualan{}                          // Initialize an empty slice of Penjualan
	err := db.Model(Penjualan{}).Find(&res).Error // Query all Penjualan records and store them in res
	if err != nil {                               // Check if there was an error
		return []Penjualan{}, err // Return an empty slice and the error
	}
	return res, nil // Return the retrieved records and nil (no error)
}

// GetPByID retrieves a single Penjualan record by its ID
func (pj *Penjualan) GetPByID(db *gorm.DB) (Penjualan, error) {
	res := Penjualan{}                                                   // Initialize an empty Penjualan
	err := db.Model(Penjualan{}).Where("id = ?", pj.ID).Take(&res).Error // Query for a Penjualan record with the given ID
	if err != nil {                                                      // Check if there was an error
		return Penjualan{}, err // Return an empty Penjualan and the error
	}
	return res, nil // Return the retrieved Penjualan and nil (no error)
}

// Update modifies an existing Penjualan record in the database
func (pj *Penjualan) Update(db *gorm.DB) error {
	err := db.Model(Penjualan{}).Where("id = ?", pj.ID).Updates(&pj).Error // Update the Penjualan record with the given ID
	if err != nil {                                                        // Check if there was an error
		return err // Return the error if the update fails
	}
	return nil // Return nil if the update is successful
}

// UpdateKodeInvoice updates the invoice code of an existing Penjualan record
func (pj *Penjualan) UpdateKodeInvoice(db *gorm.DB) error {
	err := db.Model(Penjualan{}).Where("id = ?", pj.ID).Updates(&pj).Error // Update the Penjualan record with the given ID (specifically for invoice code)
	if err != nil {                                                        // Check if there was an error
		return err // Return the error if the update fails
	}
	return nil // Return nil if the update is successful
}
