package controller

import (
	"projek/toko-retail/model"
	"projek/toko-retail/utils"
	"strconv"

	"github.com/gofiber/fiber/v2"   // Import Fiber for handling HTTP requests and responses
	"github.com/sirupsen/logrus"   // Import logrus for logging
)

// Function to create a new Barang (item)
func CreateBarang(c *fiber.Ctx) error {
	// Define a request struct for adding new Barang
	type AddBarangReq struct {
		Kode       string  `json:"kode_barang"`   // Barang code
		Nama       string  `json:"nama_barang"`   // Barang name
		HargaPokok float64 `json:"harga_pokok"`   // Cost price
		HargaJual  float64 `json:"harga_jual"`    // Selling price
		Tipe       string  `json:"tipe_barang"`   // Barang type
		Stok       uint    `json:"stok"`          // Stock quantity
		CreateBy   string  `json:"created_by"`    // Creator's name
		Histori    struct {
			Amount     int    `json:"amount"`      // Stock change amount
			Status     string `json:"status"`      // Stock change status
			Keterangan string `json:"keterangan"`  // Description of the stock change
		} `json:"histori_stok"`  // Stock history information
	}

	// Parse the request body JSON into the AddBarangReq struct
	req := new(AddBarangReq)
	if err := c.BodyParser(req); err != nil {
		// Handle JSON parsing errors
		return c.Status(fiber.StatusBadRequest).
			JSON(map[string]any{
				"message": "Invalid Body",  // Response message for invalid body
			})
	}

	// Create the new Barang (item) in the database
	barang, errCreateBarang := utils.CreateBarang(model.Barang{
		KodeBarang: req.Kode,
		Nama:       req.Nama,
		HargaPokok: req.HargaPokok,
		HargaJual:  req.HargaJual,
		TipeBarang: req.Tipe,
		Stok:       req.Stok,
		CreatedBy:  req.CreateBy,
	})

	// Create a history record for the new Barang (item)
	utils.CreateHistoriBarang(&model.Details{
		ID:         barang.ID,
		KodeBarang: req.Kode,
		Nama:       req.Nama,
		HargaPokok: req.HargaPokok,
		HargaJual:  req.HargaJual,
		TipeBarang: req.Tipe,
		Stok:       req.Stok,
		CreatedBy:  req.CreateBy,
		Histori:    []model.HistoriASKM{},
	}, req.Histori.Keterangan, int(req.Stok), req.Histori.Status)

	// Handle errors during creation and respond accordingly
	if errCreateBarang != nil {
		logrus.Printf("Error occurred: %s\n", errCreateBarang.Error())
		return c.Status(fiber.StatusInternalServerError).
			JSON(map[string]any{
				"message": "Server Error",  // Response message for server error
			})
	}

	// Return the newly created Barang's ID and kode_barang
	return c.Status(fiber.StatusOK).
		JSON(map[string]any{
			"id":          barang.ID,
			"kode_barang": barang.KodeBarang,
		})
}

// Function to retrieve all Barang (items) from the database
func GetBarang(c *fiber.Ctx) error {
	// Retrieve all Barang (items) from the database
	dataBarang, err := utils.GetBarang()
	if err != nil {
		// Handle errors during retrieval
		logrus.Error("Failed to retrieve Barang list: ", err.Error())
		return c.Status(fiber.StatusInternalServerError).JSON(
			map[string]any{
				"message": "Server Error",  // Response message for server error
			},
		)
	}

	// Return all Barang data
	return c.Status(fiber.StatusOK).JSON(
		map[string]any{
			"data":    dataBarang,
			"message": "Success Get All Barang",  // Success message
		},
	)
}

// Function to retrieve a specific Barang (item) by its ID
func GetBarangByID(c *fiber.Ctx) error {
	// Find Barang's ID from Params
	barangID, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		// Handle invalid ID format
		return c.Status(fiber.StatusBadRequest).JSON(
			map[string]interface{}{
				"message": "Invalid ID",  // Response message for invalid ID
			},
		)
	}

	// Check if there is an item with that ID
	dataBarang, err := utils.GetBarangByID(uint64(barangID))
	if err != nil {
		// Handle case where record is not found
		if err.Error() == "record not found" {
			return c.Status(fiber.StatusNotFound).JSON(
				map[string]interface{}{
					"message": "ID not found",  // Response message for ID not found
				},
			)
		}

		// Handle other errors
		return c.Status(fiber.StatusInternalServerError).JSON(
			map[string]interface{}{
				"message": err.Error(),  // Response message with error details
			},
		)
	}

	// Return the details of the Barang
	return c.Status(fiber.StatusOK).JSON(
		map[string]interface{}{
			"data":    dataBarang,
			"message": "Success",  // Success message
		},
	)
}

