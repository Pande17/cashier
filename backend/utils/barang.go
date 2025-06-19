package utils

import (
	"errors"
	"fmt"
	"projek/toko-retail/model"
	repository "projek/toko-retail/repository/config"
	"projek/toko-retail/repository/modelfunc"
	"strconv"
	"time"

	"gorm.io/gorm"
)

// Function to create a new item
func CreateBarang(data model.Barang) (model.CreateB, error) {
	// Initialize repository.Barang with model.Barang data
	repoBarang := modelfunc.Barang{
		Barang: data, // Copy the input data into the repository model
	}

	// Set timestamps and default CreatedBy if not provided
	repoBarang.CreatedAt = time.Now() // Set creation time
	repoBarang.UpdatedAt = time.Now() // Set update time
	if repoBarang.CreatedBy == "" { // Check if CreatedBy is empty
		repoBarang.CreatedBy = "SYSTEM" // Assign default value if empty
	}

	// Create new item record in the database
	err := repoBarang.Create(repository.Mysql.DB) // Save the new item record
	if err != nil {
		return model.CreateB{}, err // Return error if creation fails
	}

	// Set KodeBarang based on TipeBarang and newly created ID
	if repoBarang.TipeBarang == "MAKANAN" { // Check if the item is food
		repoBarang.KodeBarang = fmt.Sprintf("MA-%v", strconv.FormatUint(repoBarang.ID, 10)) // Generate food code
	} else if repoBarang.TipeBarang == "MINUMAN" { // Check if the item is a drink
		repoBarang.KodeBarang = fmt.Sprintf("MI-%v", strconv.FormatUint(repoBarang.ID, 10)) // Generate drink code
	} else { // For other types
		repoBarang.KodeBarang = fmt.Sprintf("L-%v", strconv.FormatUint(repoBarang.ID, 10)) // Generate default code
	}

	// Update item record with the new KodeBarang
	err = repoBarang.Update(repository.Mysql.DB) // Save the updated record
	if err != nil {
		return model.CreateB{}, err // Return error if update fails
	}

	// Fetch history data for the newly created item
	histori, err := GetASK(repoBarang.ID) // Retrieve historical data for the item
	if err != nil {
		return model.CreateB{}, err // Return error if fetching history fails
	}

	// Prepare the CreateB response with updated data and history
	createB := model.CreateB{
		ID:         repoBarang.ID,        // Assign ID to response
		KodeBarang: repoBarang.KodeBarang, // Assign KodeBarang to response
		Nama:       repoBarang.Nama,       // Assign Nama to response
		HargaPokok: repoBarang.HargaPokok, // Assign HargaPokok to response
		HargaJual:  repoBarang.HargaJual,  // Assign HargaJual to response
		TipeBarang: repoBarang.TipeBarang, // Assign TipeBarang to response
		Stok:       repoBarang.Stok,       // Assign Stok to response
		CreatedBy:  repoBarang.CreatedBy,  // Assign CreatedBy to response
		Histori:    histori,                // Include historical data in response
	}

	return createB, nil // Return the response struct
}

// Function to get a list of all items
func GetBarang() ([]model.Barang, error) {
	var barang modelfunc.Barang // Initialize repository.Barang
	return barang.GetAll(repository.Mysql.DB) // Retrieve all item records
}

// Function to get item data by its ID
func GetBarangByID(id uint64) (*model.Details, error) {
	barang := modelfunc.Barang{
		Barang: model.Barang{
			ID: id, // Set ID for the query
		},
	}
	// Include soft-deleted records in the query
	barangModel, err := barang.GetByID(repository.Mysql.DB.Unscoped()) // Fetch the record including soft-deleted
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) { // Check if the error is record not found
			return nil, fmt.Errorf("Record Not Found!") // Return a not found error
		}
		return &model.Details{}, err // Return other errors
	}

	if barangModel.ID == 0 && id != 0 { // Check if the ID is not created
		return nil, fmt.Errorf("ID no: %d not created yet!", id) // Return an error for uncreated ID
	}
	
	if barangModel.ID == 0 { // Check if ID is zero
		return nil, fmt.Errorf("You can't see this! For More Info: https://s.id/why-i-cant-see-id0") // Return an access restriction error
	}

	// Check if the record is soft-deleted
	if barangModel.DeletedAt.Valid { // Check if the record is marked as deleted
		return nil, fmt.Errorf("This record has been deleted") // Return a deleted record error
	}

	// Fetch historical data as usual
	histori, err := GetASKMByIDBarang(barangModel.ID) // Retrieve historical data for the item
	if err != nil {
		return &model.Details{}, err // Return error if fetching history fails
	}

	details := model.Details{
		ID:         barangModel.ID,        // Assign ID to response
		KodeBarang: barangModel.KodeBarang, // Assign KodeBarang to response
		Nama:       barangModel.Nama,       // Assign Nama to response
		HargaPokok: barangModel.HargaPokok, // Assign HargaPokok to response
		HargaJual:  barangModel.HargaJual,  // Assign HargaJual to response
		TipeBarang: barangModel.TipeBarang, // Assign TipeBarang to response
		Stok:       barangModel.Stok,       // Assign Stok to response
		Model:      barangModel.Model,      // Assign Model to response
		CreatedBy:  barangModel.CreatedBy,  // Assign CreatedBy to response
		Histori:    histori,                // Include historical data in response
	}

	return &details, nil // Return the response struct
}

// Function to update item data
func UpdateBarang(id uint, barang model.Barang) (model.Barang, error) {
	repoBarang := modelfunc.Barang{
		Barang: model.Barang{
			ID:         uint64(id),          // Set ID for update
			KodeBarang: barang.KodeBarang,   // Set KodeBarang for update
			Nama:       barang.Nama,         // Set Nama for update
			HargaPokok: barang.HargaPokok,   // Set HargaPokok for update
			HargaJual:  barang.HargaJual,    // Set HargaJual for update
			TipeBarang: barang.TipeBarang,   // Set TipeBarang for update
			Stok:       barang.Stok,         // Set Stok for update
			CreatedBy:  barang.CreatedBy,    // Set CreatedBy for update
		},
	}
	err := repoBarang.Update(repository.Mysql.DB) // Update the item record in the database
	return repoBarang.Barang, err // Return updated item and any errors
}

// Function to delete an item
func DeleteBarang(id uint64) error {
	barang := modelfunc.Barang{
		Barang: model.Barang{
			ID: id, // Set ID for deletion
		},
	}
	return barang.Delete(repository.Mysql.DB) // Soft delete the item record
}
