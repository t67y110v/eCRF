package nosqlstore

import (
	"context"
	"fmt"
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
			"screening.informaionconsent.signedcondition.signed":                                              isSigned,
			"screening.informaionconsent.signedcondition.color":                                               1,
			"screening.informaionconsent.dateofsigncondition.date_of_sign":                                    dateOfSign,
			"screening.informaionconsent.dateofsigncondition.color":                                           1,
			"screening.informaionconsent.timeofsigncondition.time_of_sign":                                    timeOfSign,
			"screening.informaionconsent.timeofsigncondition.color":                                           1,
			"screening.informaionconsent.receivedaninsurancepolicycondition.received_an_insurance_policy":     receivedAnInsurancePolicy,
			"screening.informaionconsent.receivedaninsurancepolicycondition.color":                            1,
			"screening.informaionconsent.receivedaninformaionconsentcondition.received_an_informaion_consent": receivedAnInformaionConsent,
			"screening.informaionconsent.receivedaninformaionconsentcondition.color":                          1,
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
			"screening.demography.sexcondition.sex":         sex,
			"screening.demography.sexcondition.color":       1,
			"screening.demography.racecondition.race":       race,
			"screening.demography.racecondition.color":      1,
			"screening.demography.birthdatecondition.date":  date,
			"screening.demography.birthdatecondition.color": 1,
			"updated_at": time.Now(),
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
			"screening.anthropometry.anthropometricdatabeenmeasuredcondition.anthropometric_data_been_measured": anthropometricDataBeenMeasured,
			"screening.anthropometry.anthropometricdatabeenmeasuredcondition.color":                             1,
			"screening.anthropometry.reasonifnotcondition.reason_if_not":                                        reasonIfNot,
			"screening.anthropometry.reasonifnotcondition.color":                                                1,
			"screening.anthropometry.dateofstartmeasuredcondition.date_of_start_measured":                       dateOfStartMeasured,
			"screening.anthropometry.dateofstartmeasuredcondition.color":                                        1,
			"screening.anthropometry.weightofbodycondition.weight_of_body":                                      weightOfBody,
			"screening.anthropometry.weightofbodycondition.color":                                               1,
			"screening.anthropometry.heightofbodycondition.height_of_body":                                      heightOfBody,
			"screening.anthropometry.heightofbodycondition.color":                                               1,
			"screening.anthropometry.indexweigthofbodycondition.index_weight_of_body":                           indexWeigthOfBody,
			"screening.anthropometry.indexweigthofbodycondition.color":                                          1,
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
			"screening.inclusioncriteria.presenceofaninformationpanelcondition.presence_of_an_information_panel":                                 presenceOfAnInformationPanel,
			"screening.inclusioncriteria.presenceofaninformationpanelcondition.color":                                                            1,
			"screening.inclusioncriteria.aged18to55yearscondition.aged_18_to_55_years":                                                           aged18To55Years,
			"screening.inclusioncriteria.aged18to55yearscondition.color":                                                                         1,
			"screening.inclusioncriteria.negativehivtestresultcondition.negative_hiv_test_result":                                                negativeHIVTestResult,
			"screening.inclusioncriteria.negativehivtestresultcondition.color":                                                                   1,
			"screening.inclusioncriteria.bodymassindexcondition.body_mass_index":                                                                 bodyMassIndex,
			"screening.inclusioncriteria.bodymassindexcondition.color":                                                                           1,
			"screening.inclusioncriteria.absenceofacuteinfectiousdiseasescondition.absence_of_acute_infectious_diseases":                         absenceOfAcuteInfectiousDiseases,
			"screening.inclusioncriteria.absenceofacuteinfectiousdiseasescondition.color":                                                        1,
			"screening.inclusioncriteria.consenttouseeffectivemethodsofcontraceptioncondition.consent_to_use_effective_methods_of_contraception": consentToUseEffectiveMethodsOfContraception,
			"screening.inclusioncriteria.consenttouseeffectivemethodsofcontraceptioncondition.color":                                             1,
			"screening.inclusioncriteria.negativepregnancytestcondition.negative_pregnancy_test":                                                 negativePregnancyTest,
			"screening.inclusioncriteria.negativepregnancytestcondition.color":                                                                   1,
			"screening.inclusioncriteria.negativedrugtestcondition.negative_drug_test":                                                           negativeDrugTest,
			"screening.inclusioncriteria.negativedrugtestcondition.color":                                                                        1,
			"screening.inclusioncriteria.negativealcoholtestcondition.negative_alcohol_test":                                                     negativeAlcoholTest,
			"screening.inclusioncriteria.negativealcoholtestcondition.color":                                                                     1,
			"screening.inclusioncriteria.nohistoryofseverepostvaccinationreactionscondition.no_history_of_severe_post_vaccination_reactions":     noHistoryOfSeverePostVaccinationReactions,
			"screening.inclusioncriteria.nohistoryofseverepostvaccinationreactionscondition.color":                                               1,
			"screening.inclusioncriteria.indicatorsbloodtestsatscreeningwithincondition.indicators_blood_tests_at_screening_within":              indicatorsBloodTestsAtScreeningWithin,
			"screening.inclusioncriteria.indicatorsbloodtestsatscreeningwithincondition.color":                                                   1,
			"screening.inclusioncriteria.nomyocardialchangescondition.no_myocardial_changes":                                                     noMyocardialChanges,
			"screening.inclusioncriteria.nomyocardialchangescondition.color":                                                                     1,
			"screening.inclusioncriteria.negativetestresultforcovidcondition.negative_test_result_for_COVID":                                     negativeTestResultForCOVID,
			"screening.inclusioncriteria.negativetestresultforcovidcondition.color":                                                              1,
			"screening.inclusioncriteria.nocontraindicationstovaccinationcondition.no_contraindications_to_vaccination":                          noContraindicationsToVaccination,
			"screening.inclusioncriteria.nocontraindicationstovaccinationcondition.color":                                                        1,
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
	anyOtherStateOfTheSubjectOfTheStudy int,

) error {
	collection := r.store.client.Database("eCRF").Collection("subjects")
	filter := bson.M{"_id": id}
	update := bson.M{
		"$set": bson.M{
			"screening.exclusionсriteria.lackofsignedinformedconsentcondition.lack_of_signed_informed_consent":                     lackOfSignedInformedConsent,
			"screening.exclusionсriteria.lackofsignedinformedconsentcondition.color":                                               1,
			"screening.exclusionсriteria.steroidtherapycondition.steroid_therapy":                                                  steroidTherapy,
			"screening.exclusionсriteria.steroidtherapycondition.color":                                                            1,
			"screening.exclusionсriteria.therapywithimmunosuppressivedrugscondition.therapy_with_immunosuppressive_drugs":          therapyWithImmunosuppressiveDrugs,
			"screening.exclusionсriteria.therapywithimmunosuppressivedrugscondition.color":                                         1,
			"screening.exclusionсriteria.femalesubjectsduringpregnancycondition.female_subjects_during_pregnancy":                  femaleSubjectsDuringPregnancy,
			"screening.exclusionсriteria.femalesubjectsduringpregnancycondition.color":                                             1,
			"screening.exclusionсriteria.strokeinlessthanoneyearcondition.stroke_in_less_than_one_year":                            strokeInLessThanOneYear,
			"screening.exclusionсriteria.strokeinlessthanoneyearcondition.color":                                                   1,
			"screening.exclusionсriteria.chronicsystemicinfectionscondition.chronic_systemic_infections":                           chronicSystemicInfections,
			"screening.exclusionсriteria.chronicsystemicinfectionscondition.color":                                                 1,
			"screening.exclusionсriteria.aggravatedallergichistorycondition.aggravated_allergic_history":                           aggravatedAllergicHistory,
			"screening.exclusionсriteria.aggravatedallergichistorycondition.color":                                                 1,
			"screening.exclusionсriteria.presenceofahistoryofneoplasmscondition.presence_of_a_history_of_neoplasms":                presenceOfAHistoryOfNeoplasms,
			"screening.exclusionсriteria.presenceofahistoryofneoplasmscondition.color":                                             1,
			"screening.exclusionсriteria.historyofsplenectomycondition.history_of_splenectomy":                                     historyOfSplenectomy,
			"screening.exclusionсriteria.historyofsplenectomycondition.color":                                                      1,
			"screening.exclusionсriteria.neutropeniacondition.neutropenia":                                                         neutropenia,
			"screening.exclusionсriteria.neutropeniacondition.color":                                                               1,
			"screening.exclusionсriteria.subjectswithactivesyphiliscondition.subjects_with_active_syphilis":                        subjectsWithActiveSyphilis,
			"screening.exclusionсriteria.subjectswithactivesyphiliscondition.color":                                                1,
			"screening.exclusionсriteria.anorexiacondition.anorexia":                                                               anorexia,
			"screening.exclusionсriteria.anorexiacondition.color":                                                                  1,
			"screening.exclusionсriteria.extensivetattooscondition.extensive_tattoos":                                              extensiveTattoos,
			"screening.exclusionсriteria.extensivetattooscondition.color":                                                          1,
			"screening.exclusionсriteria.takingnarcoticandpsychostimulantdrugscondition.taking_narcotic_and_psychostimulant_drugs": takingNarcoticAndPsychostimulantDrugs,
			"screening.exclusionсriteria.takingnarcoticandpsychostimulantdrugscondition.color":                                     1,
			"screening.exclusionсriteria.smokingmoretthantencigarettesadaycondition.smoking_more_than_ten_cigarettes_a_day":        smokingMoretThanTenCigarettesADay,
			"screening.exclusionсriteria.smokingmoretthantencigarettesadaycondition.color":                                         1,
			"screening.exclusionсriteria.alcoholintakecondition.alcohol_intake":                                                    alcoholIntake,
			"screening.exclusionсriteria.alcoholintakecondition.color":                                                             1,
			"screening.exclusionсriteria.plannedhospitalizationconditiont.planned_hospitalization":                                 plannedHospitalization,
			"screening.exclusionсriteria.plannedhospitalizationconditiont.color":                                                   1,
			"screening.exclusionсriteria.donorblooddonationcondition.donor_blood_donation":                                         donorBloodDonation,
			"screening.exclusionсriteria.donorblooddonationcondition.color":                                                        1,
			"screening.exclusionсriteria.subjectparticipationinanyotherstudycondition.subject_participation_in_any_other_study":    subjectParticipationInAnyOtherStudy,
			"screening.exclusionсriteria.subjectparticipationinanyotherstudycondition.color":                                       1,
			"screening.exclusionсriteria.anyvaccinationinthelastmonthcondition.any_vaccination_in_the_last_month":                  anyVaccinationInTheLastMonth,
			"screening.exclusionсriteria.anyvaccinationinthelastmonthcondition.color":                                              1,
			"screening.exclusionсriteria.inabilitytoreadinrussiancondition.inability_to_read_in_russian":                           inabilityToReadInRussian,
			"screening.exclusionсriteria.inabilitytoreadinrussiancondition.color":                                                  1,
			"screening.exclusionсriteria.researchcenterstaffcondition.research_center_staff":                                       researchCenterStaff,
			"screening.exclusionсriteria.researchcenterstaffcondition.color":                                                       1,
			"screening.exclusionсriteria.anyotherstateofthesubjectofthestudycondition.any_other_state_of_the_subject_of_the_study": anyOtherStateOfTheSubjectOfTheStudy,
			"screening.exclusionсriteria.anyotherstateofthesubjectofthestudycondition.color":                                       1,
		},
	}

	_, err := collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return err
	}
	return nil
}

func (r *MongoScreeningRepository) UpdateColor(ctx context.Context, id primitive.ObjectID, fieldToUpdate string, field int) error {
	collection := r.store.client.Database("eCRF").Collection("subjects")
	fmt.Println(fieldToUpdate)
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

func (r *MongoScreeningRepository) UpdateColorWithComment(ctx context.Context, id primitive.ObjectID, fieldToUpdate, reason, comment string, color int) error {
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
