package utils

import (
	"errors"
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

func CheckToken(tokenString string) (string, string, int, int, error) {
	if tokenString == "" {
		return "", "", 0, 0, errors.New("nil token")
	}

	claims := jwt.MapClaims{}

	_, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte("11we$*9sd*(@!)"), nil
	})

	if claims["id"] == nil || err != nil {
		return "", "", 0, 0, errors.New("nil id from token")
	}

	id := strconv.Itoa(int(claims["id"].(float64)))
	role := (int(claims["role"].(float64)))
	name := claims["name"].(string)

	centerId := (int(claims["centerId"].(float64)))
	return id, name, centerId, role, nil

}
