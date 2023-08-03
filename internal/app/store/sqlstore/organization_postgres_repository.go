package store

import (
	"github.com/t67y110v/web/internal/app/model/organization"
)

func (r *PostgresStoreRepository) GetOrganizationName(organizationID int) (string, error) {
	o := organization.Organization{}

	if result := r.store.db.Where("id = ?", organizationID).First(&o); result.Error != nil {

		return "", result.Error
	}
	return o.Name, nil
}

func (r *PostgresStoreRepository) GetAllOrganizations() ([]organization.Organization, error) {
	o := []organization.Organization{}
	if result := r.store.db.Order("id").Find(&o); result.Error != nil {
		return nil, result.Error
	}
	return o, nil
}

func (r *PostgresStoreRepository) AddNewOrganization(name string) error {
	o := organization.Organization{
		Name: name,
	}
	if result := r.store.db.Create(&o); result.Error != nil {
		return result.Error
	}
	return nil
}

func (r *PostgresStoreRepository) UpdateOrganization(organizationID int, name string) error {
	o := organization.Organization{
		Name: name,
	}
	if result := r.store.db.Model(&organization.Organization{}).Where("id = ?", organizationID).Updates(o); result.Error != nil {
		return result.Error
	}
	return nil
}

func (r *PostgresStoreRepository) DeleteOrganization(organizationID int) error {
	if result := r.store.db.Delete(&organization.Organization{}, organizationID); result.Error != nil {
		return result.Error
	}
	return nil
}
