package journal

import (
	"context"
	"time"

	model "github.com/t67y110v/web/internal/app/model/operation"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (o *Journal) SaveAction(ctx context.Context, method, initiator, initiatorRole, actionType string, body ...interface{}) {
	op := model.Operation{
		ID:            primitive.NewObjectID(),
		Method:        method,
		Initiator:     initiator,
		InitiatorRole: initiatorRole,
		ActionType:    actionType,
		CreatedAt:     time.Now(),
		Body:          body,
	}

	o.mgStore.Journal().SaveAction(ctx, op)
}

// func (j * Journal ) SaveProtocol ( ctx context.Context, p *model.OperationProtocol)  {
// 	op := model.Operation {

// 		Body: p,
// 	}
// 	j.mgStore.Journal().SaveAction(ctx, )
// }
