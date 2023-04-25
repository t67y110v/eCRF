package nosqlstore

import (
	"context"
	"errors"
	"time"

	model "github.com/t67y110v/web/internal/app/model/subject"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type MongoStoreRepository struct {
	store *Store
}

/*
Number            string             `bson:"number"`
CenterID          int                `bson:"center_id" `
ProtocolId        int                `bson:"protocol_id"`
Initials          string             `bson:"initials"`
*/
func (r *MongoStoreRepository) AddSubject(number, initials string, centerId, protocolId int) error {
	if number == "" || initials == "" {
		return errors.New("empty fields")
	}
	ctx := context.TODO()

	collection := r.store.client.Database("eCRF").Collection("subjects")

	subject := &model.Subject{
		ID:         primitive.NewObjectID(),
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
		Number:     number,
		CenterID:   centerId,
		ProtocolId: protocolId,
		Initials:   initials,
	}
	_, err := collection.InsertOne(ctx, subject)
	if err != nil {
		return err
	}
	return nil

}

func (r *MongoStoreRepository) GetSubjectsByProtocolId(protocolId int) ([]*model.Subject, error) {
	filter := bson.D{
		primitive.E{
			Key:   "protocol_id",
			Value: protocolId,
		},
	}
	ctx := context.TODO()
	collection := r.store.client.Database("eCRF").Collection("subjects")
	cur, err := collection.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	var subjects []*model.Subject
	for cur.Next(ctx) {
		var s model.Subject
		err := cur.Decode(&s)
		if err != nil {
			return nil, err
		}
		subjects = append(subjects, &s)
	}
	if err := cur.Err(); err != nil {
		return subjects, err
	}
	cur.Close(ctx)
	if len(subjects) == 0 {
		return subjects, nil
	}
	return subjects, nil
}

func (r *MongoStoreRepository) GetSubjectByNumber(number string) (*model.Subject, error) {
	filter := bson.D{
		primitive.E{
			Key:   "number",
			Value: number,
		},
	}
	ctx := context.TODO()
	collection := r.store.client.Database("eCRF").Collection("subjects")

	cur := collection.FindOne(ctx, filter)

	var p model.Subject

	err := cur.Decode(&p)
	if err := cur.Err(); err != nil {
		return &p, err
	}
	if err != nil {
		return nil, err
	}

	return &p, nil

}

func (r *MongoStoreRepository) DeleteSubject(number string) error {
	ctx := context.TODO()
	filter := bson.D{primitive.E{
		Key:   "number",
		Value: number,
	}}

	collection := r.store.client.Database("eCRF").Collection("subjects")

	res, err := collection.DeleteOne(ctx, filter)

	if err != nil {
		return err
	}
	if res.DeletedCount == 0 {
		return errors.New("no products were deleted")
	}
	return nil

}
