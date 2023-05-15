package nosqlstore

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type MongoScreeningRepository struct {
	store *Store
}

func (r *MongoScreeningRepository) InformaionConsent(
	ctx context.Context,
	id primitive.ObjectID,
	dateOfSign,
	timeOfSign string,
	isSigned,
	receivedAnInsurancePolicy,
	receivedAnInformaionConsent int,
) error {
	collection := r.store.client.Database("eCRF").Collection("subjects")
	filter := bson.M{"_id": id}
	update := bson.M{
		"$set": bson.M{
			"screening.informaionconsent.signed":                         isSigned,
			"screening.informaionconsent.date_of_sign":                   dateOfSign,
			"screening.informaionconsent.time_of_sign":                   timeOfSign,
			"screening.informaionconsent.received_an_insurance_policy":   receivedAnInsurancePolicy,
			"screening.informaionconsent.received_an_informaion_consent": receivedAnInformaionConsent,
		},
	}

	_, err := collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return err
	}
	return nil
}
func (r *MongoScreeningRepository) Demography(ctx context.Context, id primitive.ObjectID, sex, race int, date string) error {

	collection := r.store.client.Database("eCRF").Collection("subjects")
	filter := bson.M{"_id": id}
	update := bson.M{
		"$set": bson.M{
			"screening.demography.sex":  sex,
			"screening.demography.race": race,
			"screening.demography.date": date,
			"updated_at":                time.Now(),
		},
	}

	_, err := collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return err
	}
	return nil

}

func (r *MongoScreeningRepository) Anthropometry(
	ctx context.Context,
	id primitive.ObjectID,
	anthropometricDataBeenMeasured int,
	reasonIfNot,
	dateOfStartMeasured string,
	weightOfBody,
	heightOfBody,
	indexWeigthOfBody int,
) error {
	collection := r.store.client.Database("eCRF").Collection("subjects")
	filter := bson.M{"_id": id}
	update := bson.M{
		"$set": bson.M{
			"screening.anthropometry.anthropometric_data_been_measured": anthropometricDataBeenMeasured,
			"screening.anthropometry.reason_if_not":                     reasonIfNot,
			"screening.anthropometry.date_of_start_measured":            dateOfStartMeasured,
			"screening.anthropometry.weight_of_body":                    weightOfBody,
			"screening.anthropometry.height_of_body":                    heightOfBody,
			"screening.anthropometry.index_weight_of_body":              indexWeigthOfBody,
		},
	}

	_, err := collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return err
	}
	return nil
}

func (r *MongoScreeningRepository) InclusionCriteria(
	ctx context.Context,
	id primitive.ObjectID,
	presenceOfAnInformationPanel,
	aged18To55Years,
	negativeHIVTestResult,
	bodyMassIndex,
	absenceOfAcuteInfectiousDiseases,
	consentToUseEffectiveMethodsOfContraception,
	negativePregnancyTest,
	negativeDrugTest,
	negativeAlcoholTest,
	noHistoryOfSeverePostVaccinationReactions,
	indicatorsBloodTestsAtScreeningWithin,
	noMyocardialChanges,
	negativeTestResultForCOVID,
	noContraindicationsToVaccination int,

) error {

	collection := r.store.client.Database("eCRF").Collection("subjects")
	filter := bson.M{"_id": id}
	update := bson.M{
		"$set": bson.M{
			"screening.inclusioncriteria.presence_of_an_information_panel":                  presenceOfAnInformationPanel,
			"screening.inclusioncriteria.aged_18_to_55_years":                               aged18To55Years,
			"screening.inclusioncriteria.negative_hiv_test_result":                          negativeHIVTestResult,
			"screening.inclusioncriteria.body_mass_index":                                   bodyMassIndex,
			"screening.inclusioncriteria.absence_of_acute_infectious_diseases":              absenceOfAcuteInfectiousDiseases,
			"screening.inclusioncriteria.consent_to_use_effective_methods_of_contraception": consentToUseEffectiveMethodsOfContraception,
			"screening.inclusioncriteria.negative_pregnancy_test":                           negativePregnancyTest,
			"screening.inclusioncriteria.negative_drug_test":                                negativeDrugTest,
			"screening.inclusioncriteria.negative_alcohol_test":                             negativeAlcoholTest,
			"screening.inclusioncriteria.no_history_of_severe_post_vaccination_reactions":   noHistoryOfSeverePostVaccinationReactions,
			"screening.inclusioncriteria.indicators_blood_tests_at_screening_within":        indicatorsBloodTestsAtScreeningWithin,
			"screening.inclusioncriteria.no_myocardial_changes":                             noMyocardialChanges,
			"screening.inclusioncriteria.negative_test_result_for_COVID":                    negativeTestResultForCOVID,
			"screening.inclusioncriteria.no_contraindications_to_vaccination":               noContraindicationsToVaccination,
		},
	}

	_, err := collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return err
	}
	return nil
}

