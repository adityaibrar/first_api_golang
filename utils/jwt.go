package utils

import (
	"first_api_golang/models"
	"fmt"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

func GenerateJWT(user models.User) (string, error) {
	claims := jwt.MapClaims{
		"user_id":  user.ID,
		"username": user.Username,
		"email":    user.Email,
		"exp_id":   time.Now().Add(time.Hour * 24).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET_KEY")))
	return tokenString, err
}

func AuthMiddleware(c *fiber.Ctx) error {
	authHeader := c.Get("Authorization")
	if authHeader == "" {
		return ErrorResponse(c, fiber.StatusBadRequest, "No Token provided")
	}

	var tokenString string
	fmt.Sscanf(authHeader, "Bearer %s", &tokenString)

	if tokenString == "" {
		return ErrorResponse(c, fiber.StatusBadRequest, "Token format invalid")
	}

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("JWT_SECRET_KEY")), nil
	})

	if err != nil || !token.Valid {
		return ErrorResponse(c, fiber.StatusBadRequest, "Token is invalid")
	}
	return c.Next()
}
