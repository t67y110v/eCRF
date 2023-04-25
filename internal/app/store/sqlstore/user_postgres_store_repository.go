package store

import (
	"strconv"

	model "github.com/t67y110v/web/internal/app/model/user"
)

type PostgresStoreRepository struct {
	store *Store
}

func (r *PostgresStoreRepository) Create(u *model.User) error {

	if err := u.BeforeCreate(r.store.db); err != nil {
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

func (r *PostgresStoreRepository) GetUsers() ([]model.User, error) {
	u := []model.User{}
	if result := r.store.db.Order("id").Find(&u); result.Error != nil {
		return nil, result.Error
	}
	return u, nil
}

func (r *PostgresStoreRepository) UpdateUser(ID, role, centerId int, email, name, paswword string) error {
	u := model.User{
		Id:       ID,
		CenterID: centerId,
		Email:    email,
		Name:     name,
		Password: paswword,
	}
	if err := u.BeforeCreate(r.store.db); err != nil {
		return err
	}
	if result := r.store.db.Model(model.User{}).Where("id = ?", ID).Updates(u); result.Error != nil {
		return result.Error
	}
	return nil
}

func (r *PostgresStoreRepository) DeleteUser(Id int) error {
	if result := r.store.db.Delete(&model.User{}, Id); result.Error != nil {
		return result.Error

	}
	return nil
}
