package store

type PostgresStore interface {
	Repository() PostgresStoreRepository
}

type MongoStore interface {
	Repository() MongoStoreRepository
}
