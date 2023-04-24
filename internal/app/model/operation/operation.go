package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Operation struct {
	ID         primitive.ObjectID `bson:"_id"`
	CreatedAt  time.Time          `bson:"created_at"`
	UpdatedAt  time.Time          `bson:"updated_at"`
	Method     string             `bson:"method"`
	Code       string             `bson:"code"`
	Initiator  string             `bson:"initiator"`
	ActionType string             `bson:"action_type"`
}
