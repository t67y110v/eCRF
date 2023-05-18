package pages

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	model "github.com/t67y110v/web/internal/app/model/user"
	"github.com/t67y110v/web/internal/app/utils"
)

func (h *Pages) ProtocolPage() fiber.Handler {

	return func(c *fiber.Ctx) error {
		user := c.Locals("user").(*model.User)

		protocolId := c.Params("id")
		subjectNumber := c.Params("number")
		p_id, err := strconv.Atoi(protocolId)

		if err != nil {
			h.logger.Warningln(err)
			return utils.ErrorPage(c, err)
		}
		cName, err := h.pgStore.Repository().GetCenterName(user.CenterID)
		if err != nil {
			h.logger.Warningln(err)
			return utils.ErrorPage(c, err)
		}
		p, err := h.pgStore.Repository().GetProtocolById(p_id)
		if err != nil {
			h.logger.Warningln(err)
			return utils.ErrorPage(c, err)
		}
		s, err := h.mgStore.Subject().GetSubjectsByProtocolId(p_id)
		if err != nil {
			h.logger.Warningln(err)
			return utils.ErrorPage(c, err)
		}
		subject, err := h.mgStore.Subject().GetSubjectByNumber(subjectNumber, p_id)
		if err != nil {
			h.logger.Warningln(err)
			if len(s) == 0 {
				return c.Render("protocol/protocol_index", fiber.Map{
					"Name":         user.Name,
					"Role":         utils.GetUserRole(user.Role),
					"CLinicCenter": cName,
					"ClinicId":     user.CenterID,
					"Protocol":     p,
					"Subjects":     s,
				})
			}
			return utils.ErrorPage(c, err)
		}

		return c.Render("protocol/protocol_index", fiber.Map{
			"Name":                user.Name,
			"Role":                utils.GetUserRole(user.Role),
			"CLinicCenter":        cName,
			"ClinicId":            user.CenterID,
			"Protocol":            p,
			"Subjects":            s,
			"Subject":             subject,
			"ConsentErrors":       utils.GetInformationConsentErrors(subject),
			"DemographyErrors":    utils.GetDemographyErrorsErrors(subject),
			"AnthropometryErrors": utils.GetAnthropometryErrors(subject),
		})

	}

}
