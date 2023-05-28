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
			"screening.informaionconsent.signedcondition.signed":                                           isSigned,
			"screening.informaionconsent.signedcondition.color":                                            1,
			"screening.informaionconsent.dateofsigncondition.dateofsign":                                   dateOfSign,
			"screening.informaionconsent.dateofsigncondition.color":                                        1,
			"screening.informaionconsent.timeofsigncondition.timeofsign":                                   timeOfSign,
			"screening.informaionconsent.timeofsigncondition.color":                                        1,
			"screening.informaionconsent.receivedaninsurancepolicycondition.receivedaninsurancepolicy":     receivedAnInsurancePolicy,
			"screening.informaionconsent.receivedaninsurancepolicycondition.color":                         1,
			"screening.informaionconsent.receivedaninformaionconsentcondition.receivedaninformaionconsent": receivedAnInformaionConsent,
			"screening.informaionconsent.receivedaninformaionconsentcondition.color":                       1,
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
			"screening.anthropometry.anthropometricdatabeenmeasuredcondition.anthropometricdatabeenmeasured": anthropometricDataBeenMeasured,
			"screening.anthropometry.anthropometricdatabeenmeasuredcondition.color":                          1,
			"screening.anthropometry.reasonifnotcondition.reasonifnot":                                       reasonIfNot,
			"screening.anthropometry.reasonifnotcondition.color":                                             1,
			"screening.anthropometry.dateofstartmeasuredcondition.dateofstartmeasured":                       dateOfStartMeasured,
			"screening.anthropometry.dateofstartmeasuredcondition.color":                                     1,
			"screening.anthropometry.weightofbodycondition.weightofbody":                                     weightOfBody,
			"screening.anthropometry.weightofbodycondition.color":                                            1,
			"screening.anthropometry.heightofbodycondition.heightofbody":                                     heightOfBody,
			"screening.anthropometry.heightofbodycondition.color":                                            1,
			"screening.anthropometry.indexweigthofbodycondition.indexweightofbody":                           indexWeigthOfBody,
			"screening.anthropometry.indexweigthofbodycondition.color":                                       1,
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
			"screening.inclusioncriteria.presenceofaninformationpanelcondition.presenceofaninformationpanel":                               presenceOfAnInformationPanel,
			"screening.inclusioncriteria.presenceofaninformationpanelcondition.color":                                                      1,
			"screening.inclusioncriteria.aged18to55yearscondition.aged18to55years":                                                         aged18To55Years,
			"screening.inclusioncriteria.aged18to55yearscondition.color":                                                                   1,
			"screening.inclusioncriteria.negativehivtestresultcondition.negativehivtestresult":                                             negativeHIVTestResult,
			"screening.inclusioncriteria.negativehivtestresultcondition.color":                                                             1,
			"screening.inclusioncriteria.bodymassindexcondition.bodymassindex":                                                             bodyMassIndex,
			"screening.inclusioncriteria.bodymassindexcondition.color":                                                                     1,
			"screening.inclusioncriteria.absenceofacuteinfectiousdiseasescondition.absenceofacuteinfectiousdiseases":                       absenceOfAcuteInfectiousDiseases,
			"screening.inclusioncriteria.absenceofacuteinfectiousdiseasescondition.color":                                                  1,
			"screening.inclusioncriteria.consenttouseeffectivemethodsofcontraceptioncondition.consenttouseeffectivemethodsofcontraception": consentToUseEffectiveMethodsOfContraception,
			"screening.inclusioncriteria.consenttouseeffectivemethodsofcontraceptioncondition.color":                                       1,
			"screening.inclusioncriteria.negativepregnancytestcondition.negativepregnancytest":                                             negativePregnancyTest,
			"screening.inclusioncriteria.negativepregnancytestcondition.color":                                                             1,
			"screening.inclusioncriteria.negativedrugtestcondition.negativedrugtest":                                                       negativeDrugTest,
			"screening.inclusioncriteria.negativedrugtestcondition.color":                                                                  1,
			"screening.inclusioncriteria.negativealcoholtestcondition.negativealcoholtest":                                                 negativeAlcoholTest,
			"screening.inclusioncriteria.negativealcoholtestcondition.color":                                                               1,
			"screening.inclusioncriteria.nohistoryofseverepostvaccinationreactionscondition.nohistoryofseverepostvaccinationreactions":     noHistoryOfSeverePostVaccinationReactions,
			"screening.inclusioncriteria.nohistoryofseverepostvaccinationreactionscondition.color":                                         1,
			"screening.inclusioncriteria.indicatorsbloodtestsatscreeningwithincondition.indicatorsbloodtestsatscreeningwithin":             indicatorsBloodTestsAtScreeningWithin,
			"screening.inclusioncriteria.indicatorsbloodtestsatscreeningwithincondition.color":                                             1,
			"screening.inclusioncriteria.nomyocardialchangescondition.nomyocardialchanges":                                                 noMyocardialChanges,
			"screening.inclusioncriteria.nomyocardialchangescondition.color":                                                               1,
			"screening.inclusioncriteria.negativetestresultforcovidcondition.negativetestresultforCOVID":                                   negativeTestResultForCOVID,
			"screening.inclusioncriteria.negativetestresultforcovidcondition.color":                                                        1,
			"screening.inclusioncriteria.nocontraindicationstovaccinationcondition.nocontraindicationstovaccination":                       noContraindicationsToVaccination,
			"screening.inclusioncriteria.nocontraindicationstovaccinationcondition.color":                                                  1,
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
			"screening.exclusionсriteria.lackofsignedinformedconsentcondition.lackofsignedinformedconsent":                     lackOfSignedInformedConsent,
			"screening.exclusionсriteria.lackofsignedinformedconsentcondition.color":                                           1,
			"screening.exclusionсriteria.steroidtherapycondition.steroidtherapy":                                               steroidTherapy,
			"screening.exclusionсriteria.steroidtherapycondition.color":                                                        1,
			"screening.exclusionсriteria.therapywithimmunosuppressivedrugscondition.therapywithimmunosuppressivedrugs":         therapyWithImmunosuppressiveDrugs,
			"screening.exclusionсriteria.therapywithimmunosuppressivedrugscondition.color":                                     1,
			"screening.exclusionсriteria.femalesubjectsduringpregnancycondition.femalesubjectsduringpregnancy":                 femaleSubjectsDuringPregnancy,
			"screening.exclusionсriteria.femalesubjectsduringpregnancycondition.color":                                         1,
			"screening.exclusionсriteria.strokeinlessthanoneyearcondition.strokeinlessthanoneyear":                             strokeInLessThanOneYear,
			"screening.exclusionсriteria.strokeinlessthanoneyearcondition.color":                                               1,
			"screening.exclusionсriteria.chronicsystemicinfectionscondition.chronicsystemicinfections":                         chronicSystemicInfections,
			"screening.exclusionсriteria.chronicsystemicinfectionscondition.color":                                             1,
			"screening.exclusionсriteria.aggravatedallergichistorycondition.aggravatedallergichistory":                         aggravatedAllergicHistory,
			"screening.exclusionсriteria.aggravatedallergichistorycondition.color":                                             1,
			"screening.exclusionсriteria.presenceofahistoryofneoplasmscondition.presenceofahistoryofneoplasms":                 presenceOfAHistoryOfNeoplasms,
			"screening.exclusionсriteria.presenceofahistoryofneoplasmscondition.color":                                         1,
			"screening.exclusionсriteria.historyofsplenectomycondition.historyofsplenectomy":                                   historyOfSplenectomy,
			"screening.exclusionсriteria.historyofsplenectomycondition.color":                                                  1,
			"screening.exclusionсriteria.neutropeniacondition.neutropenia":                                                     neutropenia,
			"screening.exclusionсriteria.neutropeniacondition.color":                                                           1,
			"screening.exclusionсriteria.subjectswithactivesyphiliscondition.subjectswithactivesyphilis":                       subjectsWithActiveSyphilis,
			"screening.exclusionсriteria.subjectswithactivesyphiliscondition.color":                                            1,
			"screening.exclusionсriteria.anorexiacondition.anorexia":                                                           anorexia,
			"screening.exclusionсriteria.anorexiacondition.color":                                                              1,
			"screening.exclusionсriteria.extensivetattooscondition.extensivetattoos":                                           extensiveTattoos,
			"screening.exclusionсriteria.extensivetattooscondition.color":                                                      1,
			"screening.exclusionсriteria.takingnarcoticandpsychostimulantdrugscondition.takingnarcoticandpsychostimulantdrugs": takingNarcoticAndPsychostimulantDrugs,
			"screening.exclusionсriteria.takingnarcoticandpsychostimulantdrugscondition.color":                                 1,
			"screening.exclusionсriteria.smokingmoretthantencigarettesadaycondition.smokingmorethantencigarettesaday":          smokingMoretThanTenCigarettesADay,
			"screening.exclusionсriteria.smokingmoretthantencigarettesadaycondition.color":                                     1,
			"screening.exclusionсriteria.alcoholintakecondition.alcoholintake":                                                 alcoholIntake,
			"screening.exclusionсriteria.alcoholintakecondition.color":                                                         1,
			"screening.exclusionсriteria.plannedhospitalizationconditiont.plannedhospitalization":                              plannedHospitalization,
			"screening.exclusionсriteria.plannedhospitalizationconditiont.color":                                               1,
			"screening.exclusionсriteria.donorblooddonationcondition.donorblooddonation":                                       donorBloodDonation,
			"screening.exclusionсriteria.donorblooddonationcondition.color":                                                    1,
			"screening.exclusionсriteria.subjectparticipationinanyotherstudycondition.subjectparticipationinanyotherstudy":     subjectParticipationInAnyOtherStudy,
			"screening.exclusionсriteria.subjectparticipationinanyotherstudycondition.color":                                   1,
			"screening.exclusionсriteria.anyvaccinationinthelastmonthcondition.anyvaccinationinthelastmonth":                   anyVaccinationInTheLastMonth,
			"screening.exclusionсriteria.anyvaccinationinthelastmonthcondition.color":                                          1,
			"screening.exclusionсriteria.inabilitytoreadinrussiancondition.inabilitytoreadinrussian":                           inabilityToReadInRussian,
			"screening.exclusionсriteria.inabilitytoreadinrussiancondition.color":                                              1,
			"screening.exclusionсriteria.researchcenterstaffcondition.researchcenterstaff":                                     researchCenterStaff,
			"screening.exclusionсriteria.researchcenterstaffcondition.color":                                                   1,
			"screening.exclusionсriteria.anyotherstateofthesubjectofthestudycondition.anyotherstateofthesubjectofthestudy":     anyOtherStateOfTheSubjectOfTheStudy,
			"screening.exclusionсriteria.anyotherstateofthesubjectofthestudycondition.color":                                   1,
		},
	}

	_, err := collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return err
	}
	return nil
}

