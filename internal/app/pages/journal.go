package pages

import (
	"github.com/gofiber/fiber/v2"
	"github.com/t67y110v/web/internal/app/utils"
)

func (p *Pages) JournalPage() fiber.Handler {
	return func(c *fiber.Ctx) error {

		actions, err := p.mgStore.Repository().GetActions()
		if err != nil {
			return utils.ErrorPage(c, err)
		}

		return c.Render("journal/journal", fiber.Map{

			"Actions": actions,
		})

	}
}
