package controller

import (
	"projek/toko-retail/model"
	"projek/toko-retail/utils"
	"strconv"

	"github.com/gofiber/fiber/v2"
	logrus "github.com/sirupsen/logrus"
)

// InsertPenjualanData handles the insertion of new 'data penjualan' into the system
func InsertPenjualanData(c *fiber.Ctx) error {
	// Define the structure for the request body
	type AddPenjualanReq struct {
		ID            uint64  `json:"id"`           // Unique identifier for the sale
		KodeInvoice   string  `json:"kode_invoice"` // Invoice code
		NamaPembeli   string  `json:"nama_pembeli"` // Name of the buyer
		Subtotal      float64 `json:"subtotal"`     // Subtotal amount
		KodeDiskon    string  `json:"kode_diskon"`  // Discount code applied
		Diskon        float64 `json:"diskon"`       // Discount amount
		Total         float64 `json:"total"`        // Total amount after discount
		CreatedBy     string  `json:"created_by"`   // Person who created the sale entry
		ItemPenjualan []struct {
			Kode   string `json:"kode_barang"` // Item code
			Jumlah uint   `json:"jumlah"`      // Quantity of the item sold
		} `json:"item_penjualan"` // List of items sold
	}

	req := new(AddPenjualanReq)

	// Parse the incoming JSON body into the AddPenjualanReq struct
	if err := c.BodyParser(req); err != nil {
		// Return a Bad Request response if the body parsing fails
		return c.Status(fiber.StatusBadRequest).
			JSON(map[string]interface{}{
				"message": "Invalid Body", // Error message for invalid body
			})
	}

	// Create a Penjualan model instance with the parsed data
	penjualan := model.Penjualan{
		ID:           req.ID,
		Kode_invoice: req.KodeInvoice,
		Nama_pembeli: req.NamaPembeli,
		Subtotal:     req.Subtotal,
		Kode_diskon:  req.KodeDiskon,
		Diskon:       req.Diskon,
		Total:        req.Total,
		Created_by:   req.CreatedBy,
	}

	// Insert the sale data into the database
	_, errInsertPenjualan := utils.InsertPenjualanData(penjualan)
	if errInsertPenjualan != nil {
		// Log the error and return an Internal Server Error response if insertion fails
		logrus.Printf("Terjadi error : %s\n", errInsertPenjualan.Error())
		return c.Status(fiber.StatusInternalServerError).
			JSON(map[string]any{
				"message": "Server Error", // Error message for server error
			})
	}

	// Return a successful response if insertion succeeds
	return c.Status(fiber.StatusOK).
		JSON(map[string]any{
			"message": "Berhasil Menambahkan Penjualan", // Success message
		})
}

// GetPenjualan retrieves all sales data from the system
func GetPenjualan(c *fiber.Ctx) error {
	// Retrieve all sales data from the database
	dataPenjualan, err := utils.GetPenjualan()
	if err != nil {
		// Log the error and return an Internal Server Error response if retrieval fails
		logrus.Error("Gagal dalam mengambil list penjualan :", err.Error())
		return c.Status(fiber.StatusInternalServerError).JSON(
			map[string]any{
				"message": "Server Error", // Error message for server error
			},
		)
	}

	if dataPenjualan != nil {
		// Log the retrieved data and its length
		logrus.Info("Data Penjualan yang diterima: ", dataPenjualan)
		logrus.Info("Jumlah item dalam data penjualan: ", len(dataPenjualan))
	}

	// Return the retrieved sales data with a success message
	return c.Status(fiber.StatusOK).JSON(
		map[string]any{
			"penjualan": dataPenjualan,               // Sales data
			"message":   "Success Get All Penjualan", // Success message
		},
	)
}

// GetPenjualanByID retrieves a specific 'penjualan data' by its ID
func GetPenjualanByID(c *fiber.Ctx) error {
	// Extract and convert the sale ID from the request parameters
	penjualanID, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		// Return a Bad Request response if ID conversion fails
		return c.Status(fiber.StatusBadRequest).JSON(
			map[string]interface{}{
				"message": "Invalid ID", // Error message for invalid ID
			},
		)
	}

	// Retrieve the sale data by its ID
	dataPenjualan, err := utils.GetPenjualanByID(uint64(penjualanID))
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
			"data":    dataPenjualan, // Sale data
			"message": "Success",     // Success message
		},
	)
}
