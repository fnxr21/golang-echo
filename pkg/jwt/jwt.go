package jwtToken

import (
	"fmt"
	"os"

	"github.com/golang-jwt/jwt/v4"
)

func GenerateToken(claims *jwt.MapClaims) (string, error) {
	var SecretKey = os.Getenv("SECRET_KEY")
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	webtoken, err := token.SignedString([]byte(SecretKey))

	if err != nil {
		return "", err
	}

	return webtoken, nil
}

func DecodeToken(tokenString string) (jwt.MapClaims, error) {
	token, err := VerifyToken(tokenString)
	if err != nil {
		return nil, err
	}
	claims, isOk := token.Claims.(jwt.MapClaims)
	if isOk && token.Valid {
		return claims, nil
	}
	return nil, fmt.Errorf("invalid token")
}

func VerifyToken(tokenString string) (*jwt.Token, error) {
	var SecretKey = os.Getenv("SECRET_KEY")

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, isValid := token.Method.(*jwt.SigningMethodHMAC); !isValid {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(SecretKey), nil
	})

	if err != nil {
		return nil, err
	}
	return token, nil
}



