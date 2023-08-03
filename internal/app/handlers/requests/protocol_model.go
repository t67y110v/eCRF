package requests

type SaveProtocol struct {
	ID       int
	Status   int
	Name     string
	CenterId int
}

type AddProtocol struct {
	Name         string `json:"name"`
	Status       int    `json:"status"`
	CenterID     int    `json:"center_id"`
	Number       int    `json:"number"`
	Organization int    `json:"organization_id"`
}

type DeleteProtocol struct {
	ID int
}

type GetProtocols struct {
	UserCentersID int
}
