package responses

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type GetSubject struct {
	ID         primitive.ObjectID `json:"_id"`
	CreatedAt  time.Time          `json:"created_at"`
	UpdatedAt  time.Time          `json:"updated_at"`
	Number     string             `json:"number"`
	CenterID   int                `json:"center_id" `
	ProtocolId int                `json:"protocol_id"`
	Initials   string             `json:"initials"`
}
