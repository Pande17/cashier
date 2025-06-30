package controller

import (
	"cashier-machine/model"
	repository "cashier-machine/repository/config"
	"cashier-machine/repository/modelfunc"
	"cashier-machine/utils"

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
	var req model.Member

	// Parse the incoming JSON body into the Member struct
	if err := c.BodyParser(&req); err != nil {
		return BadRequest(c, "Data yang dimasukkan tidak sesuai", "Invalid request body")
	}

	// Update the member in the database
	member, err := utils.UpdateMemberData(req)
	if err != nil {
		logrus.Printf("Error updating member: %s\n", err.Error())
		return Conflict(c, "Server Error", "Gagal mengupdate data member")
	}

	// Return a successful response
	return OK(c, "Berhasil mengupdate data member", member)
}

// DeleteMember handles the soft deletion of a 'member'
func DeleteMember(c *fiber.Ctx) error {
	memberID := c.Params("id")

	// Soft delete the member
	err := utils.DeleteMemberData(memberID)
	if err != nil {
		logrus.Printf("Error deleting member: %s\n", err.Error())
		return Conflict(c, "Server Error", "Gagal menghapus data member")
	}

	// Return a successful response
	return OK(c, "Berhasil menghapus data member", nil)
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
