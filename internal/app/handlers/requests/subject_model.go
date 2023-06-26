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

type GetSubjectByNumber struct {
	Number string `json:"number"`
}

type UpdateColorRequest struct {
	SubjectNumber string `json:"subject_number"`
	ProtocolID    int    `json:"protocol_id"`
	FieldName     string `json:"field_name"`
	Value         int    `json:"value"`
}

type UpdateColorWithCommentRequest struct {
	SubjectNumber string `json:"subject_number"`
	ProtocolID    int    `json:"protocol_id"`
	FieldName     string `json:"field_name"`
	Value         int    `json:"value"`
	Comment       string `json:"comment"`
	Reason        string `json:"reason"`
}

type UpdateValueWithColor struct {
	SubjectNumber string `json:"subject_number"`
	ProtocolID    int    `json:"protocol_id"`
	FieldName     string `json:"field_name"`
	Value         string `json:"value"`
	Color         int    `json:"color"`
}
