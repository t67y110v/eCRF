package store

import (
	model "github.com/t67y110v/web/internal/app/model/center"
)

func (r *PostgresStoreRepository) GetCentersByOrganization(organizationID int) ([]model.Center, error) {
	c := []model.Center{}
	if result := r.store.db.Model(model.Center{}).Where("organization_id = ?", organizationID).Find(&c); result.Error != nil {

		return nil, result.Error
	}
	return c, nil
}

func (r *PostgresStoreRepository) GetCenterName(centerId int) (string, error) {
	c := model.Center{}
	// if result := r.store.db.First(&model.Center{CenterId: centerId}).Scan(c); result.Error != nil {
	// 	return "", result.Error

	// }
	if result := r.store.db.Where("center_id = ?", centerId).First(&c); result.Error != nil {

		return "", result.Error
	}
	return c.Name, nil
}

func (r *PostgresStoreRepository) GetAllCenters() ([]model.Center, error) {
	c := []model.Center{}
	if result := r.store.db.Order("center_id").Find(&c); result.Error != nil {
		return nil, result.Error
	}
	return c, nil
}

func (r *PostgresStoreRepository) AddNewCenter(name string, organizationID int) error {
	c := model.Center{
		OrganizationID: organizationID,
		Name:           name,
	}
	if result := r.store.db.Create(&c); result.Error != nil {
		return result.Error
	}
	return nil
}

func (r *PostgresStoreRepository) UpdateCenter(centerId, orgainzationID int, name string) error {
	c := model.Center{
		Name:           name,
		OrganizationID: orgainzationID,
	}
	if result := r.store.db.Model(model.Center{}).Where("center_id = ?", centerId).Updates(c); result.Error != nil {
		return result.Error
	}
	return nil
}

func (r *PostgresStoreRepository) DeleteCenter(centerId int) error {
	if result := r.store.db.Delete(&model.Center{}, centerId); result.Error != nil {
		return result.Error
	}
	return nil
}
