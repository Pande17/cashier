package controller

import (
	"cashier-machine/handler/generator"
	"cashier-machine/model"
	repository "cashier-machine/repository/config"
	"cashier-machine/utils"

	"github.com/gofiber/fiber/v2" // Import Fiber for handling HTTP requests and responses
	"github.com/sirupsen/logrus"  // Import logrus for logging
)

// Function to create a new Barang (item)
func CreateBarang(c *fiber.Ctx) error {
	// Define a request struct for adding new Barang
	// var req model.Barang
	type AddBarangReq struct {
		Nama      string  `json:"nama"`       // Name of the item
		HargaBeli float64 `json:"harga_beli"` // Purchase price of the item
		HargaJual float64 `json:"harga_jual"` // Selling price of the item
		Kategori  string  `json:"kategori"`   // Category of the item
		Stok      uint    `json:"stok"`       // Stock quantity of the item
	}

	req := new(AddBarangReq)

	// Parse the request body JSON into the Barang model
	if err := c.BodyParser(req); err != nil {
		// Handle JSON parsing errors
		return BadRequest(c, "Invalid request body", "Failed to parse request body") // Response message for invalid body
	}

	// Generate the barang code before saving the data
	kodeBarang, err := generator.GenerateKodeBarang(repository.Mysql.DB)
	if err != nil {
		// Return an error response if generating the barang code fails
		logrus.Printf("Error generating item code: %s\n", err.Error())
		return Conflict(c, "Server Error", "Gagal membuat kode barang")
	}

	// Create the new Barang (item) in the database
	barang, errCreateBarang := utils.CreateBarang(model.Barang{
		KodeBarang: kodeBarang,
		Nama:       req.Nama,
		HargaBeli:  req.HargaBeli,
		HargaJual:  req.HargaJual,
		Kategori:   req.Kategori,
		Stok:       req.Stok,
		Model: model.Model{
			CreatedBy: "admin", // Set the creator of the record
		},
	})

	// Create a history record for the new Barang (item)
	// utils.CreateHistoriBarang(&model.Details{
	// 	ID:         barang.ID,
	// 	KodeBarang: req.Kode,
	// 	Nama:       req.Nama,
	// 	HargaPokok: req.HargaPokok,
	// 	HargaJual:  req.HargaJual,
	// 	TipeBarang: req.Tipe,
	// 	Stok:       req.Stok,
	// 	// CreatedBy:  req.CreateBy,
	// 	Histori: []model.HistoriASKM{},
	// }, req.Histori.Keterangan, int(req.Stok), req.Histori.Status)

	// Handle errors during creation and respond accordingly
	if errCreateBarang != nil {
		logrus.Printf("Error occurred: %s\n", errCreateBarang.Error())
		return Conflict(c, "Gagal membuat data Barang", errCreateBarang.Error()) // Response message for creation failure
	}

	// Return the newly created Barang's ID and kode_barang
	return OK(c, "Berhasil membuat data Barang", barang)
}

// Function to retrieve all Barang (items) from the database
func GetBarang(c *fiber.Ctx) error {
	// Retrieve all Barang (items) from the database
	dataBarang, err := utils.GetBarang()
	if err != nil {
		// Handle errors during retrieval
		logrus.Error("Failed to retrieve Barang list: ", err.Error())
		return Conflict(c, "Gagal mengambil data Barang", err.Error())
	}

	// Return all Barang data
	return OK(c, "Berhasil mengambil data Barang", dataBarang)
}

// Function to retrieve a specific Barang (item) by its ID
func GetBarangByID(c *fiber.Ctx) error {
	// Find Barang's ID from Params
	kodeBarang := c.Params("id")
	// if err != nil {
	// 	// Handle invalid ID format
	// 	return BadRequest(c, "Invalid ID", "ID must be a valid") // Response message for invalid ID
	// }

	// Check if there is an item with that ID
	dataBarang, err := utils.GetBarangByID(kodeBarang)
	if err != nil {
		// Handle case where record is not found
		if err.Error() == "record not found" {
			return NotFound(c, "Data Barang tidak ditemukan", "Gagal mengambil data Barang") // Response message for not found
		}

		// Handle other errors
		return Conflict(c, "Gagal mengambil data Barang", err.Error())
	}

	// Return the details of the Barang
	return OK(c, "Berhasil mengambil data Barang", dataBarang)
}

