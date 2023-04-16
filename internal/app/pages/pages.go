package pages

import (
	"github.com/t67y110v/web/internal/app/logging"
	"github.com/t67y110v/web/internal/app/store"
)

type Pages struct {
	logger  logging.Logger
	pgStore store.PostgresStore
	mgStore store.MongoStore
}

func NewHandlers(pgstore store.PostgresStore, mgstore store.MongoStore, logger logging.Logger) *Pages {
	return &Pages{
		pgStore: pgstore,
		mgStore: mgstore,
		logger:  logger,
	}
}