func (r *MongoScreeningRepository) ExclusionСriteria(
	ctx context.Context,
	id primitive.ObjectID,
	lackOfSignedInformedConsent,
	steroidTherapy,
	therapyWithImmunosuppressiveDrugs,
	femaleSubjectsDuringPregnancy,
	strokeInLessThanOneYear,
	chronicSystemicInfections,
	aggravatedAllergicHistory,
	presenceOfAHistoryOfNeoplasms,
	historyOfSplenectomy,
	neutropenia,
	subjectsWithActiveSyphilis,
	anorexia,
	extensiveTattoos,
	takingNarcoticAndPsychostimulantDrugs,
	smokingMoretThanTenCigarettesADay,
	alcoholIntake,
	plannedHospitalization,
	donorBloodDonation,
	subjectParticipationInAnyOtherStudy,
	anyVaccinationInTheLastMonth,
	inabilityToReadInRussian,
	researchCenterStaff,
	anyOtherStateOfTheSubjectOfTheStudy bool,

) error {
	collection := r.store.client.Database("eCRF").Collection("subjects")
	filter := bson.M{"_id": id}
	update := bson.M{
		"$set": bson.M{
			"screening.exclusionсriteria.lack_of_signed_informed_consent":             lackOfSignedInformedConsent,
			"screening.exclusionсriteria.steroid_therapy":                             steroidTherapy,
			"screening.exclusionсriteria.therapy_with_immunosuppressive_drugs":        therapyWithImmunosuppressiveDrugs,
			"screening.exclusionсriteria.female_subjects_during_pregnancy":            femaleSubjectsDuringPregnancy,
			"screening.exclusionсriteria.stroke_in_less_than_one_year":                strokeInLessThanOneYear,
			"screening.exclusionсriteria.chronic_systemic_infections":                 chronicSystemicInfections,
			"screening.exclusionсriteria.aggravated_allergic_history":                 aggravatedAllergicHistory,
			"screening.exclusionсriteria.presence_of_a_history_of_neoplasms":          presenceOfAHistoryOfNeoplasms,
			"screening.exclusionсriteria.history_of_splenectomy":                      historyOfSplenectomy,
			"screening.exclusionсriteria.neutropenia":                                 neutropenia,
			"screening.exclusionсriteria.subjects_with_active_syphilis":               subjectsWithActiveSyphilis,
			"screening.exclusionсriteria.anorexia":                                    anorexia,
			"screening.exclusionсriteria.extensive_tattoos":                           extensiveTattoos,
			"screening.exclusionсriteria.taking_narcotic_and_psychostimulant_drugs":   takingNarcoticAndPsychostimulantDrugs,
			"screening.exclusionсriteria.smoking_more_than_ten_cigarettes_a_day":      smokingMoretThanTenCigarettesADay,
			"screening.exclusionсriteria.alcohol_intake":                              alcoholIntake,
			"screening.exclusionсriteria.planned_hospitalization":                     plannedHospitalization,
			"screening.exclusionсriteria.donor_blood_donation":                        donorBloodDonation,
			"screening.exclusionсriteria.subject_participation_in_any_other_study":    subjectParticipationInAnyOtherStudy,
			"screening.exclusionсriteria.any_vaccination_in_the_last_month":           anyVaccinationInTheLastMonth,
			"screening.exclusionсriteria.inability_to_read_in_russian":                inabilityToReadInRussian,
			"screening.exclusionсriteria.research_center_staff":                       researchCenterStaff,
			"screening.exclusionсriteria.any_other_state_of_the_subject_of_the_study": anyOtherStateOfTheSubjectOfTheStudy,
		},
	}

	_, err := collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return err
	}
	return nil
}
