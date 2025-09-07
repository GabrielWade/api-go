package jwt

import (
	"os"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func GenerateRefreshToken(userId uint, userName string) (string, error) {
	secret := os.Getenv("SECRET")
	ttlStr := os.Getenv("REFRESH_TOKEN_TTL")
	ttl, err := strconv.Atoi(ttlStr)
	if err != nil {
		return "", err
	}
	expiration := time.Now().Add(time.Duration(ttl) * time.Second).Unix()
	claims := jwt.MapClaims{
		"id":       userId,
		"username": userName,
		"exp":      expiration,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(secret))
}
