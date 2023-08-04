package nosqlstore

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type MongoOffSiteBlockRepository struct {
	store *Store
}

func (r *MongoOffSiteBlockRepository) AdverseEvents(ctx context.Context, id primitive.ObjectID, AdverseEventsRegistered int, DescriptionOfTheAdverseEvent string, DateOfStartAE string, isContinuesEnd int, DateOfEndAE string, Severity int, RecurringPhenomenon int, AssociationWithTheDrugUsed int, Foresight int, ConnectionBetweenAEAndDU int, RenewalAfterUse int, LocalReaction int, SubjectDropout int, MeasuresTaken int, MeasuresTakenOnUDCondition int, Exodus int, IsItSerious int, SeverityCriterion int, TestImpact int, DoseEffect int, ImpactOnHospitalStay int, RelationshipWithMedication int, Expectancy int) error {
	collection := r.store.client.Database("eCRF").Collection("subjects")
	fiter := bson.M{"_id": id}
	update := bson.M{
		"$set": bson.M{
			"OffSiteBlock.AdverseEvents.AdverseEventsRegisteredCondition.AdverseEventsRegistered":           AdverseEventsRegistered, //dateStart,
			"OffSiteBlock.AdverseEvents.AdverseEventsRegisteredCondition.color":                             1,
			"OffSiteBlock.AdverseEvents.DescriptionOfTheAdverseEventCondition.DescriptionOfTheAdverseEvent": DescriptionOfTheAdverseEvent,
			"OffSiteBlock.AdverseEvents.DescriptionOfTheAdverseEventCondition.color":                        1,
			"OffSiteBlock.AdverseEvents.DateOfStartAECondition.DateOfStartAE":                               DateOfStartAE,
			"OffSiteBlock.AdverseEvents.DateOfStartAECondition.color":                                       1,
			"OffSiteBlock.AdverseEvents.DateOfEndAECondition.DateOfEndAE":                                   DateOfEndAE,
			"OffSiteBlock.AdverseEvents.DateOfEndAECondition.color":                                         1,
			"OffSiteBlock.AdverseEvents.DateOfEndAECondition.iscontinius":                                   isContinuesEnd, //continues
			"OffSiteBlock.AdverseEvents.SeverityCondition.Severity":                                         Severity,
			"OffSiteBlock.AdverseEvents.SeverityCondition.color":                                            1,
			"OffSiteBlock.AdverseEvents.RecurringPhenomenonCondition.RecurringPhenomenon":                   RecurringPhenomenon,
			"OffSiteBlock.AdverseEvents.RecurringPhenomenonCondition.color":                                 1,
			"OffSiteBlock.AdverseEvents.AssociationWithTheDrugUsedCondition.AssociationWithTheDrugUsed":     AssociationWithTheDrugUsed,
			"OffSiteBlock.AdverseEvents.AssociationWithTheDrugUsedCondition.color":                          1,
			"OffSiteBlock.AdverseEvents.ForesightCondition.Foresight":                                       Foresight,
			"OffSiteBlock.AdverseEvents.ForesightCondition.color":                                           1,
			"OffSiteBlock.AdverseEvents.ConnectionBetweenAEAndDUCondition.ConnectionBetweenAEAndDU":         ConnectionBetweenAEAndDU,
			"OffSiteBlock.AdverseEvents.ConnectionBetweenAEAndDUCondition.color":                            1,
			"OffSiteBlock.AdverseEvents.RenewalAfterUseCondition.RenewalAfterUse":                           RenewalAfterUse,
			"OffSiteBlock.AdverseEvents.RenewalAfterUseCondition.color":                                     1,
			"OffSiteBlock.AdverseEvents.LocalReactionCondition.LocalReaction":                               LocalReaction,
			"OffSiteBlock.AdverseEvents.LocalReactionCondition.color":                                       1,
			"OffSiteBlock.AdverseEvents.SubjectDropoutCondition.SubjectDropout":                             SubjectDropout,
			"OffSiteBlock.AdverseEvents.SubjectDropoutCondition.color":                                      1,
			"OffSiteBlock.AdverseEvents.MeasuresTakenCondition.MeasuresTaken":                               MeasuresTaken,
			"OffSiteBlock.AdverseEvents.MeasuresTakenCondition.color":                                       1,
			"OffSiteBlock.AdverseEvents.MeasuresTakenOnUDCondition.MeasuresTakenOnUD":                       MeasuresTakenOnUDCondition,
			"OffSiteBlock.AdverseEvents.MeasuresTakenOnUDCondition.color":                                   1,
			"OffSiteBlock.AdverseEvents.ExodusCondition.Exodus":                                             Exodus,
			"OffSiteBlock.AdverseEvents.ExodusCondition.color":                                              1,
			"OffSiteBlock.AdverseEvents.IsItSeriousCondition.IsItSerious":                                   IsItSerious,
			"OffSiteBlock.AdverseEvents.IsItSeriousCondition.color":                                         1,
			"OffSiteBlock.AdverseEvents.SeverityCriterionCondition.SeverityCriterion":                       SeverityCriterion,
			"OffSiteBlock.AdverseEvents.SeverityCriterionCondition.color":                                   1,
			"OffSiteBlock.AdverseEvents.TestImpactCondition.TestImpact":                                     TestImpact,
			"OffSiteBlock.AdverseEvents.TestImpactCondition.color":                                          1,
			"OffSiteBlock.AdverseEvents.DoseEffectCondition.DoseEffect":                                     DoseEffect,
			"OffSiteBlock.AdverseEvents.DoseEffectCondition.color":                                          1,
			"OffSiteBlock.AdverseEvents.ImpactOnHospitalStayCondition.ImpactOnHospitalStay":                 ImpactOnHospitalStay,
			"OffSiteBlock.AdverseEvents.ImpactOnHospitalStayCondition.color":                                1,
			"OffSiteBlock.AdverseEvents.RelationshipWithMedicationCondition.RelationshipWithMedication":     RelationshipWithMedication,
			"OffSiteBlock.AdverseEvents.RelationshipWithMedicationCondition.color":                          1,
			"OffSiteBlock.AdverseEvents.ExpectancyCondition.Expectancy":                                     Expectancy,
			"OffSiteBlock.AdverseEvents.ExpectancyCondition.color":                                          1,
		},
	}

	_, err := collection.UpdateOne(ctx, fiter, update)
	if err != nil {
		return err
	}
	return nil

}
