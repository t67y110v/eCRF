package handlers

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/t67y110v/web/internal/app/utils"
)

func (h *Handlers) InformaionConsentSubject() fiber.Handler {
	return func(c *fiber.Ctx) error {
		subjectNumber := c.FormValue("subject_number")
		// TODO: add hidden input with current protocol number to correct redirect

		subject, err := h.mgStore.Subject().GetSubjectByNumber(subjectNumber)
		if err != nil {

			h.logger.Warningln(err)
			return utils.ErrorPage(c, err)
		}
		if err := h.mgStore.Screening().InformaionConsent(c.Context(), subject.ID, 1, "false", false, false, false); err != nil {
			h.logger.Warningln(err)
			return utils.ErrorPage(c, err)
		}

		return c.Redirect("/protocol/1/1")
	}
}

func (h *Handlers) DemographySubject() fiber.Handler {
	return func(c *fiber.Ctx) error {
		subjectNumber := c.FormValue("subject_number")
		// TODO: add hidden input with current protocol number to correct redirect

		subject, err := h.mgStore.Subject().GetSubjectByNumber(subjectNumber)
		if err != nil {

			h.logger.Warningln(err)
			return utils.ErrorPage(c, err)
		}
		sexSubject, err := strconv.Atoi(c.FormValue("sex"))
		if err != nil {
			h.logger.Warningln(err)
			return utils.ErrorPage(c, err)
		}
		raceSubject, err := strconv.Atoi(c.FormValue("race"))
		if err != nil {
			h.logger.Warningln(err)
			return utils.ErrorPage(c, err)
		}
		birthDateSubject := c.FormValue("birth_date")

		if err := h.mgStore.Screening().Demography(c.Context(), subject.ID, sexSubject, raceSubject, birthDateSubject); err != nil {
			h.logger.Warningln(err)
			return utils.ErrorPage(c, err)
		}
		return c.Redirect("/protocol/1/1")
	}
}

func (h *Handlers) AnthropometrySubject() fiber.Handler {
	return func(c *fiber.Ctx) error {
		subjectNumber := c.FormValue("subject_number")
		// TODO: add hidden input with current protocol number to correct redirect

		subject, err := h.mgStore.Subject().GetSubjectByNumber(subjectNumber)
		if err != nil {

			h.logger.Warningln(err)
			return utils.ErrorPage(c, err)
		}
		if err := h.mgStore.Screening().Anthropometry(c.Context(), subject.ID, false, "", "", 1, 2, 2); err != nil {
			h.logger.Warningln(err)
			return utils.ErrorPage(c, err)
		}

		return c.Redirect("/protocol/1/1")

	}
}

func (h *Handlers) InclusionCriteriaSubject() fiber.Handler {
	return func(c *fiber.Ctx) error {
		subjectNumber := c.FormValue("subject_number")
		// TODO: add hidden input with current protocol number to correct redirect

		subject, err := h.mgStore.Subject().GetSubjectByNumber(subjectNumber)
		if err != nil {

			h.logger.Warningln(err)
			return utils.ErrorPage(c, err)
		}
		if err := h.mgStore.Screening().InclusionCriteria(c.Context(), subject.ID, false, false, false, false, false, false, false, false, false, false, false, false, false); err != nil {
			h.logger.Warningln(err)
			return utils.ErrorPage(c, err)
		}

		return c.Redirect("/protocol/1/1")

	}
}

func (h *Handlers) ExclusionСriteriaSubject() fiber.Handler {
	return func(c *fiber.Ctx) error {
		subjectNumber := c.FormValue("subject_number")
		// TODO: add hidden input with current protocol number to correct redirect

		subject, err := h.mgStore.Subject().GetSubjectByNumber(subjectNumber)
		if err != nil {

			h.logger.Warningln(err)
			return utils.ErrorPage(c, err)
		}
		if err := h.mgStore.Screening().ExclusionСriteria(c.Context(), subject.ID, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false); err != nil {
			h.logger.Warningln(err)
			return utils.ErrorPage(c, err)
		}

		return c.Redirect("/protocol/1/1")

	}
}
