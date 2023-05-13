package nosqlstore

import (
	"log"

	"github.com/t67y110v/web/internal/app/store"

	"go.mongodb.org/mongo-driver/mongo"
)

type Store struct {
	client                   *mongo.Client
	mongoStoreRepository     *MongoSubjectRepository
	mongoScreeningRepository *MongoScreeningRepository
	mongoJournalRepository   *MongoJournalRepository
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
