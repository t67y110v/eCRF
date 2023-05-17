package store

import (
	"context"

	modelCenter "github.com/t67y110v/web/internal/app/model/center"
	modelOperation "github.com/t67y110v/web/internal/app/model/operation"
	modelProtocol "github.com/t67y110v/web/internal/app/model/protocol"
	modelSubject "github.com/t67y110v/web/internal/app/model/subject"
	modelUser "github.com/t67y110v/web/internal/app/model/user"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type PostgresStoreRepository interface {
	ProtocolStoreRepository
	UserStoreRepository
	CenterStoreRepository
}

type MongoSubjectRepository interface {
	AddSubject(number, initials string, centerId, protocolId int) error
	GetSubjectsByProtocolId(protocolId int) ([]*modelSubject.Subject, error)
	GetSubjectByNumber(number string, protocolID int) (*modelSubject.Subject, error)
	DeleteSubject(number string) error
}

type MongoJournalRepository interface {
	SaveAction(context.Context, modelOperation.Operation) error
	SaveProtocolAction(context.Context, modelOperation.Operation) error
	GetActions() ([]*modelOperation.Operation, error)
}

type MongoScreeningRepository interface {
	InformaionConsent(ctx context.Context, id primitive.ObjectID, dateOfSign, timeOfSign string, isSigned, receivedAnInsurancePolicy, receivedAnInformaionConsent int) error
	Demography(ctx context.Context, id primitive.ObjectID, sex, race int, date string) error
	Anthropometry(ctx context.Context, id primitive.ObjectID, anthropometricDataBeenMeasured int, reasonIfNot, dateOfStartMeasured string, weightOfBody, hightOfBody, indexWeigthOfBody int) error
	InclusionCriteria(ctx context.Context, id primitive.ObjectID, presenceOfAnInformationPanel, aged18To55Years, negativeHIVTestResult, bodyMassIndex, absenceOfAcuteInfectiousDiseases, consentToUseEffectiveMethodsOfContraception, negativePregnancyTest, negativeDrugTest, negativeAlcoholTest, noHistoryOfSeverePostVaccinationReactions, indicatorsBloodTestsAtScreeningWithin, noMyocardialChanges, negativeTestResultForCOVID, noContraindicationsToVaccination int) error
	ExclusionСriteria(ctx context.Context, id primitive.ObjectID, lackOfSignedInformedConsent, steroidTherapy, therapyWithImmunosuppressiveDrugs, femaleSubjectsDuringPregnancy, strokeInLessThanOneYear, chronicSystemicInfections, aggravatedAllergicHistory, presenceOfAHistoryOfNeoplasms, historyOfSplenectomy, neutropenia, subjectsWithActiveSyphilis, anorexia, extensiveTattoos, takingNarcoticAndPsychostimulantDrugs, smokingMoretThanTenCigarettesADay, alcoholIntake, plannedHospitalization, donorBloodDonation, subjectParticipationInAnyOtherStudy, anyVaccinationInTheLastMonth, inabilityToReadInRussian, researchCenterStaff, anyOtherStateOfTheSubjectOfTheStudy int) error
	UpdateColor(ctx context.Context, id primitive.ObjectID, fieldToUpdate string, field int) error
	UpdateColorWithComment(ctx context.Context, id primitive.ObjectID, fieldToUpdate, reason, comment string, color int) error
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
