package requests

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
