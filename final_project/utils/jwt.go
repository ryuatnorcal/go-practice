package utils

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt"
)

const SECRET_KEY = "secret_meow"

func GenerateToken(email string, userId int64) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email":  email,
		"userId": userId,
		"exp":    time.Now().Add(time.Hour * 2).Unix(),
	})
	return token.SignedString([]byte(SECRET_KEY))
}

func VerifyToken(token string) (int64, error) {
	parsedToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		// .() type checking
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, errors.New("invalid token")
		}
		return []byte(SECRET_KEY), nil
	})
	if err != nil {
		return 0, errors.New("Could not parse token")
	}
	if !parsedToken.Valid {
		return 0, errors.New("Invalid token")
	}
	claims, ok := parsedToken.Claims.(jwt.MapClaims)
	if !ok {
		return 0, errors.New("Invalid token")
	}
	// email := claims["email"].(string)
	// type conversion claims["email"].(string) to int64
	userId := int64(claims["userId"].(float64))

	if int64(claims["exp"].(float64)) < time.Now().Unix() {
		return 0, errors.New("Token expired")
	}
	return userId, nil

}
