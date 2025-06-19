package utils

import (
	"cashier-machine/model"
	repository "cashier-machine/repository/config"
	"cashier-machine/repository/modelfunc"
	"time"

	"github.com/siruspen/logrus"
)

// Function to create a new discount code
func CreateKodeDiskon(data model.Diskon) (model.Diskon, error) {
	repoDiskon := modelfunc.Diskon{
		Diskon: data, // Initialize repoDiskon with the provided discount data
	}

	repoDiskon.CreatedAt = time.Now() // Set the creation timestamp
	repoDiskon.UpdatedAt = time.Now() // Set the update timestamp

	err := repoDiskon.CreateDiskon(repository.Mysql.DB) // Save the discount record to the database
	if err != nil {
		return model.Diskon{}, err // Return an error if saving fails
	}

	return repoDiskon.Diskon, nil // Return the newly created discount record
}

// Function to get all discount codes
func GetDiskon() ([]model.Diskon, error) {
	var diskon modelfunc.Diskon
	repoDiskons, err := diskon.GetAll(repository.Mysql.DB) // Retrieve all discount records
	if err != nil {
		return nil, err // Return an error if retrieval fails
	}

	var result []model.Diskon                // Initialize a slice to store the result
	for _, repoDiskon := range repoDiskons { // Iterate through the retrieved discount records
		result = append(result, repoDiskon.Diskon) // Add each record to the result slice
	}

	return result, nil // Return the slice of discount records
}

// Function to get a discount code by its code
func GetDiskonByCode(kodeDiskon string) (model.Diskon, error) {
	logrus.Println("Searching for discount code:", kodeDiskon) // Log the discount code being searched
	diskon := modelfunc.Diskon{
		Diskon: model.Diskon{
			KodeDiskon: kodeDiskon, // Set the discount code from the parameter
		},
	}

	repoDiskon, err := diskon.GetByCode(repository.Mysql.DB, kodeDiskon) // Retrieve the discount record by its code
	if err != nil {
		logrus.Println("Error in GetByCode:", err) // Log any errors encountered
		return model.Diskon{}, err                 // Return an error if retrieval fails
	}

	logrus.Println("Discount code found:", repoDiskon.Diskon) // Log the found discount code
	return repoDiskon.Diskon, nil                             // Return the retrieved discount record
}

// Function to get a discount code by its ID
func GetDiskonByID(id uint) (model.Diskon, error) {
	diskon := modelfunc.Diskon{
		Diskon: model.Diskon{
			ID: id, // Set the ID from the parameter
		},
	}

	repoDiskon, err := diskon.GetByID(repository.Mysql.DB) // Retrieve the discount record by its ID
	if err != nil {
		return model.Diskon{}, err // Return an error if retrieval fails
	}

	return repoDiskon.Diskon, nil // Return the retrieved discount record
}

// Function to update an existing discount code
func UpdateDiskon(id uint, updatedDiskon model.Diskon) (model.Diskon, error) {
	existingDiskon := modelfunc.Diskon{
		Diskon: model.Diskon{
			ID: id, // Set the ID from the parameter
		},
	}

	if err := repository.Mysql.DB.First(&existingDiskon.Diskon).Error; err != nil {
		return model.Diskon{}, err // Return an error if the record is not found
	}

	// Update the fields of the existing discount record
	existingDiskon.Amount = updatedDiskon.Amount
	existingDiskon.Type = updatedDiskon.Type
	existingDiskon.UpdatedAt = time.Now() // Update the timestamp

	if err := repository.Mysql.DB.Save(&existingDiskon.Diskon).Error; err != nil {
		return model.Diskon{}, err // Return an error if saving fails
	}

	return existingDiskon.Diskon, nil // Return the updated discount record
}

// Function to delete a discount code by its ID
func DeleteKode(id uint64) error {
	diskon := modelfunc.Diskon{
		Diskon: model.Diskon{
			ID: uint(id), // Set the ID from the parameter
		},
	}

	return diskon.Delete(repository.Mysql.DB) // Delete the discount record from the database
}
