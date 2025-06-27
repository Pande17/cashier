package middleware

// import (
// 	"cashier-machine/utils" // Pastikan utils diimpor untuk menggunakan JWTSecret
// 	"log"
// 	"strings"
// 	"time"

// 	"github.com/dgrijalva/jwt-go"
// 	"github.com/gofiber/fiber/v2"
// )

// // GenerateToken membuat JWT baru dengan payload yang berisi informasi user
// func GenerateToken(userID string, isAdmin bool) (string, error) {
// 	// Set the expiration time for the token (e.g., 1 hour)
// 	expirationTime := time.Now().Add(1 * time.Hour)
// 	claims := &jwt.MapClaims{
// 		"id":       userID,
// 		"is_admin": isAdmin,
// 		"exp":      expirationTime.Unix(),
// 	}

// 	// Create the token with the claims and signing method
// 	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

// 	// Sign the token with the JWT secret
// 	tokenString, err := token.SignedString([]byte(JWTSecret))
// 	if err != nil {
// 		log.Println("Error signing token:", err)
// 		return "", err
// 	}

// 	return tokenString, nil
// }

// // VerifyToken memverifikasi JWT menggunakan JWTSecret
// func VerifyToken(tokenString string) (*jwt.Token, error) {
// 	// Parsing the token
// 	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
// 		// Validasi metode tanda tangan
// 		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
// 			return nil, err
// 		}
// 		return []byte(JWTSecret), nil
// 	})
// 	if err != nil {
// 		log.Println("Error verifying token:", err)
// 		return nil, err
// 	}

// 	return token, nil
// }

// // Middleware untuk memastikan hanya admin yang bisa mengakses route
// func AdminOnly(c *fiber.Ctx) error {
// 	// Mengambil Authorization header dari request
// 	authHeader := c.Get("Authorization")
// 	if authHeader == "" {
// 		return c.Status(fiber.StatusUnauthorized).JSON(map[string]interface{}{
// 			"message": "Authorization header is missing",
// 		})
// 	}

// 	// Memisahkan "Bearer token" dari header
// 	tokenString := strings.TrimPrefix(authHeader, "Bearer ")
// 	if tokenString == authHeader {
// 		// Jika tidak ada token setelah "Bearer", maka token tidak valid
// 		return c.Status(fiber.StatusUnauthorized).JSON(map[string]interface{}{
// 			"message": "Invalid token format",
// 		})
// 	}

// 	// Mengverifikasi dan mendekode token JWT
// 	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
// 		// Pastikan token menggunakan algoritma HS256
// 		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
// 			return nil, fiber.ErrUnauthorized
// 		}
// 		return []byte(utils.KUNCIRAHASIA), nil // Ambil secret key dari utilitas
// 	})

// 	if err != nil || !token.Valid {
// 		return c.Status(fiber.StatusUnauthorized).JSON(map[string]interface{}{
// 			"message": "Invalid or expired token",
// 		})
// 	}

// 	// Mengambil klaim token untuk memverifikasi apakah user adalah admin
// 	claims, ok := token.Claims.(jwt.MapClaims)
// 	if !ok || !claims["is_admin"].(bool) { // Pastikan klaim "is_admin" ada dan true
// 		return c.Status(fiber.StatusForbidden).JSON(map[string]interface{}{
// 			"message": "Access forbidden: only admins are allowed",
// 		})
// 	}

// 	// Mengambil user ID atau informasi lain dari klaim token dan menambahkannya ke context
// 	c.Locals("admin_id", claims["id"])

// 	// Lanjutkan ke handler berikutnya jika user adalah admin
// 	return c.Next()
// }
