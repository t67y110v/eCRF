package handlers

import (
	"errors"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

func createToken(id int) (string, error) {
	secret := "11we$*9sd*(@!)"

	minutesCount, _ := strconv.Atoi("15")

	claims := jwt.MapClaims{}

	claims["exp"] = time.Now().Add(time.Minute * time.Duration(minutesCount)).Unix()

	claims["id"] = id

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(secret))
}

func checkToken(cookie string) (string, error) {
	if cookie == "" {
		return "", errors.New("nil token")
	}
	tokenString := cookie

	claims := jwt.MapClaims{}

	_, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte("11we$*9sd*(@!)"), nil
	})
	if err != nil {
		return "", err
	}
	if claims["id"] == nil {
		return "", errors.New("nil id from token")
	}

	id := strconv.Itoa(int(claims["id"].(float64)))

	return id, nil

}

func loginError(c *fiber.Ctx) error {
	c.Status(fiber.StatusNotFound)
	er := &fiber.Cookie{
		Name:  "error",
		Value: "login",
	}
	c.Cookie(er)
	return c.RedirectBack("/auth")
}

func passwordError(c *fiber.Ctx) error {
	c.Status(fiber.StatusNotFound)
	er := &fiber.Cookie{
		Name:  "error",
		Value: "password",
	}
	c.Cookie(er)
	return c.RedirectBack("/auth")
}
