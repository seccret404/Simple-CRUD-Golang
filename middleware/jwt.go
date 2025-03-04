package middleware

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

func GenerateToken(userID int) (string, error) {
	secret := os.Getenv("JWT_KEY_SECRET")

	claims := jwt.MapClaims{
		"user_id": userID,
		"exp": time.Now().Add(time.Hour * 72 ).Unix(), //set batas expire
	}

	token := jwt.NewWithClaims(jwt.SigningMethodES256, claims)
	return token.SignedString([] byte(secret))

}

func JwtMiddleware(c *fiber.Ctx) error {
	// Ambil token dari header Authorization
	tokenString := c.Get("Authorization")
	if tokenString == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Token belum disertakan"})
	}

	// Pastikan token dalam format "Bearer <token>"
	parts := strings.Split(tokenString, " ")
	if len(parts) != 2 || parts[0] != "Bearer" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Format token salah, gunakan 'Bearer <token>'"})
	}
	tokenString = parts[1]

	// Ambil secret key dari environment variable
	secret := os.Getenv("JWT_KEY_SECRET")
	if secret == "" {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "JWT Secret tidak ditemukan di env"})
	}

	// Parse token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Pastikan metode signing adalah HMAC
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Metode signing token tidak valid")
		}
		return []byte(secret), nil
	})

	// Jika terjadi error saat parsing token atau token tidak valid
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Token tidak valid: " + err.Error()})
	}

	// Periksa apakah token benar-benar valid
	if !token.Valid {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Token sudah tidak berlaku"})
	}

	// Lanjut ke handler berikutnya jika token valid
	return c.Next()
}