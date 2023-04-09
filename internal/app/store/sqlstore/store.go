package store

import (
	"log"

	"github.com/t67y110v/web/internal/app/store"
	"gorm.io/gorm"
)

type Store struct {
	db                      *gorm.DB
	postgresStoreRepository *PostgresStoreRepository
}

func NewPostgresDB(db *gorm.DB) *Store {
	log.Println("PostgreSQL initialization")

	return &Store{
		db: db,
	}
}

func (s *Store) Repository() store.PostgresStoreRepository {
	if s.postgresStoreRepository != nil {
		return s.postgresStoreRepository
	}

	s.postgresStoreRepository = &PostgresStoreRepository{
		store: s,
	}
	return s.postgresStoreRepository
}
