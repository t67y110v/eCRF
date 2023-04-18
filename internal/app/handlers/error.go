package handlers

import "github.com/gofiber/fiber/v2"

func (h *Handlers) SystemError() fiber.Handler {
	return func(c *fiber.Ctx) error {

		return c.Redirect("/error/")
	}
}
