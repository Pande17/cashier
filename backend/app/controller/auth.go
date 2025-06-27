package controller

// import (
// 	"cashier-machine/utils"

// 	"github.com/gofiber/fiber/v2"
// )

// // Endpoint untuk menghasilkan token JWT
// func Login(c *fiber.Ctx) error {
// 	// Misalkan Anda mendapatkan ID dan status admin dari login
// 	userID := "12345"
// 	isAdmin := true

// 	// Generate token
// 	token, err := utils.GenerateToken(userID, isAdmin)
// 	if err != nil {
// 		return c.Status(fiber.StatusInternalServerError).JSON(map[string]interface{}{
// 			"message": "Failed to generate token",
// 		})
// 	}

// 	// Kirim token ke klien
// 	return c.JSON(map[string]interface{}{
// 		"token": token,
// 	})
// }

// // Middleware untuk memverifikasi JWT dan memastikan hanya admin yang bisa mengakses
// func AdminOnly(c *fiber.Ctx) error {
// 	authHeader := c.Get("Authorization")
// 	if authHeader == "" {
// 		return c.Status(fiber.StatusUnauthorized).JSON(map[string]interface{}{
// 			"message": "Authorization header is missing",
// 		})
// 	}

// 	tokenString := authHeader[len("Bearer "):] // Ambil token dari header

// 	// Verifikasi token
// 	_, err := utils.VerifyToken(tokenString)
// 	if err != nil {
// 		return c.Status(fiber.StatusUnauthorized).JSON(map[string]interface{}{
// 			"message": "Invalid or expired token",
// 		})
// 	}

// 	// Lanjutkan ke handler berikutnya
// 	return c.Next()
// }
