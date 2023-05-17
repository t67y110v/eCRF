package requests

type UpdateColorRequest struct {
	SubjectNumber string `json:"subject_number"`
	ProtocolID    string `json:"protocol_id"`
	FieldName     string `json:"field_name"`
	Value         int    `json:"value"`
}

type UpdateColorWithCommentRequest struct {
	SubjectNumber string `json:"subject_number"`
	ProtocolID    string `json:"protocol_id"`
	FieldName     string `json:"field_name"`
	Value         int    `json:"value"`
	Comment       string `json:"comment"`
	Reason        string `json:"reason"`
}
