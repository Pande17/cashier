package generator

import (
	"cashier-machine/model"
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

// Function to generate ID for invoice item
func GenerateIDInvoiceItem(db *gorm.DB, kodeInvoice string) (string, error) {
	// Get the current date in YYYYMMDD format
	today := time.Now().Format("20060102")

	// Initialize the count variable to 1 (since we want to start from 1)
	var count int64
	var uniqueID string
	isUnique := false

	// Loop until a unique ID is generated
	for !isUnique {
		// Count the number of items for the given invoice code
		err := db.Model(&model.InvoiceItem{}).Where("kode_invoice = ?", kodeInvoice).Count(&count).Error
		if err != nil {
			return "", err
		}

		// Increment the count and generate the ID
		idIncrement := count + 1
		idString := fmt.Sprintf("%04d", idIncrement)
		uniqueID = fmt.Sprintf("ITM%s%s", today, idString)

		// Check if the generated ID already exists in the database
		var exists int64
		err = db.Model(&model.InvoiceItem{}).Where("id = ?", uniqueID).Count(&exists).Error
		if err != nil {
			return "", err
		}

		// If ID doesn't exist, it's unique, so we can exit the loop
		if exists == 0 {
			isUnique = true
		} else {
			// If ID exists, increment the count and try again
			count++
		}
	}

	// Return the unique ID
	return uniqueID, nil
}
