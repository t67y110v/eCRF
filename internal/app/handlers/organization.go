package handlers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/t67y110v/web/internal/app/handlers/requests"
)

// @Summary Add new  organization
// @Description add new organization
// @Tags         Organization
//
//	@Accept       json
//
// @Produce json
// @Param  data body requests.AddNewOrganization true  "addorganizatione"
// @Success 200 {object} responses.DeleteProtocol
// @Failure 400 {object} responses.Error
// @Failure 500 {object} responses.Error
// @Router /api/organization/add [post]
func (h *Handlers) AddNewOrganization() fiber.Handler {

	return func(c *fiber.Ctx) error {
		req := requests.AddNewOrganization{}
		if err := c.BodyParser(&req); err != nil {
			c.Status(http.StatusBadRequest)
			return c.JSON(fiber.Map{
				"message": err.Error(),
			})
		}
		if err := h.pgStore.Repository().AddNewOrganization(req.Name); err != nil {
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

// @Summary Update organization
// @Description update organization
// @Tags         Organization
//
//	@Accept       json
//
// @Produce json
// @Param  data body requests.UpdateOrganization true  "organization"
// @Success 200 {object} responses.DeleteProtocol
// @Failure 400 {object} responses.Error
// @Failure 500 {object} responses.Error
// @Router /api/organization/update [patch]
func (h *Handlers) UpdateOrganization() fiber.Handler {
	return func(c *fiber.Ctx) error {

		req := requests.UpdateOrganization{}
		if err := c.BodyParser(&req); err != nil {
			c.Status(http.StatusBadRequest)
			return c.JSON(fiber.Map{
				"message": err.Error(),
			})
		}
		center, err := h.pgStore.Repository().GetOrganizationName(req.CenterID)
		if err != nil {
			c.Status(http.StatusBadRequest)
			return c.JSON(fiber.Map{
				"message": err.Error(),
			})
		}
		if err := h.pgStore.Repository().UpdateOrganization(req.CenterID, req.Name); err != nil {
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

// @Summary Delete organization
// @Description delete organization
// @Tags         Organization
//
//	@Accept       json
//
// @Produce json
// @Param  data body requests.DeleteOrganization true  "organization"
// @Success 200 {object} responses.DeleteProtocol
// @Failure 400 {object} responses.Error
// @Failure 500 {object} responses.Error
// @Router /api/organization/delete [delete]
func (h *Handlers) DeleteOrganization() fiber.Handler {
	return func(c *fiber.Ctx) error {
		req := requests.DeleteOrganization{}
		if err := c.BodyParser(&req); err != nil {
			c.Status(http.StatusBadRequest)
			return c.JSON(fiber.Map{
				"message": err.Error(),
			})
		}
		center, err := h.pgStore.Repository().GetOrganizationName(req.ID)
		if err != nil {
			c.Status(http.StatusBadRequest)
			return c.JSON(fiber.Map{
				"message": err.Error(),
			})
		}
		if err := h.pgStore.Repository().DeleteOrganization(req.ID); err != nil {
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

// @Summary Get all organization
// @Description getting all  organization
// @Tags         Organization
//
//	@Accept       json
//
// @Produce json
// @Success 200 {object} responses.Organization
// @Failure 400 {object} responses.Error
// @Failure 500 {object} responses.Error
// @Router /api/organization/all [get]
func (h *Handlers) GetOrganizations() fiber.Handler {
	return func(c *fiber.Ctx) error {
		centers, err := h.pgStore.Repository().GetAllOrganizations()
		if err != nil {
			c.Status(http.StatusBadRequest)
			return c.JSON(fiber.Map{
				"message": err.Error(),
			})
		}
		return c.JSON(centers)
	}
}

// @Summary Get organization name
// @Description getting organization name by id
// @Tags         Organization
//
//	@Accept       json
//
// @Produce json
// @Param id   path      string  true  "ID"
// @Success 200 {object} responses.DeleteProtocol
// @Failure 400 {object} responses.Error
// @Failure 500 {object} responses.Error
// @Router /api/organization/name/{id} [get]
func (h *Handlers) GetOrganizationName() fiber.Handler {
	return func(c *fiber.Ctx) error {
		centerID, _ := strconv.Atoi(c.Params("id"))
		name, err := h.pgStore.Repository().GetOrganizationName(centerID)
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
