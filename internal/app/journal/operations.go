package journal

import "github.com/t67y110v/web/internal/app/store"

type Journal struct {
	mgStore store.MongoStore
}

func NewOperations(mgstore store.MongoStore) *Journal {
	return &Journal{
		mgStore: mgstore,
	}
}
