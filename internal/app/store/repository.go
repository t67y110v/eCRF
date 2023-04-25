package store

import (
	modelCenter "github.com/t67y110v/web/internal/app/model/center"
	modelOperation "github.com/t67y110v/web/internal/app/model/operation"
	modelProtocol "github.com/t67y110v/web/internal/app/model/protocol"
	modelSubject "github.com/t67y110v/web/internal/app/model/subject"
	modelUser "github.com/t67y110v/web/internal/app/model/user"
)

type PostgresStoreRepository interface {
	ProtocolStoreRepository
	UserStoreRepository
	CenterStoreRepository
}

type MongoStoreRepository interface {
	SubjectStoreRepository
	JournalRepository
}

type JournalRepository interface {
	SaveAction(modelOperation.Operation) error
	GetActions() ([]*modelOperation.Operation, error)
}

type SubjectStoreRepository interface {
	AddSubject(number, initials string, centerId, protocolId int) error
	GetSubjectsByProtocolId(protocolId int) ([]*modelSubject.Subject, error)
	GetSubjectByNumber(number string) (*modelSubject.Subject, error)
	DeleteSubject(number string) error
}

type CenterStoreRepository interface {
	GetCenterName(centerId int) (string, error)
	GetAllCenters() ([]modelCenter.Center, error)
	AddNewCenter(name string) error
	UpdateCenter(centerId int, name string) error
	DeleteCenter(centerId int) error
}

type UserStoreRepository interface {
	Create(*modelUser.User) error
	FindByEmail(string) (*modelUser.User, error)
	FindByID(string) (*modelUser.User, error)
	GetUsers() ([]modelUser.User, error)
	UpdateUser(ID, role, centerId int, email, name, paswword string) error
	DeleteUser(Id int) error
}

type ProtocolStoreRepository interface {
	GetProtocols() ([]modelProtocol.Protocol, error)
	GetProtocolById(ID int) (*modelProtocol.Protocol, error)
	UpdateProtocolById(ID, status int, name string, centerId int) error
	AddProtocol(name string, status, centerID int) error
	GetProtocolsByFilter(filter string, centerId int) ([]modelProtocol.Protocol, error)
	DeleteProtocol(Id int) error
}