func (r *MongoScreeningRepository) CompletionOfScreening(
	ctx context.Context,
	id primitive.ObjectID,
	VolunteerEligible,
	NoExclusionCriteria,
	InformedOfTheRestrictions,
	VolunteerIncluded int,
	ReasonIfNot,
	CommentValue string,
) error {
	collection := r.store.client.Database("eCRF").Collection("subjects")
	filter := bson.M{"_id": id}
	update := bson.M{
		"$set": bson.M{
			"screening.completionofscreening.volunteereligiblecondition.volunteereligible":                 VolunteerEligible,
			"screening.completionofscreening.volunteereligiblecondition.color":                             1,
			"screening.completionofscreening.noexclusioncriteriacondition.noexclusioncriteria":             NoExclusionCriteria,
			"screening.completionofscreening.noexclusioncriteriacondition.color":                           1,
			"screening.completionofscreening.informedoftherestrictionscondition.informedoftherestrictions": InformedOfTheRestrictions,
			"screening.completionofscreening.informedoftherestrictionscondition.color":                     1,
			"screening.completionofscreening.volunteerincludedcondition.volunteerincluded":                 VolunteerIncluded,
			"screening.completionofscreening.volunteerincludedcondition.color":                             1,
			"screening.completionofscreening.reasonifnotcondition.reasonifnot":                             ReasonIfNot,
			"screening.completionofscreening.reasonifnotcondition.color":                                   1,
			"screening.completionofscreening.commentcondition.commentvalue":                                CommentValue,
			"screening.completionofscreening.commentcondition.color":                                       1,
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

func (r *MongoScreeningRepository) UpdateFieldIntValue(ctx context.Context, id primitive.ObjectID, fieldToUpdate, fieldValue string, value, color int) error {
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

func (r *MongoScreeningRepository) UpdateFieldStringValue(ctx context.Context, id primitive.ObjectID, fieldToUpdate, fieldValue, value string, color int) error {
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
