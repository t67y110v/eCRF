package handlers

import (
	"net/http"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/t67y110v/web/internal/app/handlers/requests"
)

// @Summary Subjects  Get
// @Description Getting all subjects by protocol id
// @Tags         Subject
//
//	@Accept       json
//
// @Produce json
// @Param protocol_id   path      string  true  "protocol_id"
// @Success 200 {object} responses.GetSubject
// @Failure 400 {object} responses.Error
// @Failure 500 {object} responses.Error
// @Router /subject/{protocol_id} [get]
func (h *Handlers) GetSubjects() fiber.Handler {
	return func(c *fiber.Ctx) error {
		protocolID, _ := strconv.Atoi(c.Params("protocol_id"))
		s, err := h.mgStore.Subject().GetSubjectsByProtocolId(protocolID)
		if err != nil {
			c.Status(http.StatusBadRequest)
			return c.JSON(fiber.Map{
				"message": err.Error(),
			})
		}
		return c.JSON(s)
	}
}

// @Summary Subject Add
// @Description creating a new protocol
// @Tags         Subject
//
//	@Accept       json
//
// @Produce json
// @Param  data body requests.AddSubject true  "addsubject"
// @Success 200 {object} responses.AddProtocol
// @Failure 400 {object} responses.Error
// @Failure 500 {object} responses.Error
// @Router /subject/add [post]
func (h *Handlers) AddSubject() fiber.Handler {
	return func(c *fiber.Ctx) error {
		req := requests.AddSubject{}
		if err := c.BodyParser(&req); err != nil {
			c.Status(http.StatusBadRequest)
			return c.JSON(fiber.Map{
				"message": err.Error(),
			})
		}

		if err := h.mgStore.Subject().AddSubject(req.Number, req.Initials, req.CenterId, req.ProtocolId); err != nil {
			c.Status(http.StatusBadRequest)
			return c.JSON(fiber.Map{
				"message": err.Error(),
			})
		}

		return c.JSON(fiber.Map{
			"message": "success",
		})
	}
}
