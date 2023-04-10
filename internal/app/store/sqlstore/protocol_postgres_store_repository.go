package store

import (
	model "github.com/t67y110v/web/internal/app/model/protocol"
)

func (r *PostgresStoreRepository) GetProtocols() ([]model.Protocol, error) {
	p := []model.Protocol{}

	if result := r.store.db.Order("id").Find(&p); result.Error != nil {
		return nil, result.Error
	}

	return p, nil

}

func (r *PostgresStoreRepository) GetProtocolsByFilter(filter string) ([]model.Protocol, error) {
	p := []model.Protocol{}
	var f string
	if filter == "=1" {
		//name
		f = "name"
	} else if filter == "=2" {
		//status
		f = "status"
	} else if filter == "=3" {
		//center
		f = "center_id"
	} else {
		if result := r.store.db.Order("id").Find(&p); result.Error != nil {
			return nil, result.Error
		}

		return p, nil
	}
	if result := r.store.db.Order(f).Find(&p); result.Error != nil {
		return nil, result.Error
	}

	return p, nil

}

func (r *PostgresStoreRepository) GetProtocolById(ID int) (*model.Protocol, error) {
	p := &model.Protocol{}
	if result := r.store.db.First(&model.Protocol{Id: ID}).Scan(p); result.Error != nil {
		return nil, result.Error

	}
	return p, nil
}

func (r *PostgresStoreRepository) UpdateProtocolById(ID, status int, name string, centerId int) error {
	if result := r.store.db.Model(model.Protocol{}).Where("id = ?", ID).Updates(model.Protocol{Name: name, Status: status, CenterId: centerId}); result.Error != nil {
		return result.Error
	}
	return nil

}

func (r *PostgresStoreRepository) AddProtocol(name string, status, centerID int) error {

	p := &model.Protocol{
		Name:     name,
		Status:   status,
		CenterId: centerID,
	}
	if result := r.store.db.Create(p); result.Error != nil {
		return result.Error
	}

	return nil
}