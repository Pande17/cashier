package modelfunc

import (
	"cashier-machine/model" // Import the model package for the Diskon model
	"fmt"

	"github.com/sirupsen/logrus" // Import logrus for logging
	"gorm.io/gorm"               // Import the GORM package for ORM functionality
)

// Define the Diskon struct which embeds the model.Diskon struct
type Diskon struct {
	model.Diskon
}

// CreateDiskon inserts a new Diskon record into the database
func (kd *Diskon) CreateDiskon(db *gorm.DB) error {
	err := db.Model(Diskon{}).Create(&kd).Error // Attempt to create a new record
	if err != nil {                             // Check if there was an error
		return err // Return the error if creation fails
	}
	return nil // Return nil if creation is successful
}

// GetAll retrieves all Diskon records from the database
func (kd *Diskon) GetAll(db *gorm.DB) ([]Diskon, error) {
	res := []Diskon{} // Initialize an empty slice of Diskon
	err := db.
		Model(Diskon{}). // Specify the Diskon model
		Find(&res).      // Query all Diskon records and store them in res
		Error            // Get any error that occurred
	if err != nil { // Check if there was an error
		return []Diskon{}, err // Return an empty slice and the error
	}
	return res, nil // Return the retrieved records and nil (no error)
}

// GetByCode retrieves a single Diskon record by its discount code
func (kd *Diskon) GetByCode(db *gorm.DB, kodediskon string) (Diskon, error) {
	res := Diskon{}                                                 // Initialize an empty Diskon
	logrus.Println("Executing query with kode_diskon:", kodediskon) // Log the discount code being queried

	// Check database connection
	sqlDB, err := db.DB() // Get the underlying SQL database connection
	if err != nil {       // Check if there was an error
		logrus.Println("Error getting database connection:", err) // Log the error
		return Diskon{}, err                                      // Return an empty Diskon and the error
	}

	if err := sqlDB.Ping(); err != nil { // Ping the database to check if the connection is alive
		logrus.Println("Database connection failed:", err) // Log the error
		return Diskon{}, err                               // Return an empty Diskon and the error
	}

	err = db.
		Where("kode_diskon = ?", kodediskon). // Query for a Diskon record with the given discount code
		First(&res).                          // Retrieve the first matching record
		Error                                 // Get any error that occurred
	if err != nil { // Check if there was an error
		logrus.Println("Error executing query:", err) // Log the error
		if err == gorm.ErrRecordNotFound {            // Check if the error is due to no record found
			return Diskon{}, fmt.Errorf("record not found") // Return a not found error
		}
		return Diskon{}, err // Return the error if it's not a record not found error
	}

	logrus.Println("Query successful, found record:", res) // Log the successful query and the found record
	return res, nil                                        // Return the retrieved Diskon and nil (no error)
}

// GetByID retrieves a single Diskon record by its ID
func (kd *Diskon) GetByID(db *gorm.DB) (Diskon, error) {
	res := Diskon{} // Initialize an empty Diskon
	err := db.
		Model(Diskon{}).        // Specify the Diskon model
		Where("id = ?", kd.ID). // Query for a Diskon record with the given ID
		Take(&res).             // Retrieve the record
		Error                   // Get any error that occurred
	if err != nil { // Check if there was an error
		return Diskon{}, err // Return an empty Diskon and the error
	}
	return res, nil // Return the retrieved Diskon and nil (no error)
}

// Delete removes a Diskon record from the database by its ID
func (kd *Diskon) Delete(db *gorm.DB) error {
	err := db.
		Model(&Diskon{}).       // Specify the Diskon model
		Where("id = ?", kd.ID). // Query for the Diskon record with the given ID
		Delete(&kd).            // Delete the record
		Error                   // Get any error that occurred
	if err != nil { // Check if there was an error
		return err // Return the error if deletion fails
	}
	return nil // Return nil if deletion is successful
}
