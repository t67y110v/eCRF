package handlers

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
)

// @Summary Journal Get
// @Description gettning jouranl of operations
// @Tags         Journal
//
//	@Accept       json
//
// @Produce json
// @Success 200 {object} responses.GetJournal
// @Failure 400 {object} responses.Error
// @Failure 500 {object} responses.Error
// @Router /api/journal/all [get]
func (h *Handlers) GetJournal() fiber.Handler {
	return func(c *fiber.Ctx) error {
		journal, err := h.mgStore.Journal().GetActions()
		if err != nil {
			c.Status(http.StatusBadRequest)
			return c.JSON(fiber.Map{
				"message": err.Error(),
			})
		}
		return c.JSON(journal)
	}
}
