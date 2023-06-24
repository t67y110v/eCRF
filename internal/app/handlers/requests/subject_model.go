package requests

type GetSubject struct {
	ProtocolID int `json:"protocol_id"`
}

type AddSubject struct {
	Number     string `json:"number"`
	Initials   string `json:"initials"`
	CenterId   int    `json:"center_id"`
	ProtocolId int    `json:"protocol_id"`
}
