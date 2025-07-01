package routes

import (
	"cashier-machine/handler/controller" // Import the controller package for handling route logic

	"github.com/gofiber/fiber/v2" // Import the Fiber package for creating and managing routes
	"github.com/gofiber/fiber/v2/middleware/cors"
)

// Function to setup the route API
func RouteSetup(r *fiber.App) {
	// r for 'route'

	// Define a route group for organizing the routes
	cashierGroup := r.Group("")

	// Middleware to handle CORS for the cashier routes
	// This middleware allows cross-origin requests to the cashier routes
	cashierGroup.Use(cors.New(cors.Config{
		AllowOrigins: "*", // Bisa disesuaikan dengan domain yang ingin diizinkan
		AllowMethods: "GET,POST,DELETE,PUT",
		AllowHeaders: "Origin, Content-Type, Accept",
	}))

	// Define routes for 'Barang'
	cashierGroup.Get("/barang", controller.GetBarang) // Route to get all Barang data
	// cashierGroup.Get("/barang/:id", controller.GetBarangByID)   // Route to get a specific Barang by ID
	// cashierGroup.Post("/barang", controller.CreateBarang)       // Route to create a new Barang record
	// cashierGroup.Put("/barang/:id", controller.UpdateBarang)    // Route to update an existing Barang by ID
	// cashierGroup.Put("/barang/stok/:id", controller.UpdateStok) // Route to update the stock of a Barang by ID
	// cashierGroup.Delete("/barang/:id", controller.DeleteBarang) // Route to delete a Barang by ID

	// Define routes for 'invoice'
	cashierGroup.Get("/invoice", controller.GetInvoices)                  // Route to get all Invoice data
	cashierGroup.Get("/invoice/:kode_invoice", controller.GetInvoiceByID) // Route to get a specific Invoice by ID
	cashierGroup.Post("/invoice", controller.InsertInvoiceData)           // Route to create a new Invoice record

	// Define routes for 'Member'
	cashierGroup.Get("/member", controller.GetAllMembers)       // Route to get all Member data
	cashierGroup.Get("/member/:id", controller.GetMemberByID)   // Route to get a specific Member by ID
	cashierGroup.Post("/member", controller.InsertMember)       // Route to create a new
	cashierGroup.Put("/member/:id", controller.UpdateMember)    // Route to update an existing Member
	cashierGroup.Delete("/member/:id", controller.DeleteMember) // Route to soft delete a Member by ID

	// Define routes for 'Kode Diskon'
	cashierGroup.Get("/kode-diskon", controller.GetKodeDiskon)         // Route to get all Kode Diskon data
	cashierGroup.Get("/kode-diskon/:id", controller.GetDiskonByID)     // Route to get a specific Kode Diskon by ID
	cashierGroup.Get("/kode-diskon-get-by-code", controller.GetByCode) // Route to get a specific Kode Diskon by Code
	cashierGroup.Post("/kode-diskon", controller.CreateKodeDiskon)     // Route to create a new Kode Diskon record
	cashierGroup.Put("/kode-diskon/:id", controller.UpdateCode)        // Route to update an existing Kode Diskon by ID
	cashierGroup.Delete("/kode-diskon/:id", controller.DeleteKode)     // Route to delete a Kode Diskon by ID
}
