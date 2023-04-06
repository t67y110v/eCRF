package store

import (
	"strconv"

	protocolModel "github.com/t67y110v/web/internal/app/model/protocol"
	model "github.com/t67y110v/web/internal/app/model/user"
)

type PostgresStoreRepository struct {
	store *Store
}

func (r *PostgresStoreRepository) Create(u *model.User) error {
	if err := u.BeforeCreate(); err != nil {
		return err
	}
	if result := r.store.db.Create(u); result.Error != nil {
		return result.Error
	}

	return nil
}

func (r *PostgresStoreRepository) FindByEmail(email string) (*model.User, error) {
	u := &model.User{}

	if result := r.store.db.Where(model.User{Email: email}).First(&u); result.Error != nil {
		return nil, result.Error
	}

	return u, nil

}

func (r *PostgresStoreRepository) FindByID(ID string) (*model.User, error) {
	u := &model.User{}

	id, err := strconv.Atoi(ID)
	if err != nil {
		return nil, err
	}

	if result := r.store.db.Where(model.User{Id: id}).First(&u); result.Error != nil {
		return nil, result.Error
	}

	return u, nil

}
func (r *PostgresStoreRepository) GetProtocols() ([]protocolModel.Protocol, error) {
	p := []protocolModel.Protocol{}

	if result := r.store.db.Find(&p); result.Error != nil {
		return nil, result.Error
	}

	return p, nil

}
