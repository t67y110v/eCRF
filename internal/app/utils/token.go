package utils

import (
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

func CreateToken(id, centerId, role int, name string) (string, error) {
	secret := "11we$*9sd*(@!)"

	minutesCount, _ := strconv.Atoi("15")

	claims := jwt.MapClaims{}

	claims["exp"] = time.Now().Add(time.Minute * time.Duration(minutesCount)).Unix()

	claims["id"] = id
	claims["name"] = name
	claims["centerId"] = centerId
	claims["role"] = role

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(secret))
}
