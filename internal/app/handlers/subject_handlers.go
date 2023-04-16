package handlers

import (
	"fmt"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/t67y110v/web/internal/app/utils"
)

func (h *Handlers) NewSubject() fiber.Handler {
	return func(c *fiber.Ctx) error {
		_, _, _, _, err := utils.CheckToken(c.Cookies("JWT"))
		if err != nil {
			return utils.LoginError(c)
		}

		centerId, err := strconv.Atoi(c.FormValue("center_id"))
		if err != nil {
			return err
		}
		protocolId, err := strconv.Atoi(c.FormValue("protocol_id"))
		if err != nil {
			return err
		}
		if err := h.mgStore.Repository().AddSubject(c.FormValue("number"), c.FormValue("initials"), centerId, protocolId); err != nil {
			return err
		}

		if err != nil {
			return err
		}

		return c.Redirect(fmt.Sprintf("/protocol/%s", c.FormValue("protocol_id")))
	}
}
