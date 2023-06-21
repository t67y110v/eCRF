package responses

type SaveProtocol struct {
	Message string `json:"message"`
}

type AddProtocol struct {
	Message string `json:"message"`
}

type DeleteProtocol struct {
	Message string `json:"message"`
}

type GetProtocols struct {
	Id       int    `json:"protocol_id"  gorm:"primaryKey"`
	Name     string `json:"name"`
	Status   int    `json:"status"`
	CenterId int    `json:"center_id"  gorm:"primaryKey"`
}
