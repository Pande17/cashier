package controller

import (
	"cashier-machine/handler/generator"
	"cashier-machine/model"
	repository "cashier-machine/repository/config"
	"cashier-machine/repository/modelfunc"
	"cashier-machine/utils"
	"time"

	"github.com/gofiber/fiber/v2"
	logrus "github.com/sirupsen/logrus"
)

// InsertMember handles the insertion of new 'member' into the system
func InsertMember(c *fiber.Ctx) error {
	var req model.Member

	// Parse the incoming JSON body into the Member struct
	if err := c.BodyParser(&req); err != nil {
		return BadRequest(c, "Data yang dimasukkan tidak sesuai", "Invalid request body")
	}

	// Generate Member ID using GenerateIDMember function
	memberID, err := generator.GenerateIDMember(repository.Mysql.DB)
	if err != nil {
		logrus.Printf("Error generating member ID: %s\n", err.Error())
		return Conflict(c, "Server Error", "Gagal membuat ID member")
	}

	// Set the generated member ID to the request
	req.ID = memberID

	// Set the created_by field to "admin" for the new member
	req.CreatedBy = "admin" // Default value, you can update this later with middleware for authenticated user

	// Insert the member into the database
	member, err := utils.InsertMemberData(req)
	if err != nil {
		logrus.Printf("Error inserting member: %s\n", err.Error())
		return Conflict(c, "Server Error", "Gagal menambahkan data member")
	}

	// Return a successful response
	return OK(c, "Berhasil menambahkan data member", member)
}

// UpdateMember handles the update of an existing 'member' details
func UpdateMember(c *fiber.Ctx) error {
	memberID := c.Params("id") // Get the member ID from the route
	if memberID == "" {
		return BadRequest(c, "ID member tidak boleh kosong", "Member ID is required")
	}

	// Check if the member exists
	existingMember, err := modelfunc.GetMemberByID(repository.Mysql.DB, memberID)
	if err != nil {
		if err.Error() == "record not found" {
			return NotFound(c, "Data member tidak ditemukan", "Gagal mengambil data member")
		}
		return Conflict(c, "Server Error", "Gagal mengambil data member")
	}

	// Parse the incoming JSON body into the Member struct
	var req model.Member
	if err := c.BodyParser(&req); err != nil {
		return BadRequest(c, "Data yang dimasukkan tidak sesuai", "Invalid request body")
	}

	// Automatically set "CreatedBy" to "admin" (or you can use middleware to get the actual admin)
	req.CreatedBy = "admin" // You can replace this with actual logic to get the admin user from JWT or session

	// Set the updated values to the existing member
	// The updated data will overwrite the old data
	existingMember.Nama = req.Nama
	existingMember.Pic = req.Pic
	existingMember.Perusahaan = req.Perusahaan
	existingMember.Kategori = req.Kategori
	existingMember.Alamat = req.Alamat
	existingMember.NoTelp = req.NoTelp
	existingMember.Status = req.Status
	existingMember.CreatedBy = req.CreatedBy // Assign the 'CreatedBy' field
	existingMember.UpdatedAt = time.Now()    // Set the UpdatedBy field to "admin" or the current user

	// Update the member in the database
	updatedMember, err := utils.UpdateMemberData(existingMember) // Pass the updated member
	if err != nil {
		logrus.Printf("Error updating member: %s\n", err.Error())
		return Conflict(c, "Server Error", "Gagal mengupdate data member")
	}

	// Return a successful response
	return OK(c, "Berhasil mengupdate data member", updatedMember)
}

// DeleteMember handles the soft deletion of a 'member'
func DeleteMember(c *fiber.Ctx) error {
	memberID := c.Params("id") // Get the member ID from the route

	// If member ID is missing, return an error
	if memberID == "" {
		return BadRequest(c, "ID member tidak boleh kosong", "Member ID is required")
	}

	// Check if the member exists
	existingMember, err := modelfunc.GetMemberByID(repository.Mysql.DB, memberID)
	if err != nil {
		if existingMember.ID == "" {
			return NotFound(c, "Data member tidak ditemukan", "Gagal mengambil data member")
		}
		return Conflict(c, "Server Error", "Gagal mengambil data member")
	}

	// Check if the member is already soft-deleted
	// if existingMember.DeletedAt != (gorm.DeletedAt{}) {
	// 	return Conflict(c, "Data member sudah dihapus", "Member already deleted")
	// }

	// Soft delete the member by updating the deleted_at column
	err = utils.DeleteMemberData(memberID)
	if err != nil {
		logrus.Printf("Error deleting member: %s\n", err.Error())
		return Conflict(c, "Server Error", "Gagal menghapus data member")
	}

	return OK(c, "Data member sudah berhasil dihapus", nil)
}

// GetMemberByID retrieves a specific member by their ID
func GetMemberByID(c *fiber.Ctx) error {
	memberID := c.Params("id")

	// Retrieve the member from the database
	member, err := modelfunc.GetMemberByID(repository.Mysql.DB, memberID)
	if err != nil {
		if err.Error() == "record not found" {
			return NotFound(c, "Data member tidak ditemukan", "Gagal mengambil data member")
		}
		return Conflict(c, "Server Error", "Gagal mengambil data member")
	}

	// Return the retrieved member data
	return OK(c, "Berhasil mengambil data member", member)
}

// GetAllMembers retrieves all members from the system
func GetAllMembers(c *fiber.Ctx) error {
	logrus.Info("Fetching all members...")
	members, err := modelfunc.GetAllMembers(repository.Mysql.DB)
	if err != nil {
		logrus.Errorf("Error: %s", err.Error())
		return Conflict(c, "Server Error", "Gagal mengambil data member")
	}
	return OK(c, "Berhasil mengambil seluruh data member", members)
}
