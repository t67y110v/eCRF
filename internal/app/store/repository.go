package store

import (
	mongoModel "github.com/t67y110v/web/internal/app/model/product"
	postgresModelProtocol "github.com/t67y110v/web/internal/app/model/protocol"
	postgresModelUser "github.com/t67y110v/web/internal/app/model/user"
)

type PostgresStoreRepository interface {
	Create(*postgresModelUser.User) error
	FindByEmail(string) (*postgresModelUser.User, error)
	FindByID(string) (*postgresModelUser.User, error)
	GetProtocols() ([]postgresModelProtocol.Protocol, error)
}

type MongoStoreRepository interface {
	AddProduct(name, category, imgPath, description string, price, discount int) error
	GetAllProducts() ([]*mongoModel.Product, error)
	FilterByCategory(filter string) ([]*mongoModel.Product, error)
	DeleteProduct(value string) error
}
