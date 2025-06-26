package utils

import (
	"cashier-machine/model"
	repository "cashier-machine/repository/config"
	"cashier-machine/repository/modelfunc"
	"fmt"
	"time"

	"gorm.io/gorm"
)

// Fungsi untuk menghasilkan kode invoice
func GenerateKodeInvoice(db *gorm.DB) (string, error) {
	// Mendapatkan tanggal hari ini dalam format YYYYMMDD
	today := time.Now().Format("20060102")

	// Mencari jumlah invoice yang sudah ada untuk hari ini
	var count int64
	err := db.Model(&model.Invoice{}).Where("DATE(tanggal_beli) = ?", time.Now().Format("2006-01-02")).Count(&count).Error
	if err != nil {
		return "", err
	}

	// Increment ID berdasarkan jumlah invoice yang ada + 1
	idIncrement := count + 1

	// Format ID increment dengan 4 digit (0001, 0002, dst)
	idString := fmt.Sprintf("%04d", idIncrement)

	// Membuat kode invoice
	kodeInvoice := fmt.Sprintf("INV%s%s", today, idString)

	return kodeInvoice, nil
}

// Function to insert sales data into the database
func InsertInvoiceData(data model.Invoice) (model.Invoice, error) {
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

	// Convert model.Invoice to modelfunc.Invoice
	invoice := modelfunc.Invoice{
		Invoice: data, // Initialize with the provided sales data
	}

	// Save the sales data to the database to get the generated ID
	err := invoice.CreateInvoice(repository.Mysql.DB)
	if err != nil {
		return data, err // Return the error if saving fails
	}

	// Generate the invoice code after saving the sales data
	data.ID = invoice.Invoice.ID
	data.Kode_invoice = GenerateInvoice(data.ID)

	// Update the sales data with the newly generated invoice code
	invoice.Invoice.Kode_invoice = data.Kode_invoice
	err = invoice.Update(repository.Mysql.DB)
	if err != nil {
		return data, err // Return the error if updating fails
	}

	return invoice.Invoice, nil // Return the updated sales record
}

// Function to get all sales data
func GetInvoices() ([]model.Invoice, error) {
	var invoice modelfunc.Invoice
	invoiceList, err := invoice.GetAll(repository.Mysql.DB) // Retrieve all sales records
	if err != nil {
		return nil, err // Return the error if retrieval fails
	}

	// Convert []modelfunc.Invoice to []model.Invoice
	result := make([]model.Invoice, len(invoiceList)) // Initialize a result slice
	for i, inv := range invoiceList {                 // Iterate through the retrieved sales records
		result[i] = inv.Invoice // Add each record to the result slice
	}

	return result, nil // Return the slice of sales records
}

// Function to get sales data by ID
func GetInvoiceByID(kode string) (model.Invoice, error) {
	invoice := modelfunc.Invoice{
		Invoice: model.Invoice{
			KodeInvoice: kode, // Set the ID from the parameter
		},
	}
	result, err := invoice.GetPByID(repository.Mysql.DB) // Retrieve the sales record by ID
	if err != nil {
		return model.Invoice{}, err // Return the error if retrieval fails
	}
	return result.Invoice, nil // Return the retrieved sales record
}
