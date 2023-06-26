package nosqlstore

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type MongoColorRepository struct {
	store *Store
}

func (r *MongoColorRepository) UpdateColor(ctx context.Context, id primitive.ObjectID, fieldToUpdate string, field int) error {
	collection := r.store.client.Database("eCRF").Collection("subjects")

	filter := bson.M{"_id": id}
	update := bson.M{
		"$set": bson.M{
			fmt.Sprintf("%scolor", fieldToUpdate): field,
		},
	}

	_, err := collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return err
	}
	return nil
}

func (r *MongoColorRepository) UpdateColorWithComment(ctx context.Context, id primitive.ObjectID, fieldToUpdate, reason, comment string, color int) error {
	collection := r.store.client.Database("eCRF").Collection("subjects")
	filter := bson.M{"_id": id}
	update := bson.M{
		"$set": bson.M{
			fmt.Sprintf("%scolor", fieldToUpdate):   color,
			fmt.Sprintf("%sreason", fieldToUpdate):  reason,
			fmt.Sprintf("%scomment", fieldToUpdate): comment,
		},
	}
	_, err := collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return err
	}
	return nil

}

func (r *MongoColorRepository) UpdateFieldIntValue(ctx context.Context, id primitive.ObjectID, fieldToUpdate, fieldValue string, value, color int) error {
	collection := r.store.client.Database("eCRF").Collection("subjects")
	filter := bson.M{"_id": id}
	update := bson.M{
		"$set": bson.M{
			fmt.Sprintf("%scolor", fieldToUpdate):          color,
			fmt.Sprintf("%s%s", fieldToUpdate, fieldValue): value,
		},
	}
	_, err := collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return err
	}
	return nil
}

func (r *MongoColorRepository) UpdateFieldStringValue(ctx context.Context, id primitive.ObjectID, fieldToUpdate, fieldValue, value string, color int) error {
	collection := r.store.client.Database("eCRF").Collection("subjects")
	filter := bson.M{"_id": id}
	update := bson.M{
		"$set": bson.M{
			fmt.Sprintf("%scolor", fieldToUpdate):          color,
			fmt.Sprintf("%s%s", fieldToUpdate, fieldValue): value,
		},
	}
	_, err := collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return err
	}
	return nil
}
