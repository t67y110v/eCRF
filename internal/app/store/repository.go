package store

import (
	postgresModelCenter "github.com/t67y110v/web/internal/app/model/center"
	mongoModel "github.com/t67y110v/web/internal/app/model/product"
	postgresModelProtocol "github.com/t67y110v/web/internal/app/model/protocol"
	posgresModelSubject "github.com/t67y110v/web/internal/app/model/subject"

	postgresModelUser "github.com/t67y110v/web/internal/app/model/user"
)

type PostgresStoreRepository interface {
	Create(*postgresModelUser.User) error
	FindByEmail(string) (*postgresModelUser.User, error)
	FindByID(string) (*postgresModelUser.User, error)
	GetProtocols() ([]postgresModelProtocol.Protocol, error)
	GetProtocolById(ID int) (*postgresModelProtocol.Protocol, error)
	UpdateProtocolById(ID, status int, name string, centerId int) error
	AddProtocol(name string, status, centerID int) error
	GetProtocolsByFilter(filter string, centerId int) ([]postgresModelProtocol.Protocol, error)
	GetSubjects() ([]posgresModelSubject.Subject, error)
	GetCenterName(centerId int) (string, error)
	GetAllCenters() ([]postgresModelCenter.Center, error)
	GetUsers() ([]postgresModelUser.User, error)
	UpdateUser(ID, role, centerId int, email, name, paswword string) error
	AddNewCenter(name string) error
	UpdateCenter(centerId int, name string) error
}

type MongoStoreRepository interface {
	AddProduct(name, category, imgPath, description string, price, discount int) error
	GetAllProducts() ([]*mongoModel.Product, error)
	FilterByCategory(filter string) ([]*mongoModel.Product, error)
	DeleteProduct(value string) error
}
