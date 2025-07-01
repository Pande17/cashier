package controller

import (
	"cashier-machine/handler/generator"
	"cashier-machine/model"
	repository "cashier-machine/repository/config"
	"cashier-machine/utils"
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
	logrus "github.com/sirupsen/logrus"
)

// InsertInvoiceData handles the insertion of new 'data invoice' into the system
func InsertInvoiceData(c *fiber.Ctx) error {
	// Define the structure for the request body
	type AddInvoiceReq struct {
		MemberID        string  `json:"member_id"`        // ID of the buyer
		JatuhTempo      string  `json:"jatuh_tempo"`      // Due date for payment
		Status          string  `json:"status"`           // Status of the transaction (e.g., "paid", "unpaid")
		Ppn             float64 `json:"ppn"`              // Value-added tax applied to the transaction
		BiayaPengiriman float64 `json:"biaya_pengiriman"` // Shipping cost for the transaction
		DiskonTotal     float64 `json:"diskon_total"`     // Total discount applied to the transaction
		// Diskon          float64 `json:"diskon"`           // Discount amount applied
		InvoiceItems []struct {
			Kode         string  `json:"kode_barang"`   // Item code
			Jumlah       uint    `json:"quantity"`      // Quantity of the item sold
			Harga        float64 `json:"harga"`         // Price of the item
			DiskonBarang float64 `json:"diskon_barang"` // Discount applied to the item
		} `json:"invoice_items"` // List of items sold
	}

	req := new(AddInvoiceReq)

	// Parse the incoming JSON body into the AddInvoiceReq struct
	if err := c.BodyParser(req); err != nil {
		// Return a Bad Request response if the body parsing fails
		return BadRequest(c, "Data yang dimasukkan tidak sesuai", "Invalid request body")
	}

	// Generate the invoice code before saving the data
	kodeInvoice, err := generator.GenerateKodeInvoice(repository.Mysql.DB)
	if err != nil {
		// Return an error response if generating the invoice code fails
		logrus.Printf("Error generating invoice code: %s\n", err.Error())
		return Conflict(c, "Server Error", "Gagal membuat kode invoice")
	}

	// Load the timezone
	loc, err := utils.SetTimezone()
	if err != nil {
		log.Printf("Error setting timezone: %s\n", err.Error())
		return Conflict(c, "Server Error", "Failed to set timezone")
	}

	// Get the current time in the selected timezone
	currentTime := utils.GetCurrentTimeInTimezone(loc)
	// formattedDate := utils.FormatDate(currentTime) // Format the current date to "13-Jun-2025"
	logrus.Printf("Current time in timezone: %s\n", currentTime)

	// Convert jatuh_tempo string to time.Time, handle null case
	var jatuhTempo *time.Time
	if req.JatuhTempo != "" {
		parsedDate, err := time.Parse("02-Jan-2006", req.JatuhTempo) // Parse date format (13-Jun-2025)
		if err != nil {
			return BadRequest(c, "Invalid date format", "Jatuh tempo format harus 'dd-Mon-yyyy'")
		}
		jatuhTempo = &parsedDate
	}

	// Calculate Subtotal and Total based on invoice items
	subtotal := CalculateSubtotal(req.InvoiceItems)
	total := CalculateTotal(subtotal, req.Ppn, req.DiskonTotal, req.BiayaPengiriman)

	// Create a Invoice model instance with the parsed data
	invoice := model.Invoice{
		KodeInvoice:     kodeInvoice, // Set the generated invoice code
		MemberID:        req.MemberID,
		TanggalBeli:     currentTime,          // Set the purchase date (current date)
		JatuhTempo:      jatuhTempo,           // Set due date, can be nil if not provided
		Status:          req.Status,           // Set the status of the transaction
		Ppn:             &req.Ppn,             // Set PPN (if applicable)
		BiayaPengiriman: &req.BiayaPengiriman, // Set shipping cost
		Subtotal:        subtotal,             // Subtotal from items
		DiskonTotal:     req.DiskonTotal,      // Total discount
		// Diskon:          req.Diskon,           // Individual discount
		Total: total, // Total after discount
		Model: model.Model{
			CreatedBy: "admin",     // Set the creator of the invoice
			CreatedAt: currentTime, // Set the creation time
			UpdatedAt: currentTime, // Set the update time
		},
	}
	logrus.Printf("Invoice Code to be inserted: %+v\n", kodeInvoice)

	// Insert the invoice data into the database
	_, errInsertInvoice := utils.InsertInvoiceData(invoice)
	if errInsertInvoice != nil {
		// Log the error and return an Internal Server Error response if insertion fails
		logrus.Printf("Error inserting invoice data: %s\n", errInsertInvoice.Error())
		return Conflict(c, "Server Error", "Gagal menambahkan data invoice")
	}

	// Insert the invoice items into the invoice_items table
	for _, item := range req.InvoiceItems {
		// Generate ID for the invoice item
		idInvoiceItem, err := generator.GenerateIDInvoiceItem(repository.Mysql.DB, kodeInvoice)
		if err != nil {
			// Log the error and return an Internal Server Error response if ID generation fails
			logrus.Printf("Error generating invoice item ID: %s\n", err.Error())
			return Conflict(c, "Server Error", "Gagal membuat ID untuk item invoice")
		}

		// Create an InvoiceItem model instance
		invoiceItem := model.InvoiceItem{
			ID:           idInvoiceItem,                                                  // Use the generated ID
			KodeInvoice:  kodeInvoice,                                                    // Set the invoice code
			KodeBarang:   item.Kode,                                                      // Item code
			Quantity:     item.Jumlah,                                                    // Quantity of the item sold
			Harga:        item.Harga,                                                     // Price of the item
			DiskonBarang: item.DiskonBarang,                                              // Discount applied to the item
			ItemTotal:    CalculateItemTotal(item.Harga, item.Jumlah, item.DiskonBarang), // Calculate item total
			Model: model.Model{
				CreatedBy: "admin",     // Set the creator of the item
				CreatedAt: currentTime, // Set the creation time
				UpdatedAt: currentTime, // Set the update time
			},
		}

		// Insert the item into the database
		errInsertItem := utils.InsertInvoiceItem(invoiceItem)
		if errInsertItem != nil {
			// Log the error and return an Internal Server Error response if insertion fails
			logrus.Printf("Error inserting invoice item: %s\n", errInsertItem.Error())
			return Conflict(c, "Server Error", "Gagal menambahkan data item invoice")
		}
	}

	// Return a successful response if insertion succeeds
	return OK(c, "Berhasil menambahkan data invoice", invoice)
}

