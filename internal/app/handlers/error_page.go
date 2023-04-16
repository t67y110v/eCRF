package handlers

import "github.com/gofiber/fiber/v2"

func (h *Handlers) ErrorPage() fiber.Handler {
	return func(c *fiber.Ctx) error {
		err := c.Cookies("Error")

		return c.Render("errors/access_error",
			fiber.Map{
				"Error": err,
			},
		)
	}
}
