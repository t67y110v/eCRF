package operations

import (
	"time"

	model "github.com/t67y110v/web/internal/app/model/operation"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (o *Operations) SaveAction(method, code, initiator, actionType string) {
	op := model.Operation{
		Method:     method,
		Code:       code,
		Initiator:  initiator,
		ActionType: actionType,
		ID:         primitive.NewObjectID(),
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
	}

	o.mgStore.Repository().SaveAction(op)
}