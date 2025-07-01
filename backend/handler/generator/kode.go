package generator

import (
	"cashier-machine/model"
	"fmt"
	"strconv"
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

// GenerateIDInvoiceItem generates a unique ID for an invoice item based on the invoice code.
func GenerateIDInvoiceItem(db *gorm.DB, kodeInvoice string) (string, error) {
	// Get the current date in YYYYMMDD format
	today := time.Now().Format("20060102")

	// Variable to store the next ID to be generated
	var lastItem model.InvoiceItem

	// Get the last item based on the given invoice code, ordered by ID
	err := db.Model(&model.InvoiceItem{}).
		Where("kode_invoice = ?", kodeInvoice).
		Order("id desc").
		First(&lastItem).Error

	if err != nil && err.Error() != "record not found" {
		return "", err
	}

	// Generate the new ID by incrementing the last number
	var newID string
	if err != nil && err.Error() == "record not found" {
		// If no record found, it means this is the first item for this invoice
		newID = fmt.Sprintf("ITM%s0001", today)
	} else {
		// Extract the numeric part of the last ID and increment it
		lastID := lastItem.ID[len(lastItem.ID)-4:] // Get the last 4 digits
		lastIDInt, err := strconv.Atoi(lastID)
		if err != nil {
			return "", err
		}
		newID = fmt.Sprintf("ITM%s%04d", today, lastIDInt+1)
	}

	// Return the newly generated ID
	return newID, nil
}

// Fungsi untuk menghasilkan kode ID Member
func GenerateIDMember(db *gorm.DB) (string, error) {
	// Mendapatkan tanggal hari ini dalam format YYYYMMDD
	today := time.Now().Format("20060102")

	// Mencari jumlah member yang sudah ada untuk hari ini
	var count int64
	err := db.Model(&model.Member{}).Where("DATE(created_at) = ?", time.Now().Format("2006-01-02")).Count(&count).Error
	if err != nil {
		return "", err
	}

	// Increment ID berdasarkan jumlah member yang ada + 1
	idIncrement := count + 1

	// Format ID increment dengan 4 digit (0001, 0002, dst)
	idString := fmt.Sprintf("%04d", idIncrement)

	// Membuat kode ID Member
	kodeMember := fmt.Sprintf("MEM%s%s", today, idString)

	return kodeMember, nil
}