package pages

import (
	"fmt"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/t67y110v/web/internal/app/utils"
)

func (h *Pages) ProtocolPage() fiber.Handler {

	return func(c *fiber.Ctx) error {
		_, userName, userCentrerID, userRole, err := utils.CheckToken(c.Cookies("JWT"))
		if err != nil {
			return utils.LoginError(c)
		}

		protocolId := c.Params("id")
		subjectNumber := c.Params("number")
		p_id, err := strconv.Atoi(protocolId)

		if err != nil {
			return utils.ErrorPage(c, err)
		}

		p, err := h.pgStore.Repository().GetProtocolById(p_id)
		if err != nil {
			return utils.ErrorPage(c, err)
		}
		s, err := h.mgStore.Repository().GetSubjectsByProtocolId(p_id)
		if err != nil {
			return utils.ErrorPage(c, err)
		}
		subject, err := h.mgStore.Repository().GetSubjectByNumber(subjectNumber)
		if err != nil {
			return utils.ErrorPage(c, err)
		}

		cName, err := h.pgStore.Repository().GetCenterName(userCentrerID)
		if err != nil {
			return utils.ErrorPage(c, err)
		}
		return c.Render("protocol/protocol_index", fiber.Map{
			"Name":         userName,
			"Role":         utils.GetUserRole(userRole),
			"CLinicCenter": cName,
			"ClinicId":     userCentrerID,
			"Protocol":     p,
			"Subjects":     s,
			"Subject":      subject,
		})

	}

}

func (h *Pages) ProtocolEdit() fiber.Handler {
	return func(c *fiber.Ctx) error {
		_, userName, userCentrerID, userRole, err := utils.CheckToken(c.Cookies("JWT"))
		if err != nil {
			return utils.LoginError(c)
		}

		protocol_id := c.Params("id")
		p_id, err := strconv.Atoi(protocol_id)
		if err != nil {
			return c.Redirect(fmt.Sprintf("/protocol/edit/%s", protocol_id))
		}
		p, err := h.pgStore.Repository().GetProtocolById(p_id)
		if err != nil {
			return c.Redirect("/main/filter=0") // 404
		}
		cName, err := h.pgStore.Repository().GetCenterName(userCentrerID)
		if err != nil {
			return utils.ErrorPage(c, err)
		}
		return c.Render("protocol/protocol_edit", fiber.Map{
			"Name":         userName,
			"Role":         utils.GetUserRole(userRole),
			"CLinicCenter": cName,
			"Protocol":     p,
		})
	}
}

func (h *Pages) ProtocolNew() fiber.Handler {
	return func(c *fiber.Ctx) error {
		_, userName, userCentrerID, userRole, err := utils.CheckToken(c.Cookies("JWT"))
		if err != nil {
			return utils.LoginError(c)
		}
		cName, err := h.pgStore.Repository().GetCenterName(userCentrerID)
		if err != nil {
			return utils.ErrorPage(c, err)
		}
		return c.Render("protocol/protocol_new", fiber.Map{
			"Name":         userName,
			"Role":         utils.GetUserRole(userRole),
			"CLinicCenter": cName,
		})
	}
}
