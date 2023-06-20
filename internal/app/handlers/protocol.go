package handlers

import (
	"fmt"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/t67y110v/web/internal/app/utils"
)

func (h *Handlers) SaveProtocol() fiber.Handler {
	return func(c *fiber.Ctx) error {
		name := c.FormValue("name")
		status := c.FormValue("status")
		centerId := c.FormValue("center")
		protocolId := c.FormValue("id")
		p_id, err := strconv.Atoi(protocolId)
		if err != nil {
			return c.Redirect(fmt.Sprintf("/protocol/edit/%s", protocolId))
		}
		s, err := strconv.Atoi(status)
		if err != nil {
			return c.Redirect(fmt.Sprintf("/protocol/edit/%s", protocolId))
		}
		c_id, err := strconv.Atoi(centerId)
		if err != nil {
			return c.Redirect(fmt.Sprintf("/protocol/edit/%s", protocolId))
		}

		err = h.pgStore.Repository().UpdateProtocolById(p_id, s, name, c_id)
		if err != nil {
			return c.Redirect(fmt.Sprintf("/protocol/edit/%s", protocolId))
		}

		go h.operations.SaveAction(c.Context(), "SaveProtocol", "200", c.Locals("name").(string), "Обновление протокола")

		return c.Redirect("/main/filter=0")
	}
}

func (h *Handlers) AddProtocol() fiber.Handler {
	return func(c *fiber.Ctx) error {
		name := c.FormValue("name")
		status := c.FormValue("status")
		centerId := c.FormValue("center")

		s, err := strconv.Atoi(status)
		if err != nil {
			return utils.ErrorPage(c, err)
		}
		c_id, err := strconv.Atoi(centerId)
		if err != nil {
			return utils.ErrorPage(c, err)
		}
		err = h.pgStore.Repository().AddProtocol(name, s, c_id)
		if err != nil {
			return utils.ErrorPage(c, err)
		}
		go h.operations.SaveAction(c.Context(), "AddProtocol", "200", c.Locals("name").(string), "Добавление протокола")
		return c.Redirect("/main/filter=0")
	}
}

func (h *Handlers) DeleteProtocol() fiber.Handler {
	return func(c *fiber.Ctx) error {

		protocolId, err := strconv.Atoi(c.FormValue("id"))
		if err != nil {

			return utils.ErrorPage(c, err)
		}

		if err := h.pgStore.Repository().DeleteProtocol(protocolId); err != nil {
			return utils.ErrorPage(c, err)
		}

		go h.operations.SaveAction(c.Context(), "DeleteProtocol", "200", c.Locals("name").(string), "Удаление протокола")

		return c.Redirect("/main/filter=0")
	}
}
