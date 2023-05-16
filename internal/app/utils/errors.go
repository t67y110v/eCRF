package utils

import (
	"github.com/gofiber/fiber/v2"
)

func ErrorPage(c *fiber.Ctx, err error) error {
	c.Status(fiber.StatusNotFound)
	er := &fiber.Cookie{
		Name:  "Error",
		Value: err.Error(),
	}
	c.Cookie(er)
	return c.Redirect("/error/")
}

func LoginError(c *fiber.Ctx) error {
	c.Status(fiber.StatusNotFound)
	er := &fiber.Cookie{
		Name:  "error",
		Value: "login",
	}
	c.Cookie(er)
	return c.RedirectBack("/auth")
}

func PasswordError(c *fiber.Ctx) error {
	c.Status(fiber.StatusNotFound)
	er := &fiber.Cookie{
		Name:  "error",
		Value: "password",
	}
	c.Cookie(er)
	return c.RedirectBack("/auth")
}
