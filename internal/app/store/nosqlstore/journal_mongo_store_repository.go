package nosqlstore

import (
	"context"

	model "github.com/t67y110v/web/internal/app/model/operation"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func (r *MongoStoreRepository) SaveAction(operation model.Operation) error {

	ctx := context.TODO()

	collection := r.store.client.Database("eCRF").Collection("operations")

	_, err := collection.InsertOne(ctx, operation)
	if err != nil {
		return err
	}
	return nil

}

func (r *MongoStoreRepository) GetActions() ([]*model.Operation, error) {

	filter := bson.D{{}}
	ctx := context.TODO()
	collection := r.store.client.Database("eCRF").Collection("operations")

	cur, err := collection.Find(ctx, filter)

	var operations []*model.Operation

	if err != nil {
		return nil, err
	}
	for cur.Next(ctx) {
		var o model.Operation
		err := cur.Decode(&o)
		if err != nil {
			return nil, err
		}
		operations = append(operations, &o)
	}
	if err := cur.Err(); err != nil {
		return operations, err
	}
	cur.Close(ctx)
	if len(operations) == 0 {
		return operations, mongo.ErrNoDocuments

	}
	return operations, nil

}
