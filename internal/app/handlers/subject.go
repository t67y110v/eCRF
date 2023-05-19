package handlers

import (
	"fmt"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/t67y110v/web/internal/app/utils"
)

func (h *Handlers) NewSubject() fiber.Handler {
	return func(c *fiber.Ctx) error {

		centerId, err := strconv.Atoi(c.FormValue("center_id"))
		if err != nil {
			return utils.ErrorPage(c, err)
		}
		protocolId, err := strconv.Atoi(c.FormValue("protocol_id"))
		if err != nil {
			return utils.ErrorPage(c, err)
		}
		if err := h.mgStore.Subject().AddSubject(c.FormValue("number"), c.FormValue("initials"), centerId, protocolId); err != nil {
			return utils.ErrorPage(c, err)
		}

		return c.Redirect(fmt.Sprintf("/protocol/%s/%s", c.FormValue("protocol_id"), c.FormValue("number")))
	}
}

func (h *Handlers) DeleteSubject() fiber.Handler {
	return func(c *fiber.Ctx) error {
		if err := h.mgStore.Subject().DeleteSubject(c.FormValue("number")); err != nil {
			return utils.ErrorPage(c, err)
		}
		return c.Redirect(fmt.Sprintf("/protocol/%s/1", c.FormValue("protocol_id")))
	}
}
