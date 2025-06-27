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

	// Generate the invoice code before saving the sales data
	kodeInvoice, err := GenerateKodeInvoice(repository.Mysql.DB)
	if err != nil {
		return data, err // Return the error if generating the invoice code fails
	}

	// Assign the generated invoice code to the invoice data
	data.KodeInvoice = kodeInvoice

	// Convert model.Invoice to modelfunc.Invoice
	invoice := modelfunc.Invoice{
		Invoice: data, // Initialize with the provided sales data
	}

	// Save the sales data to the database to get the generated ID
	err = invoice.CreateInvoice(repository.Mysql.DB)
	if err != nil {
		return data, err // Return the error if saving fails
	}

	// Update the sales data with the newly generated invoice code
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
