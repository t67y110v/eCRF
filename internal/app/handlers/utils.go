package handlers

import (
	"errors"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

func createToken(id, centerId, role int, name string) (string, error) {
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

func checkToken(cookie string) (string, string, int, int, error) {
	if cookie == "" {
		return "", "", 0, 0, errors.New("nil token")
	}
	tokenString := cookie

	claims := jwt.MapClaims{}

	_, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte("11we$*9sd*(@!)"), nil
	})
	if err != nil {
		return "", "", 0, 0, err
	}
	if claims["id"] == nil {
		return "", "", 0, 0, errors.New("nil id from token")
	}

	id := strconv.Itoa(int(claims["id"].(float64)))
	role := (int(claims["role"].(float64)))
	name := claims["name"].(string)

	centerId := (int(claims["centerId"].(float64)))
	return id, name, centerId, role, nil

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

func getUserRole(roleId int) string {
	m := make(map[int]string)
	m[1] = "Администратор"
	m[2] = "Исследователь"
	m[3] = "Главный исследователь"
	m[4] = "Монитор"
	m[5] = "Дата Менеджер"
	m[6] = "Аудитор Контроля Качества"
	m[7] = "Медицинский монитор"
	m[8] = "Специалист по фармаконадзору"
	return m[roleId]
}
