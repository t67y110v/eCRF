package store

type PostgresStore interface {
	Repository() PostgresStoreRepository
}

type MongoStore interface {
	Subject() MongoSubjectRepository
	Screening() MongoScreeningRepository
	Journal() MongoJournalRepository
	Color() MongoColorRepository
	OffSiteBlock() MongoOffSiteBlockRepository
}
