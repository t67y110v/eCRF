package handlers

import (
	"fmt"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/t67y110v/web/internal/app/handlers/requests"
)

// @Summary AdverseEventsSubject
// @Description AdverseEventSubject value of subject
// @Tags         Subject.Adverse
//
//	@Accept       json
//
// @Produce json
// @Param  data body requests.AdverseEvents true  "adverseevents"
// @Success 200 {object} responses.AddProtocol
// @Failure 400 {object} responses.Error
// @Failure 500 {object} responses.Error
// @Router /api/subject/offsiteblock/adverseevent [patch]
func (h *Handlers) AdverseEventsSubject() fiber.Handler {
	return func(c *fiber.Ctx) error {
		req := requests.AdverseEvents{}
		if err := c.BodyParser(&req); err != nil {
			c.Status(http.StatusBadRequest)
			return c.JSON(fiber.Map{
				"message": err.Error(),
			})
		}
		subject, err := h.mgStore.Subject().GetSubjectByNumber(req.SubjectNumber, req.ProtocolID)
		if err != nil {
			c.Status(http.StatusBadRequest)
			return c.JSON(fiber.Map{
				"message": err.Error(),
			})
		}
		if err := h.mgStore.OffSiteBlock().AdverseEvents(c.Context(), subject.ID,
			req.AdverseEventsRegistered, req.DescriptionOfTheAdverseEvent,
			req.DateOfStartAE, req.IsContinuesEnd, req.DateOfEndAE,
			req.Severity, req.RecurringPhenomenon,
			req.AssociationWithTheDrugUsed, req.Foresight,
			req.ConnectionBetweenAEAndDU, req.RenewalAfterUse,
			req.LocalReaction, req.SubjectDropout,
			req.MeasuresTaken, req.MeasuresTakenOnUD,
			req.Exodus, req.IsItSerious,
			req.SeverityCriterion, req.TestImpact,
			req.DoseEffect, req.ImpactOnHospitalStay,
			req.RelationshipWithMedication, req.Expectancy,
		); err != nil {
			c.Status(http.StatusBadRequest)
			return c.JSON(fiber.Map{
				"message": err.Error(),
			})
		}
		go h.journal.SaveAction(c.Context(), fmt.Sprintf("Протокол:%d Субьект:%s, Вневизитный блок - Нежелательные явления, первый ввод данных", req.ProtocolID, req.SubjectNumber), c.Cookies("token_name"), c.Cookies("token_role"), "patch", req)
		return c.JSON(fiber.Map{
			"message": "success",
		})
	}
}
