package utils

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var jwtKey = []byte(os.Getenv("JWT_SECRET"))

func Generate_Token(email string) (string, error) {
	claims := &jwt.MapClaims{
		"email": email,
		"exp":   time.Now().Add(15 * time.Minute).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtKey)
}

func Parse_Token(tokenStr string) (*jwt.MapClaims, error) {
	token, err := jwt.ParseWithClaims(tokenStr, &jwt.MapClaims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})

	if claims, ok := token.Claims.(*jwt.MapClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, err
}
