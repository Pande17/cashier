package modelfunc

import (
	"cashier-machine/model" // Import the model package for the Invoice model

	"gorm.io/gorm" // Import the GORM package for ORM functionality
)

// Define the Invoice struct which embeds the model.Invoice struct
type Invoice struct {
	model.Invoice
}

// CreateInvoice inserts a new Invoice record into the database
func (inv *Invoice) CreateInvoice(db *gorm.DB) error {
	err := db.Create(&inv).Error // Attempt to create a new Invoice record
	if err != nil {              // Check if there was an error
		return err // Return the error if creation fails
	}
	return nil // Return nil if creation is successful
}

// GetAll retrieves all Invoice records from the database
func (inv *Invoice) GetAll(db *gorm.DB) ([]Invoice, error) {
	res := []Invoice{}                                 // Initialize an empty slice of Invoice
	err := db.Preload("InvoiceItems").Find(&res).Error // Query all Invoice records and store them in res
	if err != nil {                                    // Check if there was an error
		return []Invoice{}, err // Return an empty slice and the error
	}
	return res, nil // Return the retrieved records and nil (no error)
}

// GetPByID retrieves a single Invoice record by its ID
func (inv *Invoice) GetInvByID(db *gorm.DB) (Invoice, error) {
	res := Invoice{}                                                                                               // Initialize an empty Invoice
	err := db.Model(Invoice{}).Preload("InvoiceItems").Where("kode_invoice = ?", inv.KodeInvoice).Take(&res).Error // Query for a Invoice record with the given ID
	if err != nil {                                                                                                // Check if there was an error
		return Invoice{}, err // Return an empty Invoice and the error
	}
	return res, nil // Return the retrieved Invoice and nil (no error)
}

// Update modifies an existing Invoice record in the database
func (inv *Invoice) Update(db *gorm.DB) error {
	err := db.Model(Invoice{}).Where("kode_invoice = ?", inv.KodeInvoice).Updates(&inv).Error // Update the Invoice record with the given ID
	if err != nil {                                                                           // Check if there was an error
		return err // Return the error if the update fails
	}
	return nil // Return nil if the update is successful
}

// UpdateKodeInvoice updates the invoice code of an existing Invoice record
func (inv *Invoice) UpdateKodeInvoice(db *gorm.DB) error {
	err := db.Model(Invoice{}).Where("kode_invoice = ?", inv.KodeInvoice).Updates(&inv).Error // Update the Invoice record with the given ID (specifically for invoice code)
	if err != nil {                                                                           // Check if there was an error
		return err // Return the error if the update fails
	}
	return nil // Return nil if the update is successful
}
