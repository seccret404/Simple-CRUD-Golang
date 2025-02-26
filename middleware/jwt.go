package middleware

import (
	"os"
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

func JwtMiddleware(c *fiber.Ctx) error{
	tokenString := c.Get("Authorization")

	if tokenString == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error" : "Token nya belum ade bro.."})
	}

	secret := os.Getenv("JWT_KEY_SECRET")
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error){
		return []byte(secret), nil
	})

	if err != nil || !token.Valid {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error" : "token invalid"})
	}

	return c.Next()

}