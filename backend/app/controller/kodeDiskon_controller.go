package controller

import (
	"fmt"
	"projek/toko-retail/model"
	"projek/toko-retail/utils"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

// CreateKodeDiskon handles the creation of a new discount code
func CreateKodeDiskon(c *fiber.Ctx) error {
	type AddKodeDiskon struct {
		Kode_diskon string  `json:"kode_diskon"` // Discount code
		Amount      float64 `json:"amount"`      // Discount amount
		Type        string  `json:"type"`        // Discount type (PERCENT or FIXED)
	}

	// Parse the request body into AddKodeDiskon struct
	req := new(AddKodeDiskon)
	if err := c.BodyParser(req); err != nil {
		return c.Status(fiber.StatusBadRequest).
			JSON(map[string]any{
				"message": "Invalid Body", // Response message for invalid body
			})
	}

	// Create a new discount code
	diskon, errDiskon := utils.CreateKodeDiskon(model.Diskon{
		KodeDiskon: req.Kode_diskon,
		Amount:     req.Amount,
		Type:       req.Type,
	})

	// Handle errors during creation
	if errDiskon != nil {
		logrus.Printf("Terjadi error : %s\n", errDiskon.Error())
		return c.Status(fiber.StatusInternalServerError).
			JSON(map[string]any{
				"message": "Server Error", // Response message for server error
			})
	}

	// Return the created discount code
	return c.Status(fiber.StatusOK).
		JSON(map[string]any{
			"data": diskon, // Response message containing the created discount code
		})
}

// GetKodeDiskon retrieves all discount codes
func GetKodeDiskon(c *fiber.Ctx) error {
	dataDiskon, err := utils.GetDiskon()
	if err != nil {
		logrus.Error("Gagal dalam mengambil list Diskon: ", err.Error())
		return c.Status(fiber.StatusInternalServerError).JSON(
			map[string]any{
				"message": "Server Error", // Response message for server error
			},
		)
	}

	// Return the list of all discount codes
	return c.Status(fiber.StatusOK).JSON(
		map[string]any{
			"data":  dataDiskon, // Response message containing the list of discount codes
			"title": "Success Get All Kode Diskon",
		},
	)
}

// GetByCode retrieves a discount code by its code and applies it to a subtotal
func GetByCode(c *fiber.Ctx) error {
	DiskonCode := c.Query("kode-diskon")
	SubtotalStr := c.Query("subtotal")

	fmt.Println("Received kode-diskon:", DiskonCode)
	fmt.Println("Received subtotal:", SubtotalStr)

	// Validate the discount code
	if DiskonCode == "" {
		fmt.Println("kode-diskon is empty")
		return c.Status(fiber.StatusBadRequest).JSON(
			map[string]any{
				"message": "Invalid kode-diskon", // Response message for invalid discount code
			},
		)
	}

	// Retrieve the discount code details
	dataDiskon, err := utils.GetDiskonByCode(DiskonCode)
	if err != nil {
		fmt.Println("Error retrieving discount code:", err)
		if err.Error() == "record not found" {
			return c.Status(fiber.StatusNotFound).JSON(
				map[string]any{
					"message": "Discount Code not found", // Response message for discount code not found
				},
			)
		}
		return c.Status(fiber.StatusInternalServerError).JSON(
			map[string]any{
				"message": "Server Error", // Response message for server error
			},
		)
	}

	// Calculate the final amount after applying the discount
	var response fiber.Map
	if SubtotalStr != "" {
		subtotal, err := strconv.ParseFloat(SubtotalStr, 64)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(
				map[string]any{
					"message": "Invalid Subtotal", // Response message for invalid subtotal
				},
			)
		}

		var finalAmount float64
		if dataDiskon.Type == "PERCENT" {
			finalAmount = subtotal - (subtotal * (dataDiskon.Amount / 100))
		} else {
			finalAmount = subtotal - dataDiskon.Amount
		}

		response = fiber.Map{
			"subtotal": subtotal,
			"diskon":   dataDiskon.Amount,
			"total":    finalAmount, // Response message with calculated final amount
		}
	} else {
		response = fiber.Map{
			"data": dataDiskon, // Response message containing the discount code details
		}
	}

	// Return the calculated final amount or discount code details
	return c.JSON(response)
}