// Function to update Barang (item) by ID
func UpdateBarang(c *fiber.Ctx) error {
	// Find Barang's ID from Params
	barangID, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		// Handle invalid ID format
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid ID",  // Response message for invalid ID
		})
	}

	// Parse the request body JSON into the Barang model
	var updatedBarang model.Barang
	if err := c.BodyParser(&updatedBarang); err != nil {
		// Handle JSON parsing errors
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid request body",  // Response message for invalid body
		})
	}

	// Update the Barang in the database
	dataBarang, err := utils.UpdateBarang(uint(barangID), updatedBarang)
	if err != nil {
		// Handle errors during update
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to update item",  // Response message for update failure
		})
	}

	// Return the updated Barang's ID
	return c.Status(fiber.StatusOK).JSON(
		map[string]any{
			"id": dataBarang.ID,
		},
	)
}

// Function to update stock of Barang (item) by ID
func UpdateStok(c *fiber.Ctx) error {
	// Convert params to find ID
	barangID, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		// Handle invalid ID format
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid ID",  // Response message for invalid ID
		})
	}

	// Define a new request struct for stock and history
	var requestData struct {
		Stok        uint          `json:"stok"`         // New stock quantity
		HistoriStok model.Histori `json:"histori_stok"` // Stock history information
	}

	// Parse the request body JSON into the struct
	if err := c.BodyParser(&requestData); err != nil {
		// Handle JSON parsing errors
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid request body",  // Response message for invalid body
		})
	}

	// Retrieve the existing Barang to update it
	existingBarang, err := utils.GetBarangByID(uint64(barangID))
	if err != nil {
		// Handle errors during retrieval
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to retrieve item",  // Response message for retrieval failure
		})
	}

	// Update the stock of the Barang
	existingBarang.Stok = requestData.Stok

	// Update the Barang in the database
	updatedBarang := model.Barang{
		ID:         existingBarang.ID,
		KodeBarang: existingBarang.KodeBarang,
		Nama:       existingBarang.Nama,
		HargaPokok: existingBarang.HargaPokok,
		HargaJual:  existingBarang.HargaJual,
		TipeBarang: existingBarang.TipeBarang,
		Stok:       existingBarang.Stok,
	}
	updatedBarang, err = utils.UpdateBarang(uint(existingBarang.ID), updatedBarang)
	if err != nil {
		// Handle errors during update
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to update item",  // Response message for update failure
		})
	}

	// Create the history record
	newHistori, err := utils.CreateHistoriBarang(existingBarang, requestData.HistoriStok.Keterangan, requestData.HistoriStok.Amount, requestData.HistoriStok.Status)
	if err != nil {
		// Handle errors during history creation
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to create history record",  // Response message for history creation failure
		})
	}

	// Return the updated stock and history information
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"id":          updatedBarang.ID,
		"kode_barang": updatedBarang.KodeBarang,
		"stok":        updatedBarang.Stok,
		"histori_stok": map[string]interface{}{
			"amount":     newHistori.Amount,
			"status":     newHistori.Status,
			"keterangan": newHistori.Keterangan,
		},
	})
}

// Function to soft delete a Barang (item) by ID
func DeleteBarang(c *fiber.Ctx) error {
	// Convert params to find ID
	barangID, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		// Handle invalid ID format
		return c.Status(fiber.StatusBadRequest).JSON(
			map[string]any{
				"message": "Invalid ID",  // Response message for invalid ID
			},
		)
	}

	// Attempt to delete the Barang
	err = utils.DeleteBarang(uint64(barangID))
	if err != nil {
		// Handle cases where the record is not found or other errors
		if err.Error() == "record not found" {
			return c.Status(fiber.StatusNotFound).JSON(
				map[string]any{
					"message": "ID not found",  // Response message for ID not found
				},
			)
		}
	}

	// Return confirmation of successful deletion
	return c.Status(fiber.StatusOK).JSON(map[string]any{
		"message": "Deleted Successfully",  // Response message for successful deletion
	})
}