// Function to update Barang (item) by ID
func UpdateBarang(c *fiber.Ctx) error {
	// Find Barang's ID from Params
	kodeBarang := c.Params("id")

	// Define a request struct for updating Barang
	var updatedBarang model.Barang
	if err := c.BodyParser(&updatedBarang); err != nil {
		return BadRequest(c, "Invalid request body", "Failed to parse request body")
	}

	// Ensure kode_barang is not empty
	updatedBarang.KodeBarang = kodeBarang

	// Update the Barang in the database
	updatedData, err := utils.UpdateBarang(kodeBarang, updatedBarang)
	if err != nil {
		return Conflict(c, "Gagal memperbarui data Barang", err.Error())
	}

	// Return the updated Barang's ID
	return OK(c, "Berhasil memperbarui data Barang", updatedData)
}

// // Function to update stock of Barang (item) by ID
// func UpdateStok(c *fiber.Ctx) error {
// 	// Convert params to find ID
// 	barangID, err := strconv.Atoi(c.Params("id"))
// 	if err != nil {
// 		// Handle invalid ID format
// 		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
// 			"message": "Invalid ID", // Response message for invalid ID
// 		})
// 	}

// 	// Define a new request struct for stock and history
// 	var requestData struct {
// 		Stok        uint          `json:"stok"`         // New stock quantity
// 		HistoriStok model.Histori `json:"histori_stok"` // Stock history information
// 	}

// 	// Parse the request body JSON into the struct
// 	if err := c.BodyParser(&requestData); err != nil {
// 		// Handle JSON parsing errors
// 		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
// 			"message": "Invalid request body", // Response message for invalid body
// 		})
// 	}

// 	// Retrieve the existing Barang to update it
// 	existingBarang, err := utils.GetBarangByID(uint64(barangID))
// 	if err != nil {
// 		// Handle errors during retrieval
// 		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
// 			"message": "Failed to retrieve item", // Response message for retrieval failure
// 		})
// 	}

// 	// Update the stock of the Barang
// 	existingBarang.Stok = requestData.Stok

// 	// Update the Barang in the database
// 	updatedBarang := model.Barang{
// 		ID:         existingBarang.ID,
// 		KodeBarang: existingBarang.KodeBarang,
// 		Nama:       existingBarang.Nama,
// 		HargaPokok: existingBarang.HargaPokok,
// 		HargaJual:  existingBarang.HargaJual,
// 		TipeBarang: existingBarang.TipeBarang,
// 		Stok:       existingBarang.Stok,
// 	}
// 	updatedBarang, err = utils.UpdateBarang(uint(existingBarang.ID), updatedBarang)
// 	if err != nil {
// 		// Handle errors during update
// 		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
// 			"message": "Failed to update item", // Response message for update failure
// 		})
// 	}

// 	// Create the history record
// 	newHistori, err := utils.CreateHistoriBarang(existingBarang, requestData.HistoriStok.Keterangan, requestData.HistoriStok.Amount, requestData.HistoriStok.Status)
// 	if err != nil {
// 		// Handle errors during history creation
// 		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
// 			"message": "Failed to create history record", // Response message for history creation failure
// 		})
// 	}

// 	// Return the updated stock and history information
// 	return c.Status(fiber.StatusOK).JSON(fiber.Map{
// 		"id":          updatedBarang.ID,
// 		"kode_barang": updatedBarang.KodeBarang,
// 		"stok":        updatedBarang.Stok,
// 		"histori_stok": map[string]interface{}{
// 			"amount":     newHistori.Amount,
// 			"status":     newHistori.Status,
// 			"keterangan": newHistori.Keterangan,
// 		},
// 	})
// }

// Function to soft delete a Barang (item) by ID
func DeleteBarang(c *fiber.Ctx) error {
	// Convert params to find ID
	kodeBarang := c.Params("id")
	// if err != nil {
	// 	// Handle invalid ID format
	// 	return BadRequest(c, "Invalid ID", "ID must be a valid") // Response message for invalid ID
	// }

	// Attempt to delete the Barang
	err := utils.DeleteBarang(kodeBarang)
	if err != nil {
		// Handle cases where the record is not found or other errors
		if err.Error() == "record not found" {
			return NotFound(c, "Data Barang tidak ditemukan", "Gagal menghapus data Barang") // Response message for not found
		}
	}

	// Return confirmation of successful deletion
	return OK(c, "Berhasil menghapus data Barang", nil) // Response message for successful deletion
}
