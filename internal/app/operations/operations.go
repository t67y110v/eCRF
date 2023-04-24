package operations

import "github.com/t67y110v/web/internal/app/store"

type Operations struct {
	mgStore store.MongoStore
}

func NewOperations(mgstore store.MongoStore) *Operations {
	return &Operations{
		mgStore: mgstore,
	}
}
