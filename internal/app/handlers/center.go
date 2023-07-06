package handlers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/t67y110v/web/internal/app/handlers/requests"
)

// @Summary Add new  center
// @Description add new center
// @Tags         Center
//
//	@Accept       json
//
// @Produce json
// @Param  data body requests.AddCenter true  "addcenter"
// @Success 200 {object} responses.DeleteProtocol
// @Failure 400 {object} responses.Error
// @Failure 500 {object} responses.Error
// @Router /api/center/add [post]
func (h *Handlers) AddNewCenter() fiber.Handler {

	return func(c *fiber.Ctx) error {
		req := requests.AddCenter{}
		if err := c.BodyParser(&req); err != nil {
			c.Status(http.StatusBadRequest)
			return c.JSON(fiber.Map{
				"message": err.Error(),
			})
		}
		if err := h.pgStore.Repository().AddNewCenter(req.Name); err != nil {
			c.Status(http.StatusBadRequest)
			return c.JSON(fiber.Map{
				"message": err.Error(),
			})
		}
		go h.journal.SaveAction(c.Context(), fmt.Sprintf("Добавление центра %s", req.Name), c.Cookies("token_name"), c.Cookies("token_role"), "create", req)
		return c.JSON(fiber.Map{
			"message": "success",
		})

	}

}

// @Summary Update center
// @Description update center
// @Tags         Center
//
//	@Accept       json
//
// @Produce json
// @Param  data body requests.UpdateCenter true  "updatecenter"
// @Success 200 {object} responses.DeleteProtocol
// @Failure 400 {object} responses.Error
// @Failure 500 {object} responses.Error
// @Router /api/center/update [patch]
func (h *Handlers) UpdateCenter() fiber.Handler {
	return func(c *fiber.Ctx) error {

		req := requests.UpdateCenter{}
		if err := c.BodyParser(&req); err != nil {
			c.Status(http.StatusBadRequest)
			return c.JSON(fiber.Map{
				"message": err.Error(),
			})
		}
		center, err := h.pgStore.Repository().GetCenterName(req.CenterID)
		if err != nil {
			c.Status(http.StatusBadRequest)
			return c.JSON(fiber.Map{
				"message": err.Error(),
			})
		}
		if err := h.pgStore.Repository().UpdateCenter(req.CenterID, req.Name); err != nil {
			c.Status(http.StatusBadRequest)
			return c.JSON(fiber.Map{
				"message": err.Error(),
			})
		}
		go h.journal.SaveAction(c.Context(), fmt.Sprintf("Обновление центра %s|id:%d", req.Name, req.CenterID), c.Cookies("token_name"), c.Cookies("token_role"), "update", req, center)
		return c.JSON(fiber.Map{
			"message": "success",
		})
	}
}

// @Summary Delete center
// @Description delete center
// @Tags         Center
//
//	@Accept       json
//
// @Produce json
// @Param  data body requests.DeleteCenter true  "deletecenter"
// @Success 200 {object} responses.DeleteProtocol
// @Failure 400 {object} responses.Error
// @Failure 500 {object} responses.Error
// @Router /api/center/delete [delete]
func (h *Handlers) DeleteCenter() fiber.Handler {
	return func(c *fiber.Ctx) error {
		req := requests.DeleteCenter{}
		if err := c.BodyParser(&req); err != nil {
			c.Status(http.StatusBadRequest)
			return c.JSON(fiber.Map{
				"message": err.Error(),
			})
		}
		center, err := h.pgStore.Repository().GetCenterName(req.ID)
		if err != nil {
			c.Status(http.StatusBadRequest)
			return c.JSON(fiber.Map{
				"message": err.Error(),
			})
		}
		if err := h.pgStore.Repository().DeleteCenter(req.ID); err != nil {
			c.Status(http.StatusBadRequest)
			return c.JSON(fiber.Map{
				"message": err.Error(),
			})
		}
		go h.journal.SaveAction(c.Context(), fmt.Sprintf("Удаление центра %s|id:%d", center, req.ID), c.Cookies("token_name"), c.Cookies("token_role"), "delete", req)
		return c.JSON(fiber.Map{
			"message": "success",
		})
	}
}

// @Summary Get all centers
// @Description getting all  centers
// @Tags         Center
//
//	@Accept       json
//
// @Produce json
// @Success 200 {object} responses.Center
// @Failure 400 {object} responses.Error
// @Failure 500 {object} responses.Error
// @Router /api/center/all [get]
func (h *Handlers) GetCenters() fiber.Handler {
	return func(c *fiber.Ctx) error {
		centers, err := h.pgStore.Repository().GetAllCenters()
		if err != nil {
			c.Status(http.StatusBadRequest)
			return c.JSON(fiber.Map{
				"message": err.Error(),
			})
		}
		return c.JSON(centers)
	}
}

// @Summary Get center name
// @Description getting center name by id
// @Tags         Center
//
//	@Accept       json
//
// @Produce json
// @Param id   path      string  true  "ID"
// @Success 200 {object} responses.DeleteProtocol
// @Failure 400 {object} responses.Error
// @Failure 500 {object} responses.Error
// @Router /api/center/name/{id} [get]
func (h *Handlers) GetCenterName() fiber.Handler {
	return func(c *fiber.Ctx) error {
		centerID, _ := strconv.Atoi(c.Params("id"))
		name, err := h.pgStore.Repository().GetCenterName(centerID)
		if err != nil {
			c.Status(http.StatusBadRequest)
			return c.JSON(fiber.Map{
				"message": err.Error(),
			})
		}
		return c.JSON(fiber.Map{
			"message": name,
		})
	}
}