// GetInvoices retrieves all sales data from the system
func GetInvoices(c *fiber.Ctx) error {
	// Retrieve all sales data from the database
	dataInvoices, err := utils.GetInvoices()
	if err != nil {
		// Log the error and return an Internal Server Error response if retrieval fails
		logrus.Error("Gagal dalam mengambil list invoice :", err.Error())
		return Conflict(c, "Server Error", "Gagal mengambil data invoice")
	}

	// if dataInvoices != nil {
	// 	// Log the retrieved data and its length
	// 	logrus.Info("Data Invoice yang diterima: ", dataInvoices)
	// 	logrus.Info("Jumlah item dalam data invoice: ", len(dataInvoices))
	// }

	// Return the retrieved sales data with a success message
	return OK(c, "Berhasil mengambil seluruh data invoice", dataInvoices)
}

// GetInvoiceByID retrieves a specific 'invoice data' by its ID
func GetInvoiceByID(c *fiber.Ctx) error {
	// Extract and convert the sale ID from the request parameters
	kodeInvoice := c.Params("kode_invoice")

	// Declare dataInvoice as a pointer to model.Invoice
	var dataInvoice model.Invoice

	// Retrieve the sale data by its ID
	dataInvoice, err := utils.GetInvoiceByID(kodeInvoice)
	if err != nil {
		if err.Error() == "record not found" {
			// Return a Not Found response if no record is found with the given ID
			return NotFound(c, "Data invoice tidak ditemukan", "Gagal mengambil data invoice")
		}

		// Return an Internal Server Error response for other errors
		return Conflict(c, "Server Error", "Gagal mengambil data invoice")
	}

	// Log the retrieved data and its length

	logrus.Info("Data Invoice yang diterima: ", dataInvoice)
	logrus.Info("Jumlah item dalam data invoice: ", len(dataInvoice.InvoiceItems))

	// Log individual details of the invoice, e.g., kode_invoice, total, and items
	logrus.Info("Kode Invoice: ", dataInvoice.KodeInvoice)
	logrus.Info("Total Invoice: ", dataInvoice.Total)

	// Log item details in the invoice
	for i, item := range dataInvoice.InvoiceItems {
		logrus.Infof("Item %d: Kode Barang: %s, Quantity: %d, Harga: %f, Item Total: %f",
			i+1, item.KodeBarang, item.Quantity, item.Harga, item.ItemTotal)
	}

	// Return the specific sale's data with a success message
	return OK(c, "Berhasil mengambil data invoice", dataInvoice)
}
