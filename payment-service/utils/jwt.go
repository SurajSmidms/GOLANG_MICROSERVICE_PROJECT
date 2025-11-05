package utils

import (
	"errors"
	"os"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// ValidateToken verifies the JWT and returns the username if valid
func ValidateToken(tokenString string) (string, error) {
	secret := os.Getenv("ACCESS_SECRET")
	if secret == "" {
		return "", errors.New("ACCESS_SECRET not set in environment")
	}

	// Remove "Bearer " prefix if present
	if strings.HasPrefix(tokenString, "Bearer ") {
		tokenString = strings.TrimPrefix(tokenString, "Bearer ")
	}

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return []byte(secret), nil
	})

	if err != nil {
		return "", err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		// Optional expiration check
		if exp, ok := claims["exp"].(float64); ok && time.Now().Unix() > int64(exp) {
			return "", errors.New("token expired")
		}

		username, _ := claims["username"].(string)
		return username, nil
	}

	return "", errors.New("invalid token")
}
