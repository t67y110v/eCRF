package handlers

import (
	"net/http"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/t67y110v/web/internal/app/handlers/requests"
)

// @Summary Protocol Save
// @Description saving of updating prtocol
// @Tags         Protocol
//
//	@Accept       json
//
// @Produce json
// @Param  data body requests.SaveProtocol true  "saveprotocol"
// @Success 200 {object} responses.SaveProtocol
// @Failure 400 {object} responses.Error
// @Failure 500 {object} responses.Error
// @Router /protocols/save [patch]
func (h *Handlers) SaveProtocol() fiber.Handler {
	return func(c *fiber.Ctx) error {

		req := requests.SaveProtocol{}
		if err := c.BodyParser(&req); err != nil {
			c.Status(http.StatusBadRequest)
			return c.JSON(fiber.Map{
				"message": err.Error(),
			})
		}
		err := h.pgStore.Repository().UpdateProtocolById(req.ID, req.Status, req.Name, req.CenterId)
		if err != nil {
			c.Status(http.StatusBadRequest)
			return c.JSON(fiber.Map{
				"message": err.Error(),
			})
		}

		go h.operations.SaveAction(c.Context(), "SaveProtocol", "200", req.Name, "Обновление протокола")

		return c.JSON(fiber.Map{
			"message": "success",
		})
	}
}

// @Summary Protocol Add
// @Description creating a new protocol
// @Tags         Protocol
//
//	@Accept       json
//
// @Produce json
// @Param  data body requests.AddProtocol true  "addprotocol"
// @Success 200 {object} responses.AddProtocol
// @Failure 400 {object} responses.Error
// @Failure 500 {object} responses.Error
// @Router /protocols/add [post]
func (h *Handlers) AddProtocol() fiber.Handler {
	return func(c *fiber.Ctx) error {
		req := requests.AddProtocol{}
		if err := c.BodyParser(&req); err != nil {
			c.Status(http.StatusBadRequest)
			return c.JSON(fiber.Map{
				"message": err.Error(),
			})
		}
		err := h.pgStore.Repository().AddProtocol(req.Name, req.Status, req.CenterID)
		if err != nil {
			c.Status(http.StatusBadRequest)
			return c.JSON(fiber.Map{
				"message": err.Error(),
			})
		}
		go h.operations.SaveAction(c.Context(), "AddProtocol", "200", req.Name, "Добавление протокола")
		c.Status(http.StatusBadRequest)
		return c.JSON(fiber.Map{
			"message": "success",
		})
	}
}

// @Summary Protocol Delete
// @Description Deleting a  protocol  by id
// @Tags         Protocol
//
//	@Accept       json
//
// @Produce json
// @Param  data body requests.DeleteProtocol true  "deleteprotocol"
// @Success 200 {object} responses.DeleteProtocol
// @Failure 400 {object} responses.Error
// @Failure 500 {object} responses.Error
// @Router /protocols/delete [delete]
func (h *Handlers) DeleteProtocol() fiber.Handler {
	return func(c *fiber.Ctx) error {
		req := requests.DeleteProtocol{}
		if err := c.BodyParser(&req); err != nil {
			c.Status(http.StatusBadRequest)
			return c.JSON(fiber.Map{
				"message": err.Error(),
			})
		}

		if err := h.pgStore.Repository().DeleteProtocol(req.ID); err != nil {
			c.Status(http.StatusBadRequest)
			return c.JSON(fiber.Map{
				"message": err.Error(),
			})
		}

		go h.operations.SaveAction(c.Context(), "DeleteProtocol", "200", "test", "Удаление протокола")

		return c.JSON(fiber.Map{
			"message": "success",
		})
	}
}

// @Summary Protocols Get
// @Description Getting all protocols from users center id and filter params
// @Tags         Protocol
//
//	@Accept       json
//
// @Produce json
// @Param filter   path      string  true  "Filter"
// @Param center   path      string  true  "Center"
// @Success 200 {object} responses.GetProtocols
// @Failure 400 {object} responses.Error
// @Failure 500 {object} responses.Error
// @Router /protocols/{filter}/{center}  [get]
func (h *Handlers) GetProtocols() fiber.Handler {
	return func(c *fiber.Ctx) error {

		filter := c.Params("filter")
		userCenter, _ := strconv.Atoi(c.Params("center"))

		p, err := h.pgStore.Repository().GetProtocolsByFilter(filter, userCenter)
		if err != nil {
			c.Status(http.StatusBadRequest)
			return c.JSON(fiber.Map{
				"message": err.Error(),
			})
		}

		return c.JSON(p)
	}
}
