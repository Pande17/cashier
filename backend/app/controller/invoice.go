package controller

import (
	"cashier-machine/model"
	"cashier-machine/utils"
	"strconv"

	"github.com/gofiber/fiber/v2"
	logrus "github.com/sirupsen/logrus"
)

// InsertInvoiceData handles the insertion of new 'data invoice' into the system
func InsertInvoiceData(c *fiber.Ctx) error {
	// Define the structure for the request body
	type AddInvoiceReq struct {
		MemberID    string  `json:"member_id"`   // ID of the buyer
		Subtotal    float64 `json:"subtotal"`    // Subtotal amount
		KodeDiskon  string  `json:"kode_diskon"` // Discount code applied
		Diskon      float64 `json:"diskon"`      // Discount amount
		Total       float64 `json:"total"`       // Total amount after discount
		CreatedBy   string  `json:"created_by"`  // Person who created the sale entry
		ItemInvoice []struct {
			Kode   string `json:"kode_barang"` // Item code
			Jumlah uint   `json:"jumlah"`      // Quantity of the item sold
		} `json:"item_invoice"` // List of items sold
	}

	req := new(AddInvoiceReq)

	// Parse the incoming JSON body into the AddInvoiceReq struct
	if err := c.BodyParser(req); err != nil {
		// Return a Bad Request response if the body parsing fails
		return c.Status(fiber.StatusBadRequest).
			JSON(map[string]interface{}{
				"message": "Invalid Body", // Error message for invalid body
			})
	}

	// Create a Invoice model instance with the parsed data
	invoice := model.Invoice{
		KodeInvoice:     req.KodeInvoice,
		MemberID:        req.MemberID,
		TanggalBeli:     req.NamaPembeli,
		JatuhTempo:      req.Subtotal,
		Ppn:             req.KodeDiskon,
		BiayaPengiriman: req.Diskon,
		Subtotal:        req.Total,
		DiskonTotal:     req.Diskon,
		Diskon:          req.Diskon,
		Total:           req.Total,
		Model: model.Model{
			CreatedBy: req.CreatedBy,          // Set the creator of the sale entry
			CreatedAt: utils.GetCurrentTime(), // Set the creation time
			UpdatedAt: utils.GetCurrentTime(), // Set the update time
		},
	}

	// Insert the sale data into the database
	_, errInsertInvoice := utils.InsertInvoiceData(invoice)
	if errInsertInvoice != nil {
		// Log the error and return an Internal Server Error response if insertion fails
		logrus.Printf("Terjadi error : %s\n", errInsertInvoice.Error())
		return c.Status(fiber.StatusInternalServerError).
			JSON(map[string]any{
				"message": "Server Error", // Error message for server error
			})
	}

	// Return a successful response if insertion succeeds
	return c.Status(fiber.StatusOK).
		JSON(map[string]any{
			"message": "Berhasil Menambahkan Invoice", // Success message
		})
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

	if dataInvoices != nil {
		// Log the retrieved data and its length
		logrus.Info("Data Invoice yang diterima: ", dataInvoices)
		logrus.Info("Jumlah item dalam data invoice: ", len(dataInvoices))
	}

	// Return the retrieved sales data with a success message
	return OK(c, "Berhasil mengambil data invoice", dataInvoices)
}

// GetInvoiceByID retrieves a specific 'invoice data' by its ID
func GetInvoiceByID(c *fiber.Ctx) error {
	// Extract and convert the sale ID from the request parameters
	invoiceID, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		// Return a Bad Request response if ID conversion fails
		return c.Status(fiber.StatusBadRequest).JSON(
			map[string]interface{}{
				"message": "Invalid ID", // Error message for invalid ID
			},
		)
	}

	// Retrieve the sale data by its ID
	dataInvoice, err := utils.GetInvoiceByID(uint64(invoiceID))
	if err != nil {
		if err.Error() == "record not found" {
			// Return a Not Found response if no record is found with the given ID
			return c.Status(fiber.StatusNotFound).JSON(
				map[string]interface{}{
					"message": "ID not found", // Error message for ID not found
				},
			)
		}

		// Return an Internal Server Error response for other errors
		return c.Status(fiber.StatusInternalServerError).JSON(
			map[string]interface{}{
				"message": "Server Error", // Error message for server error
			},
		)
	}

	// Return the specific sale's data with a success message
	return c.Status(fiber.StatusOK).JSON(
		map[string]interface{}{
			"data":    dataInvoice, // Sale data
			"message": "Success",   // Success message
		},
	)
}
