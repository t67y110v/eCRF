package nosqlstore

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (r *MongoStoreRepository) Demography(ctx context.Context, id primitive.ObjectID, sex, race int, date string) error {

	collection := r.store.client.Database("eCRF").Collection("subjects")
	filter := bson.M{"_id": id}
	update := bson.M{
		"$set": bson.M{
			"demographysubject.sex":  sex,
			"demographysubject.race": race,
			"demographysubject.date": date,
			"updated_at":             time.Now(),
		},
	}

	_, err := collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return err
	}
	return nil

}

func (r *MongoStoreRepository) InformaionConsent(
	ctx context.Context,
	id primitive.ObjectID,
	dateOfSign int,
	timeOfSign string,
	isSigned,
	receivedAnInsurancePolicy,
	receivedAnInformaionConsent bool,
) error {
	collection := r.store.client.Database("eCRF").Collection("subjects")
	filter := bson.M{"_id": id}
	update := bson.M{
		"$set": bson.M{
			"informaionconsent.signed":                         isSigned,
			"informaionconsent.date_of_sign":                   dateOfSign,
			"informaionconsent.time_of_sign":                   timeOfSign,
			"informaionconsent.received_an_insurance_policy":   receivedAnInsurancePolicy,
			"informaionconsent.received_an_informaion_consent": receivedAnInformaionConsent,
		},
	}

	_, err := collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return err
	}
	return nil
}

func (r *MongoStoreRepository) Anthropometry(
	ctx context.Context,
	id primitive.ObjectID,
	anthropometricDataBeenMeasured bool,
	reasonIfNot,
	dateOfStartMeasured string,
	weightOfBody,
	hightOfBody,
	indexWeigthOfBody int,
) error {
	collection := r.store.client.Database("eCRF").Collection("subjects")
	filter := bson.M{"_id": id}
	update := bson.M{
		"$set": bson.M{
			"anthropometry.anthropometric_data_been_measured": anthropometricDataBeenMeasured,
			"anthropometry.reason_if_not":                     reasonIfNot,
			"anthropometry.date_of_start_measured":            dateOfStartMeasured,
			"anthropometry.weight_of_body":                    weightOfBody,
			"anthropometry.hight_of_body":                     hightOfBody,
			"anthropometry.index_weight_of_body":              indexWeigthOfBody,
		},
	}

	_, err := collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return err
	}
	return nil
}
