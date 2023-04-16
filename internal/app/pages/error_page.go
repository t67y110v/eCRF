package pages

import "github.com/gofiber/fiber/v2"

func (h *Pages) ErrorPage() fiber.Handler {
	return func(c *fiber.Ctx) error {
		err := c.Cookies("Error")

		return c.Render("errors/access_error",
			fiber.Map{
				"Error": err,
			},
		)
	}
}
