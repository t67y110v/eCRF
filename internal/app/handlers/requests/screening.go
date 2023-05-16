package requests

type Request struct {
	SNum  string `json:"subject_number"`
	PID   string `json:"protocol_id"`
	FName string `json:"field_name"`
	V     int    `json:"value"`
}
