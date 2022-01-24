package jwt

import (
	"os"
	"time"

	// external packages
	"github.com/golang-jwt/jwt"
	"github.com/pkg/errors"
	// project packages
)

var SecretKey = []byte(os.Getenv("jwt_secret"))

type CustomClaims struct {
	ID uint `json:"id"`
	jwt.StandardClaims
}

func GenerateToken(userId uint) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["id"] = userId
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()
	tokenString, err := token.SignedString(SecretKey)
	if err != nil {
		return "", errors.Wrap(err, "Error in Generating key")
	}
	return tokenString, nil
}

func ParseToken(tokenStr string) (uint, error) {
	token, err := jwt.ParseWithClaims(tokenStr, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return SecretKey, nil
	})
	if err != nil {
		return 0, err
	}

	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		return claims.ID, nil
	} else {
		return 0, err
	}
}
