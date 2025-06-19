package routes

import (
	"cashier-machine/app/controller" // Import the controller package for handling route logic

	"github.com/gofiber/fiber/v2" // Import the Fiber package for creating and managing routes
)

// Function to setup the route API
func RouteSetup(r *fiber.App) {
	// r for 'route'

	// Define a route group for organizing the routes
	retailGroup := r.Group("")

	// Define routes for 'Barang'
	retailGroup.Get("/barang", controller.GetBarang)           // Route to get all Barang data
	retailGroup.Get("/barang/:id", controller.GetBarangByID)   // Route to get a specific Barang by ID
	retailGroup.Post("/barang", controller.CreateBarang)       // Route to create a new Barang record
	retailGroup.Put("/barang/:id", controller.UpdateBarang)    // Route to update an existing Barang by ID
	retailGroup.Put("/barang/stok/:id", controller.UpdateStok) // Route to update the stock of a Barang by ID
	retailGroup.Delete("/barang/:id", controller.DeleteBarang) // Route to delete a Barang by ID

	// Define routes for 'Penjualan'
	retailGroup.Get("/penjualan", controller.GetPenjualan)         // Route to get all Penjualan data
	retailGroup.Get("/penjualan/:id", controller.GetPenjualanByID) // Route to get a specific Penjualan by ID
	retailGroup.Post("/penjualan", controller.InsertPenjualanData) // Route to create a new Penjualan record

	// Define routes for 'Kode Diskon'
	retailGroup.Get("/kode-diskon", controller.GetKodeDiskon)         // Route to get all Kode Diskon data
	retailGroup.Get("/kode-diskon/:id", controller.GetDiskonByID)     // Route to get a specific Kode Diskon by ID
	retailGroup.Get("/kode-diskon-get-by-code", controller.GetByCode) // Route to get a specific Kode Diskon by Code
	retailGroup.Post("/kode-diskon", controller.CreateKodeDiskon)     // Route to create a new Kode Diskon record
	retailGroup.Put("/kode-diskon/:id", controller.UpdateCode)        // Route to update an existing Kode Diskon by ID
	retailGroup.Delete("/kode-diskon/:id", controller.DeleteKode)     // Route to delete a Kode Diskon by ID
}
