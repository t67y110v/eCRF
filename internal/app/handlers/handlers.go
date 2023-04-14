package handlers

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/t67y110v/web/internal/app/logging"
	"github.com/t67y110v/web/internal/app/store"
)

type Handlers struct {
	logger  logging.Logger
	pgStore store.PostgresStore
	mgStore store.MongoStore
}

func NewHandlers(pgstore store.PostgresStore, mgstore store.MongoStore, logger logging.Logger) *Handlers {
	return &Handlers{
		pgStore: pgstore,
		mgStore: mgstore,
		logger:  logger,
	}
}

func (h *Handlers) Test() fiber.Handler {
	return func(c *fiber.Ctx) error {
		fmt.Println(c.Params("counter"))
		fmt.Println(c.Params("subject_id"))
		return c.JSON(fiber.Map{
			"counter":    c.Params("counter"),
			"subject_id": c.Params("subject_id"),
		})
	}
}
