package utils

import (
	"cashier-machine/model"
	repository "cashier-machine/repository/config"
	"cashier-machine/repository/modelfunc"
	"fmt"
	"time"
)

// Function to generate an invoice code automatically
func GenerateKodeInvoice(id uint64) string {
	kodeInvoice := fmt.Sprintf("INV%d", id) // Format the invoice code with the given ID
	return kodeInvoice
}

// Function to insert sales data into the database
func InsertPenjualanData(data model.Invoice) (model.Invoice, error) {
	data.Model.CreatedAt = time.Now() // Set the creation timestamp
	data.Model.UpdatedAt = time.Now() // Set the update timestamp

	// Get the discount based on the discount code
	if data.Kode_diskon != "" {
		diskon, err := GetDiskonByCode(data.Kode_diskon) // Retrieve discount information
		if err != nil {
			return data, err // Return the error if retrieval fails
		}

		// Calculate the discount amount
		var diskonAmount float64
		if diskon.Type == "PERCENT" {
			diskonAmount = data.Subtotal * (diskon.Amount / 100) // Calculate percentage-based discount
		} else {
			diskonAmount = diskon.Amount // Set fixed amount discount
		}

		// Apply the discount to the subtotal
		data.Diskon = diskonAmount
		data.Total = data.Subtotal - data.Diskon
	} else {
		data.Diskon = 0 // No discount applied
		data.Total = data.Subtotal
	}

	// Convert model.Penjualan to modelfunc.Penjualan
	penjualan := modelfunc.Penjualan{
		Penjualan: data, // Initialize with the provided sales data
	}

	// Save the sales data to the database to get the generated ID
	err := penjualan.CreatePenjualan(repository.Mysql.DB)
	if err != nil {
		return data, err // Return the error if saving fails
	}

	// Generate the invoice code after saving the sales data
	data.ID = penjualan.Penjualan.ID
	data.Kode_invoice = GenerateInvoice(data.ID)

	// Update the sales data with the newly generated invoice code
	penjualan.Penjualan.Kode_invoice = data.Kode_invoice
	err = penjualan.Update(repository.Mysql.DB)
	if err != nil {
		return data, err // Return the error if updating fails
	}

	return penjualan.Penjualan, nil // Return the updated sales record
}

// Function to get all sales data
func GetInvoices() ([]model.Invoice, error) {
	var penjualan modelfunc.Penjualan
	penjualanList, err := penjualan.GetAll(repository.Mysql.DB) // Retrieve all sales records
	if err != nil {
		return nil, err // Return the error if retrieval fails
	}

	// Convert []modelfunc.Penjualan to []model.Penjualan
	result := make([]model.Penjualan, len(penjualanList)) // Initialize a result slice
	for i, pj := range penjualanList {                    // Iterate through the retrieved sales records
		result[i] = pj.Penjualan // Add each record to the result slice
	}

	return result, nil // Return the slice of sales records
}

// Function to get sales data by ID
func GetPenjualanByID(id uint64) (model.Penjualan, error) {
	penjualan := modelfunc.Penjualan{
		Penjualan: model.Penjualan{
			ID: id, // Set the ID from the parameter
		},
	}
	result, err := penjualan.GetPByID(repository.Mysql.DB) // Retrieve the sales record by ID
	if err != nil {
		return model.Penjualan{}, err // Return the error if retrieval fails
	}
	return result.Penjualan, nil // Return the retrieved sales record
}
