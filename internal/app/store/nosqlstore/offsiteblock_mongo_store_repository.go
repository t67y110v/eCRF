package nosqlstore

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type MongoOffSiteBlockRepository struct {
	store *Store
}

func (r *MongoOffSiteBlockRepository) AdverseEvents(ctx context.Context, id primitive.ObjectID, AdverseEventsRegistered int, DescriptionOfTheAdverseEvent string, DateOfStartAE string, isContinuesEnd int, DateOfEndAE string, Severity int, RecurringPhenomenon int, AssociationWithTheDrugUsed int, Foresight int, ConnectionBetweenAEAndDU int, RenewalAfterUse int, LocalReaction int, SubjectDropout int, MeasuresTaken int, MeasuresTakenOnUDCondition int, Exodus int, IsItSerious int, SeverityCriterion int, TestImpact int, DoseEffect int, ImpactOnHospitalStay int, RelationshipWithMedication int, Expectancy int, count int) error {
	collection := r.store.client.Database("eCRF").Collection("subjects")

	/*
		update := bson.M{
			"$push": bson.M{
				"offsiteblock.events_slice": bson.M{
					"$each": []interface{}{bson.M{
						"AdverseEventsRegisteredCondition.AdverseEventsRegistered": AdverseEventsRegistered,
					}},
				},
			},

	*/
	fiter := bson.M{"_id": id}
	update := bson.M{
		"$set": bson.M{
			fmt.Sprintf("offsiteblock.adverseevents.%d.adverseeventsregisteredcondition.adverseeventsregistered", count):           AdverseEventsRegistered, //dateStart,
			fmt.Sprintf("offsiteblock.adverseevents.%d.adverseeventsregisteredcondition.color", count):                             1,
			fmt.Sprintf("offsiteblock.adverseevents.%d.descriptionoftheadverseeventcondition.descriptionoftheadverseevent", count): DescriptionOfTheAdverseEvent,
			fmt.Sprintf("offsiteblock.adverseevents.%d.descriptionoftheadverseeventcondition.color", count):                        1,
			fmt.Sprintf("offsiteblock.adverseevents.%d.dateofstartaecondition.dateofstartae", count):                               DateOfStartAE,
			fmt.Sprintf("offsiteblock.adverseevents.%d.dateofstartaecondition.color", count):                                       1,
			fmt.Sprintf("offsiteblock.adverseevents.%d.dateofendaecondition.dateofendae", count):                                   DateOfEndAE,
			fmt.Sprintf("offsiteblock.adverseevents.%d.dateofendaecondition.color", count):                                         1,
			fmt.Sprintf("offsiteblock.adverseevents.%d.dateofendaecondition.iscontinius", count):                                   isContinuesEnd, //continues
			fmt.Sprintf("offsiteblock.adverseevents.%d.severitycondition.severity", count):                                         Severity,
			fmt.Sprintf("offsiteblock.adverseevents.%d.severitycondition.color", count):                                            1,
			fmt.Sprintf("offsiteblock.adverseevents.%d.recurringphenomenoncondition.recurringphenomenon", count):                   RecurringPhenomenon,
			fmt.Sprintf("offsiteblock.adverseevents.%d.recurringphenomenoncondition.color", count):                                 1,
			fmt.Sprintf("offsiteblock.adverseevents.%d.associationwiththedrugusedcondition.associationwiththedrugused", count):     AssociationWithTheDrugUsed,
			fmt.Sprintf("offsiteblock.adverseevents.%d.associationwiththedrugusedcondition.color", count):                          1,
			fmt.Sprintf("offsiteblock.adverseevents.%d.foresightcondition.foresight", count):                                       Foresight,
			fmt.Sprintf("offsiteblock.adverseevents.%d.foresightcondition.color", count):                                           1,
			fmt.Sprintf("offsiteblock.adverseevents.%d.connectionbetweenaeandducondition.connectionbetweenaeanddu", count):         ConnectionBetweenAEAndDU,
			fmt.Sprintf("offsiteblock.adverseevents.%d.connectionbetweenaeandducondition.color", count):                            1,
			fmt.Sprintf("offsiteblock.adverseevents.%d.renewalafterusecondition.renewalafteruse", count):                           RenewalAfterUse,
			fmt.Sprintf("offsiteblock.adverseevents.%d.renewalafterusecondition.color", count):                                     1,
			fmt.Sprintf("offsiteblock.adverseevents.%d.localreactioncondition.localreaction", count):                               LocalReaction,
			fmt.Sprintf("offsiteblock.adverseevents.%d.localreactioncondition.color", count):                                       1,
			fmt.Sprintf("offsiteblock.adverseevents.%d.subjectdropoutcondition.subjectdropout", count):                             SubjectDropout,
			fmt.Sprintf("offsiteblock.adverseevents.%d.subjectdropoutcondition.color", count):                                      1,
			fmt.Sprintf("offsiteblock.adverseevents.%d.measurestakencondition.measurestaken", count):                               MeasuresTaken,
			fmt.Sprintf("offsiteblock.adverseevents.%d.measurestakencondition.color", count):                                       1,
			fmt.Sprintf("offsiteblock.adverseevents.%d.measurestakenonudcondition.measurestakenonud", count):                       MeasuresTakenOnUDCondition,
			fmt.Sprintf("offsiteblock.adverseevents.%d.measurestakenonudcondition.color", count):                                   1,
			fmt.Sprintf("offsiteblock.adverseevents.%d.exoduscondition.exodus", count):                                             Exodus,
			fmt.Sprintf("offsiteblock.adverseevents.%d.exoduscondition.color", count):                                              1,
			fmt.Sprintf("offsiteblock.adverseevents.%d.isitseriouscondition.isitserious", count):                                   IsItSerious,
			fmt.Sprintf("offsiteblock.adverseevents.%d.isitseriouscondition.color", count):                                         1,
			fmt.Sprintf("offsiteblock.adverseevents.%d.severitycriterioncondition.severitycriterion", count):                       SeverityCriterion,
			fmt.Sprintf("offsiteblock.adverseevents.%d.severitycriterioncondition.color", count):                                   1,
			fmt.Sprintf("offsiteblock.adverseevents.%d.testimpactcondition.testimpact", count):                                     TestImpact,
			fmt.Sprintf("offsiteblock.adverseevents.%d.testimpactcondition.color", count):                                          1,
			fmt.Sprintf("offsiteblock.adverseevents.%d.doseeffectcondition.doseeffect", count):                                     DoseEffect,
			fmt.Sprintf("offsiteblock.adverseevents.%d.doseeffectcondition.color", count):                                          1,
			fmt.Sprintf("offsiteblock.adverseevents.%d.impactonhospitalstaycondition.impactonhospitalstay", count):                 ImpactOnHospitalStay,
			fmt.Sprintf("offsiteblock.adverseevents.%d.impactonhospitalstaycondition.color", count):                                1,
			fmt.Sprintf("offsiteblock.adverseevents.%d.relationshipwithmedicationcondition.relationshipwithmedication", count):     RelationshipWithMedication,
			fmt.Sprintf("offsiteblock.adverseevents.%d.relationshipwithmedicationcondition.color", count):                          1,
			fmt.Sprintf("offsiteblock.adverseevents.%d.expectancycondition.expectancy", count):                                     Expectancy,
			fmt.Sprintf("offsiteblock.adverseevents.%d.expectancycondition.color", count):                                          1,
		},
	}

	_, err := collection.UpdateOne(ctx, fiter, update)
	if err != nil {
		return err
	}
	return nil

}
