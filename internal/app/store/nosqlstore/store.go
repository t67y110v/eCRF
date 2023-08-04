package nosqlstore

import (
	"log"

	"github.com/t67y110v/web/internal/app/store"

	"go.mongodb.org/mongo-driver/mongo"
)

type Store struct {
	client                      *mongo.Client
	mongoStoreRepository        *MongoSubjectRepository
	mongoScreeningRepository    *MongoScreeningRepository
	mongoJournalRepository      *MongoJournalRepository
	mongoColorRepository        *MongoColorRepository
	mongoOffSiteBlockRepository *MongoOffSiteBlockRepository
}

func NewMongoDB(client *mongo.Client) *Store {
	log.Println("MongoDB initialization")

	return &Store{
		client: client,
	}
}

func (s *Store) Subject() store.MongoSubjectRepository {
	if s.mongoStoreRepository != nil {
		return s.mongoStoreRepository
	}
	s.mongoStoreRepository = &MongoSubjectRepository{
		store: s,
	}
	return s.mongoStoreRepository
}

func (s *Store) Screening() store.MongoScreeningRepository {
	if s.mongoScreeningRepository != nil {
		return s.mongoScreeningRepository
	}
	s.mongoScreeningRepository = &MongoScreeningRepository{
		store: s,
	}
	return s.mongoScreeningRepository
}

func (s *Store) Journal() store.MongoJournalRepository {
	if s.mongoJournalRepository != nil {
		return s.mongoJournalRepository
	}

	s.mongoJournalRepository = &MongoJournalRepository{
		store: s,
	}
	return s.mongoJournalRepository
}

func (s *Store) Color() store.MongoColorRepository {
	if s.mongoColorRepository != nil {
		return s.mongoColorRepository
	}
	s.mongoColorRepository = &MongoColorRepository{
		store: s,
	}
	return s.mongoColorRepository
}

func (s *Store) OffSiteBlock() store.MongoOffSiteBlockRepository {
	if s.mongoOffSiteBlockRepository != nil {
		return s.mongoOffSiteBlockRepository
	}
	s.mongoOffSiteBlockRepository = &MongoOffSiteBlockRepository{
		store: s,
	}
	return s.mongoOffSiteBlockRepository
}
