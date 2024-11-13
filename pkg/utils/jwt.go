package utils

import (
	"fmt"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type Claims struct {
	Username string `json:"username"`
	UserID   int    `json:"userID"`
	IsAdmin  bool   `json:"is_admin"`
	jwt.StandardClaims
}

func GenerateToken(username string, userID int, isAdmin bool) (string, error) {
	expirationTime := time.Now().Add(10 * time.Hour)
	claims := &Claims{
		Username: username,
		UserID:   userID,
		IsAdmin:  isAdmin,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	fmt.Println("token : ", token)
	return token.SignedString([]byte(os.Getenv("JWT_SECRET")))
}

func ValidateToken(tokenString string) (*Claims, error) {
	claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWT_SECRET")), nil
	})

	if err != nil || !token.Valid {
		return nil, err
	}
	return claims, nil
}
