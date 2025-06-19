package modelfunc

import (
	"projek/toko-retail/model"  // Import the model package for the Barang model
	"gorm.io/gorm"              // Import the GORM package for ORM functionality
)

// Define the Barang struct which embeds the model.Barang struct
type Barang struct {
	model.Barang
}

// Create inserts a new Barang record into the database
func (br *Barang) Create(db *gorm.DB) error {
	if err := db.Create(br).Error; err != nil {  // Attempt to create a new record
        return err  // Return the error if creation fails
    }
    return nil  // Return nil if creation is successful
}

// GetAll retrieves all Barang records from the database
func (br *Barang) GetAll(db *gorm.DB) ([]model.Barang, error) {
	res := []model.Barang{}  // Initialize an empty slice of Barang
	err := db.Model(model.Barang{}).Find(&res).Error  // Query all Barang records and store them in res
	if err != nil {  // Check if there was an error
		return []model.Barang{}, err  // Return an empty slice and the error
	}
	return res, err  // Return the retrieved records and error (if any)
}

// GetByID retrieves a single Barang record by its ID
func (br *Barang) GetByID(db *gorm.DB) (model.Barang, error) {
	res := model.Barang{}  // Initialize an empty Barang
	err := db.Model(model.Barang{}).Where("id = ?", br.ID).Find(&res).Error  // Query for a Barang record with the given ID
	if err != nil {  // Check if there was an error
		return res, err  // Return the empty Barang and the error
	}
	return res, nil  // Return the retrieved Barang and nil (no error)
}

// Update modifies an existing Barang record in the database
func (br *Barang) Update(db *gorm.DB) error {
	err := db.Model(model.Barang{}).Where("id = ?", br.ID).Updates(&br).Error  // Update the Barang record with the given ID
	if err != nil {  // Check if there was an error
		return err  // Return the error if the update fails
	}
	return nil  // Return nil if the update is successful
}

// Delete removes a Barang record from the database by its ID
func (br *Barang) Delete(db *gorm.DB) error {
	err := db.Where("id = ?", br.ID).Delete(&model.Barang{}).Error  // Delete the Barang record with the given ID
	if err != nil {  // Check if there was an error
		return err  // Return the error if the deletion fails
	}
	return nil  // Return nil if the deletion is successful
}