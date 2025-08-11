package utils

import (
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var secretKey = "fcbarcelona"

func GenerateToken(email string, id int64) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email":  email,
		"userId": id,
		"exp":    time.Now().Add(time.Hour * 2).Unix(),
	})

	return token.SignedString([]byte(secretKey))
}

func VerifyToken(tokenString string) (int64, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(secretKey), nil
	})

	if err != nil || !token.Valid {
		return 0, errors.New("invalid token")
	}

	claims := token.Claims.(jwt.MapClaims)
	fmt.Printf("Claims: %#v\n", claims)
	uid, ok := claims["userId"].(float64)
	if !ok {
		return 0, errors.New("invalid user id")
	}
	return int64(uid), nil
}
