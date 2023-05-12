package nosqlstore

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (r *MongoStoreRepository) Demography(ctx context.Context, id primitive.ObjectID, sex, race int, date string) error {

	collection := r.store.client.Database("eCRF").Collection("subjects")
	filter := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"sex": sex, "race": race, "date": date, "updated_at": time.Now()}}

	res, err := collection.UpdateOne(ctx, filter, update)
	fmt.Println(res)
	if err != nil {
		return err
	}
	return nil

}
