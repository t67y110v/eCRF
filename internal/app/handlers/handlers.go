package handlers

import (
	"github.com/t67y110v/web/internal/app/journal"
	"github.com/t67y110v/web/internal/app/logging"
	"github.com/t67y110v/web/internal/app/store"
)

type Handlers struct {
	logger  logging.Logger
	pgStore store.PostgresStore
	mgStore store.MongoStore
	journal journal.Journal
}

func NewHandlers(pgstore store.PostgresStore, mgstore store.MongoStore, logger logging.Logger) *Handlers {
	return &Handlers{
		pgStore: pgstore,
		mgStore: mgstore,
		logger:  logger,
		journal: *journal.NewOperations(mgstore),
	}
}
