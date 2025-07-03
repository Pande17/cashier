package generator

import (
	"cashier-machine/model"
	"fmt"
	"strconv"
	"time"

	"github.com/siruspen/logrus"
	"gorm.io/gorm"
)

// Fungsi untuk menghasilkan kode barang
func GenerateKodeBarang(db *gorm.DB) (string, error) {
	// Mendapatkan tanggal hari ini dalam format YYYYMMDD
	today := time.Now().Format("20060102")

	// Mencari jumlah barang yang sudah ada untuk hari ini
	var count int64
	err := db.Model(&model.Barang{}).Where("DATE(created_at) = ?", time.Now().Format("2006-01-02")).Count(&count).Error
	if err != nil {
		return "", err
	}

	// Increment ID berdasarkan jumlah barang yang ada + 1
	idIncrement := count + 1

	// Format ID increment dengan 3 digit (001, 002, dst)
	idString := fmt.Sprintf("%04d", idIncrement)

	// Membuat kode barang
	kodeBarang := fmt.Sprintf("BRG%s%s", today, idString)

	return kodeBarang, nil
}

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

func GenerateIDMember(db *gorm.DB) (string, error) {
	// Get today's date in YYYYMMDD format
	today := time.Now().Format("20060102")

	// Find the highest existing ID for today
	var lastMember model.Member
	err := db.Model(&model.Member{}).
		Where("id LIKE ?", fmt.Sprintf("MEM%s%%", today)).
		Order("id DESC").
		First(&lastMember).Error
	if err != nil && err.Error() != "record not found" {
		return "", err
	}

	logrus.Printf("Last member ID found: %s\n", lastMember.ID)

	// Extract the numeric part of the last ID (skip MEM + today)
	var lastIDNum int
	if err == nil {
		// Extract the last 4 digits from the last member ID
		_, err := fmt.Sscanf(lastMember.ID[len("MEM"+today):], "%d", &lastIDNum)
		if err != nil {
			return "", err
		}
	}

	// If no record was found, start from 1
	if err != nil || lastIDNum == 0 {
		lastIDNum = 1
	} else {
		// Increment the ID by 1
		lastIDNum++
	}

	// Format the new ID with leading zeros (e.g., MEM202507010002)
	newID := fmt.Sprintf("MEM%s%04d", today, lastIDNum)

	// Check if the ID is already taken (including soft-deleted records)
	var existingMember model.Member
	err = db.First(&existingMember, "id = ?", newID).Error
	if err == nil {
		// If the member exists (including soft deleted), recursively generate a new ID
		return GenerateIDMember(db)
	}

	return newID, nil
}
