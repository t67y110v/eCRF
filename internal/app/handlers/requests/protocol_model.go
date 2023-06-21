package requests

type SaveProtocol struct {
	ID       int
	Status   int
	Name     string
	CenterId int
}

type AddProtocol struct {
	Name     string
	Status   int
	CenterID int
}

type DeleteProtocol struct {
	ID int
}

type GetProtocols struct {
	UserCentersID int
}
