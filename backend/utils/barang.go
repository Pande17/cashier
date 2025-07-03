package utils

import (
	"cashier-machine/model"
	repository "cashier-machine/repository/config"
	"cashier-machine/repository/modelfunc"
	"errors"
	"fmt"
	"time"

	"gorm.io/gorm"
)

// Function to create a new item
func CreateBarang(data model.Barang) (model.Barang, error) {
	// Initialize repository.Barang with model.Barang data
	repoBarang := modelfunc.Barang{
		Barang: data, // Copy the input data into the repository model
	}

	// Set timestamps and default CreatedBy if not provided
	repoBarang.CreatedAt = time.Now() // Set creation time
	repoBarang.UpdatedAt = time.Now() // Set update time
	if repoBarang.CreatedBy == "" {   // Check if CreatedBy is empty
		repoBarang.CreatedBy = "SYSTEM" // Assign default value if empty
	}

	// Create new item record in the database
	err := repoBarang.Create(repository.Mysql.DB) // Save the new item record
	if err != nil {
		return model.Barang{}, err // Return error if creation fails
	}

	// Update item record with the new KodeBarang
	err = repoBarang.Update(repository.Mysql.DB) // Save the updated record
	if err != nil {
		return model.Barang{}, err // Return error if update fails
	}

	// Fetch history data for the newly created item
	// histori, err := GetASK(repoBarang.ID) // Retrieve historical data for the item
	// if err != nil {
	// 	return model.CreateB{}, err // Return error if fetching history fails
	// }

	// Prepare the CreateB response with updated data and history
	createB := model.Barang{
		KodeBarang: repoBarang.KodeBarang, // Assign KodeBarang to response
		Nama:       repoBarang.Nama,       // Assign Nama to response
		HargaBeli:  repoBarang.HargaBeli,  // Assign HargaPokok to response
		HargaJual:  repoBarang.HargaJual,  // Assign HargaJual to response
		Kategori:   repoBarang.Kategori,   // Assign TipeBarang to response
		Stok:       repoBarang.Stok,       // Assign Stok to response
		Model: model.Model{
			CreatedAt: time.Now().In(time.Local), // Set creation time
			UpdatedAt: time.Now().In(time.Local), // Set update time
			CreatedBy: repoBarang.CreatedBy,      // Assign CreatedAt to response
		},
	}

	return createB, nil // Return the response struct
}

// Function to get a list of all items
func GetBarang() ([]model.Barang, error) {
	var barang modelfunc.Barang               // Initialize repository.Barang
	return barang.GetAll(repository.Mysql.DB) // Retrieve all item records
}

// Function to get item data by its ID
func GetBarangByID(kode string) (*model.Barang, error) {
	barang := modelfunc.Barang{
		Barang: model.Barang{
			KodeBarang: kode, // Set ID for the query
		},
	}
	// Include soft-deleted records in the query
	barangModel, err := barang.GetByID(repository.Mysql.DB.Unscoped()) // Fetch the record including soft-deleted
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) { // Check if the error is record not found
			return nil, err // Return a not found error
		}
		return &model.Barang{}, err // Return other errors
	}

	if barangModel.KodeBarang == "" && kode != "" { // Check if the ID is not created
		return nil, err // Return an error for uncreated ID
	}

	if barangModel.KodeBarang == "" { // Check if ID is zero
		return nil, fmt.Errorf("you can't see this! For More Info: https://s.id/why-i-cant-see-id0") // Return an access restriction error
	}

	// Check if the record is soft-deleted
	if barangModel.DeletedAt.Valid { // Check if the record is marked as deleted
		return nil, fmt.Errorf("this record has been deleted") // Return a deleted record error
	}

	// // Fetch historical data as usual
	// histori, err := GetASKMByIDBarang(barangModel.ID) // Retrieve historical data for the item
	// if err != nil {
	// 	return &model.Details{}, err // Return error if fetching history fails
	// }

	details := model.Barang{
		KodeBarang: barangModel.KodeBarang, // Assign KodeBarang to response
		Nama:       barangModel.Nama,       // Assign Nama to response
		Kategori:   barangModel.Kategori,   // Assign HargaPokok to response
		HargaJual:  barangModel.HargaJual,  // Assign HargaJual to response
		HargaBeli:  barangModel.HargaBeli,  // Assign TipeBarang to response
		Stok:       barangModel.Stok,       // Assign Stok to response
		Model:      barangModel.Model,      // Assign Model to response
		// CreatedBy:  barangModel.CreatedBy,  // Assign CreatedBy to response
		// Histori: histori, // Include historical data in response
	}

	return &details, nil // Return the response struct
}

// Function to update item data
func UpdateBarang(kodeBarang string, barang model.Barang) (model.Barang, error) {
	repoBarang := modelfunc.Barang{
		Barang: barang,
	}

	// Update timestamps
	repoBarang.UpdatedAt = time.Now()
	// Perform the update operation
	err := repoBarang.Update(repository.Mysql.DB)
	if err != nil {
		return model.Barang{}, err
	}
	// Retrieve the updated data to return
	updatedBarang, err := GetBarangByID(kodeBarang)
	if err != nil {
		return model.Barang{}, err
	}
	return *updatedBarang, nil
}

// Function to delete an item
func DeleteBarang(kodeBarang string) error {
	barang := modelfunc.Barang{
		Barang: model.Barang{
			KodeBarang: kodeBarang, // Set ID for deletion
		},
	}
	return barang.Delete(repository.Mysql.DB) // Soft delete the item record
}
