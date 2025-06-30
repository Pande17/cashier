package utils

import (
	"cashier-machine/model"
	repository "cashier-machine/repository/config"
	"cashier-machine/repository/modelfunc"
	"log"
)

// Function to insert Invoice Item data into the database
func InsertInvoiceItem(invoiceItem model.InvoiceItem) error {
	// Insert the invoice item into the database
	if err := repository.Mysql.DB.Create(&invoiceItem).Error; err != nil {
		return err
	}
	return nil
}

// Function to insert Invoice data into the database
func InsertInvoiceData(data model.Invoice) (model.Invoice, error) {
	// Create an Invoice model instance with the provided data
	invoice := modelfunc.Invoice{
		Invoice: data, // Initialize with the provided invoice data
	}

	// Save the invoice data to the database (this will insert the invoice into the database)
	err := invoice.CreateInvoice(repository.Mysql.DB)
	if err != nil {
		log.Println("Error creating invoice:", err)
		return data, err // Return the error if creation fails
	}

	// Insert the invoice items into the database
	for _, item := range data.InvoiceItems {
		invoiceItem := model.InvoiceItem{
			KodeInvoice:  data.KodeInvoice,
			KodeBarang:   item.KodeBarang,
			Quantity:     item.Quantity,
			Harga:        item.Harga,
			DiskonBarang: item.DiskonBarang,
			ItemTotal:    item.ItemTotal, // Assuming ItemTotal is already calculated
		}

		// Insert each invoice item
		errInsertItem := InsertInvoiceItem(invoiceItem)
		if errInsertItem != nil {
			log.Println("Error inserting invoice item:", errInsertItem)
			return data, errInsertItem // Return the error if inserting item fails
		}
	}

	return invoice.Invoice, nil // Return the updated invoice after successfully inserting it and its items
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
	result, err := invoice.GetInvByID(repository.Mysql.DB) // Retrieve the sales record by ID
	if err != nil {
		return model.Invoice{}, err // Return the error if retrieval fails
	}
	return result.Invoice, nil // Return the retrieved sales record
}
