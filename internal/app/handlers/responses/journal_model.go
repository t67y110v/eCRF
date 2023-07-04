package responses

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type GetJournal struct {
	ID            primitive.ObjectID `bson:"_id"`
	CreatedAt     time.Time          `bson:"created_at"`
	Method        string             `bson:"method"`
	ActionType    string             `bson:"action_type"`
	Initiator     string             `bson:"initiator"`
	InitiatorRole string             `bson:"initiator_role"`
	Body          interface{}        `bson:"body"`
}
