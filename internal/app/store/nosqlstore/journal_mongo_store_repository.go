package nosqlstore

import (
	"context"

	model "github.com/t67y110v/web/internal/app/model/operation"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (r *MongoStoreRepository) SaveProtocolAction(ctx context.Context, operation model.Operation) error {
	return nil
}

func (r *MongoStoreRepository) SaveAction(ctx context.Context, operation model.Operation) error {

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
	opts := options.Find()
	opts.SetSort(bson.D{{Key: "updated_at", Value: -1}})
	cur, err := collection.Find(ctx, filter, opts)

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