// GetDiskonByID retrieves a discount code by its ID
func GetDiskonByID(c *fiber.Ctx) error {
	DiskonID, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(
			map[string]any{
				"message": "Invalid ID", // Response message for invalid ID
			},
		)
	}

	// Retrieve the discount code details by ID
	dataDiskon, err := utils.GetDiskonByID(uint(DiskonID))
	if err != nil {
		if err.Error() == "record not found" {
			return c.Status(fiber.StatusNotFound).JSON(
				map[string]any{
					"message": "ID not found", // Response message for ID not found
				},
			)
		}
	}

	// Return the retrieved discount code details
	return c.Status(fiber.StatusOK).JSON(
		map[string]any{
			"data": dataDiskon, // Response message containing the discount code details
		},
	)
}

// UpdateCode updates an existing discount code
func UpdateCode(c *fiber.Ctx) error {
	DiskonID, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid ID", // Response message for invalid ID
		})
	}

	// Parse the request body into updatedDiskon struct
	var updatedDiskon model.Diskon
	if err := c.BodyParser(&updatedDiskon); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid request body", // Response message for invalid body
		})
	}

	// Update the discount code details
	dataDiskon, err := utils.UpdateDiskon(uint(DiskonID), updatedDiskon)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to update item", // Response message for update failure
		})
	}

	// Return the updated discount code details
	return c.Status(fiber.StatusOK).JSON(
		map[string]any{
			"data": dataDiskon, // Response message containing the updated discount code details
		},
	)
}

// DeleteKode deletes an existing discount code by ID
func DeleteKode(c *fiber.Ctx) error {
	KodeID, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(
			map[string]any{
				"message": "Invalid ID", // Response message for invalid ID
			},
		)
	}

	// Delete the discount code by ID
	err = utils.DeleteKode(uint64(KodeID))
	if err != nil {
		if err.Error() == "record not found" {
			return c.Status(fiber.StatusNotFound).JSON(
				map[string]any{
					"message": "ID not found", // Response message for ID not found
				},
			)
		}
	}

	// Return a success message after deletion
	return c.Status(fiber.StatusOK).JSON(map[string]any{
		"message": "Deleted Successfully", // Response message for successful deletion
	})
}

// ApplyDiskon applies a discount code to a subtotal and returns the final amount
func ApplyDiskon(c *fiber.Ctx) error {
	type DiscountRequest struct {
		KodeDiskon string  `json:"kode_diskon"` // Discount code
		Subtotal   float64 `json:"subtotal"`   // Subtotal amount
	}

	// Parse the request body into DiscountRequest struct
	req := new(DiscountRequest)
	if err := c.BodyParser(req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(
			map[string]any{
				"message": "Invalid request body", // Response message for invalid body
			},
		)
	}

	// Retrieve the discount code details
	dataDiskon, err := utils.GetDiskonByCode(req.KodeDiskon)
	if err != nil {
		if err.Error() == "record not found" {
			return c.Status(fiber.StatusNotFound).JSON(
				map[string]any{
					"message": "Discount code not found", // Response message for discount code not found
				},
			)
		}
		return c.Status(fiber.StatusInternalServerError).JSON(
			map[string]any{
				"message": "Server error", // Response message for server error
			},
		)
	}

	// Calculate the final amount after applying the discount
	var finalAmount float64
	if dataDiskon.Type == "PERCENT" {
		finalAmount = req.Subtotal - (req.Subtotal * (dataDiskon.Amount / 100))
	} else {
		finalAmount = req.Subtotal - dataDiskon.Amount
	}

	// Return the calculated final amount
	return c.Status(fiber.StatusOK).JSON(
		map[string]any{
			"subtotal": req.Subtotal,
			"diskon":   dataDiskon.Amount,
			"total":    finalAmount, // Response message with calculated final amount
		},
	)
}
