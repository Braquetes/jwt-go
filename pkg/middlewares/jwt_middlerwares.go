package middlewares

import (
	"fmt"
	"time"

	"github.com/braquetes/jwt-go/pkg/domain"
	"github.com/dgrijalva/jwt-go"
)

func GenerateToken(jwtParams *domain.JWT_Params) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["userId"] = jwtParams.Username
	claims["email"] = jwtParams.Email
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()

	tokenString, err := token.SignedString([]byte("llave-secreta-xd"))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func ValidateToken(tokenString string) (*domain.JWT_Params, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method")
		}
		return []byte("llave-secreta-xd"), nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		userId := claims["userId"].(string)
		email := claims["email"].(string)
		return &domain.JWT_Params{Username: userId, Email: email}, nil
	} else {
		return nil, err
	}
}
