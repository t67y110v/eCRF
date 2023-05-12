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
		if err := h.mgStore.Repository().AddSubject(c.FormValue("number"), c.FormValue("initials"), centerId, protocolId); err != nil {
			return utils.ErrorPage(c, err)
		}

		return c.Redirect(fmt.Sprintf("/protocol/%s/1", c.FormValue("protocol_id")))
	}
}

func (h *Handlers) DemographySubject() fiber.Handler {
	return func(c *fiber.Ctx) error {
		subjectNumber := c.FormValue("subject_number")
		subject, err := h.mgStore.Repository().GetSubjectByNumber(subjectNumber)
		if err != nil {
			fmt.Println("- ", subjectNumber)
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

		if err := h.mgStore.Repository().Demography(c.Context(), subject.ID, sexSubject, raceSubject, birthDateSubject); err != nil {
			h.logger.Warningln(err)
			return utils.ErrorPage(c, err)
		}
		return c.Redirect("/protocol/1/1")
	}
}

func (h *Handlers) DeleteSubject() fiber.Handler {
	return func(c *fiber.Ctx) error {
		if err := h.mgStore.Repository().DeleteSubject(c.FormValue("number")); err != nil {
			return utils.ErrorPage(c, err)
		}
		return c.Redirect(fmt.Sprintf("/protocol/%s/1", c.FormValue("protocol_id")))
	}
}
