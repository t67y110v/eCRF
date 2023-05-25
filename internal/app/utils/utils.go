package utils

import (
	model "github.com/t67y110v/web/internal/app/model/subject"
)

func GetUserRole(roleId int) string {
	m := make(map[int]string)
	m[1] = "Администратор"
	m[2] = "Исследователь"
	m[3] = "Главный исследователь"
	m[4] = "Монитор"
	m[5] = "Дата Менеджер"
	m[6] = "Аудитор Контроля Качества"
	m[7] = "Медицинский монитор"
	m[8] = "Специалист по фармаконадзору"
	return m[roleId]
}

func GetFieldName(field string) string {
	m := make(map[string]string)
	m["signed"] = "screening.informaionconsent.signedcondition."
	m["date_of_sign"] = "screening.informaionconsent.dateofsigncondition."
	m["time_of_sign"] = "screening.informaionconsent.timeofsigncondition."
	m["original"] = "screening.informaionconsent.receivedaninsurancepolicycondition."
	m["consent"] = "screening.informaionconsent.receivedaninformaionconsentcondition."

	m["sex"] = "screening.demography.sexcondition."
	m["race"] = "screening.demography.racecondition."
	m["birth_date"] = "screening.demography.birthdatecondition."

	m["data_been_measured"] = "screening.anthropometry.anthropometricdatabeenmeasuredcondition."
	m["if_not"] = "screening.anthropometry.reasonifnotcondition."
	m["date_of_start"] = "screening.anthropometry.dateofstartmeasuredcondition."
	m["weight"] = "screening.anthropometry.weightofbodycondition."
	m["height"] = "screening.anthropometry.heightofbodycondition."
	m["index"] = "screening.anthropometry.indexweigthofbodycondition."

	m["inclusion1"] = "screening.inclusioncriteria.presenceofaninformationpanelcondition."
	m["inclusion2"] = "screening.inclusioncriteria.aged18to55yearscondition."
	m["inclusion3"] = "screening.inclusioncriteria.negativehivtestresultcondition."
	m["inclusion4"] = "screening.inclusioncriteria.bodymassindexcondition."
	m["inclusion5"] = "screening.inclusioncriteria.absenceofacuteinfectiousdiseasescondition."
	m["inclusion6"] = "screening.inclusioncriteria.consenttouseeffectivemethodsofcontraceptioncondition."
	m["inclusion7"] = "screening.inclusioncriteria.negativepregnancytestcondition."
	m["inclusion8"] = "screening.inclusioncriteria.negativedrugtestcondition."
	m["inclusion9"] = "screening.inclusioncriteria.negativealcoholtestcondition."
	m["inclusion10"] = "screening.inclusioncriteria.nohistoryofseverepostvaccinationreactionscondition."
	m["inclusion11"] = "screening.inclusioncriteria.indicatorsbloodtestsatscreeningwithincondition."
	m["inclusion12"] = "screening.inclusioncriteria.nomyocardialchangescondition."
	m["inclusion13"] = "screening.inclusioncriteria.negativetestresultforcovidcondition."
	m["inclusion14"] = "screening.inclusioncriteria.nocontraindicationstovaccinationcondition."

	m["exclusion1"] = "screening.exclusionсriteria.lackofsignedinformedconsentcondition."
	m["exclusion2"] = "screening.exclusionсriteria.steroidtherapycondition."
	m["exclusion3"] = "screening.exclusionсriteria.therapywithimmunosuppressivedrugscondition."
	m["exclusion4"] = "screening.exclusionсriteria.femalesubjectsduringpregnancycondition."
	m["exclusion5"] = "screening.exclusionсriteria.strokeinlessthanoneyearcondition."
	m["exclusion6"] = "screening.exclusionсriteria.chronicsystemicinfectionscondition."
	m["exclusion7"] = "screening.exclusionсriteria.aggravatedallergichistorycondition."
	m["exclusion8"] = "screening.exclusionсriteria.presenceofahistoryofneoplasmscondition."
	m["exclusion9"] = "screening.exclusionсriteria.historyofsplenectomycondition."
	m["exclusion10"] = "screening.exclusionсriteria.neutropeniacondition."
	m["exclusion11"] = "screening.exclusionсriteria.subjectswithactivesyphiliscondition."
	m["exclusion12"] = "screening.exclusionсriteria.anorexiacondition."
	m["exclusion13"] = "screening.exclusionсriteria.extensivetattooscondition."
	m["exclusion14"] = "screening.exclusionсriteria.takingnarcoticandpsychostimulantdrugscondition."
	m["exclusion15"] = "screening.exclusionсriteria.smokingmoretthantencigarettesadaycondition."
	m["exclusion16"] = "screening.exclusionсriteria.alcoholintakecondition."
	m["exclusion17"] = "screening.exclusionсriteria.plannedhospitalizationconditiont."
	m["exclusion18"] = "screening.exclusionсriteria.donorblooddonationcondition."
	m["exclusion19"] = "screening.exclusionсriteria.subjectparticipationinanyotherstudycondition."
	m["exclusion20"] = "screening.exclusionсriteria.anyvaccinationinthelastmonthcondition."
	m["exclusion21"] = "screening.exclusionсriteria.inabilitytoreadinrussiancondition."
	m["exclusion22"] = "screening.exclusionсriteria.researchcenterstaffcondition."
	m["exclusion23"] = "screening.exclusionсriteria.anyotherstateofthesubjectofthestudycondition."

	m["completion1"] = "screening.completionofscreening.volunteereligiblecondition."
	m["completion2"] = "screening.completionofscreening.noexclusioncriteriacondition."
	m["completion3"] = "screening.completionofscreening.informedoftherestrictionscondition."
	m["completion4"] = "screening.completionofscreening.volunteerincludedcondition."
	m["completion5"] = "screening.completionofscreening.reasonifnotcondition."
	m["completion6"] = "screening.completionofscreening.commentcondition."
	return m[field]
}

func GetInformationConsentErrors(subject *model.Subject) []*model.InformationConsentErrors {
	var errors []*model.InformationConsentErrors
	if subject.Screening.InformaionConsent.SignedCondition.Color == 3 || subject.Screening.InformaionConsent.SignedCondition.Color == 4 {
		errors = append(errors, &model.InformationConsentErrors{
			Field:    "Информированное согласие подписано?",
			Reasons:  subject.Screening.InformaionConsent.SignedCondition.Reason,
			Comments: subject.Screening.InformaionConsent.SignedCondition.Comment})
	}
	if subject.Screening.InformaionConsent.DateOfSignCondition.Color == 3 || subject.Screening.InformaionConsent.DateOfSignCondition.Color == 4 {
		errors = append(errors, &model.InformationConsentErrors{
			Field:    "Дата подписания",
			Reasons:  subject.Screening.InformaionConsent.DateOfSignCondition.Reason,
			Comments: subject.Screening.InformaionConsent.DateOfSignCondition.Comment})

	}
	if subject.Screening.InformaionConsent.TimeOfSignCondition.Color == 3 || subject.Screening.InformaionConsent.TimeOfSignCondition.Color == 4 {
		errors = append(errors, &model.InformationConsentErrors{
			Field:    "Время подписания",
			Reasons:  subject.Screening.InformaionConsent.TimeOfSignCondition.Reason,
			Comments: subject.Screening.InformaionConsent.TimeOfSignCondition.Comment})
	}
	if subject.Screening.InformaionConsent.ReceivedAnInsurancePolicyCondition.Color == 3 || subject.Screening.InformaionConsent.ReceivedAnInsurancePolicyCondition.Color == 4 {
		errors = append(errors, &model.InformationConsentErrors{
			Field:    "Доброволец получил на руки оригинал страхового Полиса участника исследования ?",
			Reasons:  subject.Screening.InformaionConsent.ReceivedAnInsurancePolicyCondition.Reason,
			Comments: subject.Screening.InformaionConsent.ReceivedAnInsurancePolicyCondition.Comment})
	}
	if subject.Screening.InformaionConsent.ReceivedAnInformaionConsentCondition.Color == 3 || subject.Screening.InformaionConsent.ReceivedAnInformaionConsentCondition.Color == 4 {
		errors = append(errors, &model.InformationConsentErrors{
			Field:    "Доброволец получил на руки один экземпляр информированного согласия",
			Reasons:  subject.Screening.InformaionConsent.ReceivedAnInformaionConsentCondition.Reason,
			Comments: subject.Screening.InformaionConsent.ReceivedAnInformaionConsentCondition.Comment})
	}

	return errors
}

func GetDemographyErrorsErrors(subject *model.Subject) []*model.InformationConsentErrors {
	var errors []*model.InformationConsentErrors
	if subject.Screening.Demography.SexCondition.Color == 3 || subject.Screening.Demography.SexCondition.Color == 4 {
		errors = append(errors, &model.InformationConsentErrors{
			Field:    "Пол",
			Reasons:  subject.Screening.Demography.SexCondition.Reason,
			Comments: subject.Screening.Demography.SexCondition.Comment})
	}
	if subject.Screening.Demography.RaceCondition.Color == 3 || subject.Screening.Demography.RaceCondition.Color == 4 {
		errors = append(errors, &model.InformationConsentErrors{
			Field:    "Раса",
			Reasons:  subject.Screening.Demography.RaceCondition.Reason,
			Comments: subject.Screening.Demography.RaceCondition.Comment})

	}
	if subject.Screening.Demography.BirthDateCondition.Color == 3 || subject.Screening.Demography.BirthDateCondition.Color == 4 {
		errors = append(errors, &model.InformationConsentErrors{
			Field:    "Дата рождения",
			Reasons:  subject.Screening.Demography.BirthDateCondition.Reason,
			Comments: subject.Screening.Demography.BirthDateCondition.Comment})
	}

	return errors
}

func GetAnthropometryErrors(subject *model.Subject) []*model.InformationConsentErrors {
	var errors []*model.InformationConsentErrors
	if subject.Screening.Anthropometry.AnthropometricDataBeenMeasuredCondition.Color == 3 || subject.Screening.Anthropometry.AnthropometricDataBeenMeasuredCondition.Color == 4 {
		errors = append(errors, &model.InformationConsentErrors{
			Field:    "Измерение антропометрических данных проведено?",
			Reasons:  subject.Screening.Anthropometry.AnthropometricDataBeenMeasuredCondition.Reason,
			Comments: subject.Screening.Anthropometry.AnthropometricDataBeenMeasuredCondition.Comment})
	}
	if subject.Screening.Anthropometry.ReasonIfNotCondition.Color == 3 || subject.Screening.Anthropometry.ReasonIfNotCondition.Color == 4 {
		errors = append(errors, &model.InformationConsentErrors{
			Field:    "Если НЕТ, то укажите причину",
			Reasons:  subject.Screening.Anthropometry.ReasonIfNotCondition.Reason,
			Comments: subject.Screening.Anthropometry.ReasonIfNotCondition.Comment})

	}
	if subject.Screening.Anthropometry.DateOfStartMeasuredCondition.Color == 3 || subject.Screening.Anthropometry.DateOfStartMeasuredCondition.Color == 4 {
		errors = append(errors, &model.InformationConsentErrors{
			Field:    "Дата проведения измерений",
			Reasons:  subject.Screening.Anthropometry.DateOfStartMeasuredCondition.Reason,
			Comments: subject.Screening.Anthropometry.DateOfStartMeasuredCondition.Comment})
	}
	if subject.Screening.Anthropometry.WeightOfBodyCondition.Color == 3 || subject.Screening.Anthropometry.WeightOfBodyCondition.Color == 4 {
		errors = append(errors, &model.InformationConsentErrors{
			Field:    "Масса тела",
			Reasons:  subject.Screening.Anthropometry.WeightOfBodyCondition.Reason,
			Comments: subject.Screening.Anthropometry.WeightOfBodyCondition.Comment})
	}
	if subject.Screening.Anthropometry.HeightOfBodyCondition.Color == 3 || subject.Screening.Anthropometry.HeightOfBodyCondition.Color == 4 {
		errors = append(errors, &model.InformationConsentErrors{
			Field:    "Рост",
			Reasons:  subject.Screening.Anthropometry.HeightOfBodyCondition.Reason,
			Comments: subject.Screening.Anthropometry.HeightOfBodyCondition.Comment})

	}
	if subject.Screening.Anthropometry.IndexWeigthOfBodyCondition.Color == 3 || subject.Screening.Anthropometry.IndexWeigthOfBodyCondition.Color == 4 {
		errors = append(errors, &model.InformationConsentErrors{
			Field:    "Индекс массы тела(рассчетное значение)",
			Reasons:  subject.Screening.Anthropometry.IndexWeigthOfBodyCondition.Reason,
			Comments: subject.Screening.Anthropometry.IndexWeigthOfBodyCondition.Comment})
	}
	return errors
}

func GetInclusionErrors(subject *model.Subject) []*model.InformationConsentErrors {
	var errors []*model.InformationConsentErrors
	if subject.Screening.InclusionCriteria.PresenceOfAnInformationPanelCondition.Color == 3 || subject.Screening.InclusionCriteria.PresenceOfAnInformationPanelCondition.Color == 4 {
		errors = append(errors, &model.InformationConsentErrors{
			Field:    "Наличие письменного информированного согласия субъекта на участие в исследовании  ",
			Reasons:  subject.Screening.InclusionCriteria.PresenceOfAnInformationPanelCondition.Reason,
			Comments: subject.Screening.InclusionCriteria.PresenceOfAnInformationPanelCondition.Comment})
	}
	if subject.Screening.InclusionCriteria.Aged18To55YearsCondition.Color == 3 || subject.Screening.InclusionCriteria.Aged18To55YearsCondition.Color == 4 {
		errors = append(errors, &model.InformationConsentErrors{
			Field:    "Мужчины и женщины в возрасте от 18 до 55 лет",
			Reasons:  subject.Screening.InclusionCriteria.Aged18To55YearsCondition.Reason,
			Comments: subject.Screening.InclusionCriteria.Aged18To55YearsCondition.Comment})
	}
	if subject.Screening.InclusionCriteria.NegativeHIVTestResultCondition.Color == 3 || subject.Screening.InclusionCriteria.NegativeHIVTestResultCondition.Color == 4 {
		errors = append(errors, &model.InformationConsentErrors{
			Field:    "Отрицательный результат исследования на ВИЧ, гепатиты, сифилис",
			Reasons:  subject.Screening.InclusionCriteria.NegativeHIVTestResultCondition.Reason,
			Comments: subject.Screening.InclusionCriteria.NegativeHIVTestResultCondition.Comment})
	}
	if subject.Screening.InclusionCriteria.BodyMassIndexCondition.Color == 3 || subject.Screening.InclusionCriteria.BodyMassIndexCondition.Color == 4 {
		errors = append(errors, &model.InformationConsentErrors{
			Field:    "Индекс массы тела (ИМТ) составляет 18.5≤ИМТ≤30  ",
			Reasons:  subject.Screening.InclusionCriteria.BodyMassIndexCondition.Reason,
			Comments: subject.Screening.InclusionCriteria.BodyMassIndexCondition.Comment})
	}
	if subject.Screening.InclusionCriteria.AbsenceOfAcuteInfectiousDiseasesCondition.Color == 3 || subject.Screening.InclusionCriteria.AbsenceOfAcuteInfectiousDiseasesCondition.Color == 4 {
		errors = append(errors, &model.InformationConsentErrors{
			Field:    "Отсутствие острых инфекционных и/или респираторных заболеваний по меньшей мере в течение 14-ти дней до включения в исследование",
			Reasons:  subject.Screening.InclusionCriteria.ConsentToUseEffectiveMethodsOfContraceptionCondition.Reason,
			Comments: subject.Screening.InclusionCriteria.ConsentToUseEffectiveMethodsOfContraceptionCondition.Comment})
	}
	if subject.Screening.InclusionCriteria.ConsentToUseEffectiveMethodsOfContraceptionCondition.Color == 3 || subject.Screening.InclusionCriteria.ConsentToUseEffectiveMethodsOfContraceptionCondition.Color == 4 {
		errors = append(errors, &model.InformationConsentErrors{
			Field:    "Согласие на использование эффективных методов контрацепции в ходе всего периода участия в исследовании",
			Reasons:  subject.Screening.InclusionCriteria.ConsentToUseEffectiveMethodsOfContraceptionCondition.Reason,
			Comments: subject.Screening.InclusionCriteria.ConsentToUseEffectiveMethodsOfContraceptionCondition.Comment})
	}
	if subject.Screening.InclusionCriteria.NegativePregnancyTestCondition.Color == 3 || subject.Screening.InclusionCriteria.NegativePregnancyTestCondition.Color == 4 {
		errors = append(errors, &model.InformationConsentErrors{
			Field:    "Отрицательный тест на беременность по результатам исследования мочи на визите скрининга (для женщин детородного возраста)",
			Reasons:  subject.Screening.InclusionCriteria.NegativePregnancyTestCondition.Reason,
			Comments: subject.Screening.InclusionCriteria.NegativePregnancyTestCondition.Comment})
	}
	if subject.Screening.InclusionCriteria.NegativeDrugTestCondition.Color == 3 || subject.Screening.InclusionCriteria.NegativeDrugTestCondition.Color == 4 {
		errors = append(errors, &model.InformationConsentErrors{
			Field:    "Отрицательный тест на наличие наркотических и психостимулирующих средств в моче на визите скрининга",
			Reasons:  subject.Screening.InclusionCriteria.NegativeDrugTestCondition.Reason,
			Comments: subject.Screening.InclusionCriteria.NegativeDrugTestCondition.Comment})
	}
	if subject.Screening.InclusionCriteria.NegativeAlcoholTestCondition.Color == 3 || subject.Screening.InclusionCriteria.NegativeAlcoholTestCondition.Color == 4 {
		errors = append(errors, &model.InformationConsentErrors{
			Field:    "Отрицательный тест на содержание алкоголя на визите скрининга",
			Reasons:  subject.Screening.InclusionCriteria.NegativeAlcoholTestCondition.Reason,
			Comments: subject.Screening.InclusionCriteria.NegativeAlcoholTestCondition.Comment})
	}
	if subject.Screening.InclusionCriteria.NoHistoryOfSeverePostVaccinationReactionsCondition.Color == 3 || subject.Screening.InclusionCriteria.NoHistoryOfSeverePostVaccinationReactionsCondition.Color == 4 {
		errors = append(errors, &model.InformationConsentErrors{
			Field:    "Отсутствие в анамнезе выраженных поствакцинальных реакций или поствакцинальных осложнений на предыдущее применение иммунобиологических препаратов ",
			Reasons:  subject.Screening.InclusionCriteria.NoHistoryOfSeverePostVaccinationReactionsCondition.Reason,
			Comments: subject.Screening.InclusionCriteria.NoHistoryOfSeverePostVaccinationReactionsCondition.Comment})
	}
	if subject.Screening.InclusionCriteria.IndicatorsBloodTestsAtScreeningWithinCondition.Color == 3 || subject.Screening.InclusionCriteria.IndicatorsBloodTestsAtScreeningWithinCondition.Color == 4 {
		errors = append(errors, &model.InformationConsentErrors{
			Field:    "Показатели общего и биохимического анализа крови на скрининге в пределах 1,1 х ВГРИ – 0,9 х НГРИ",
			Reasons:  subject.Screening.InclusionCriteria.IndicatorsBloodTestsAtScreeningWithinCondition.Reason,
			Comments: subject.Screening.InclusionCriteria.IndicatorsBloodTestsAtScreeningWithinCondition.Comment})
	}
	if subject.Screening.InclusionCriteria.NoMyocardialChangesCondition.Color == 3 || subject.Screening.InclusionCriteria.NoMyocardialChangesCondition.Color == 4 {
		errors = append(errors, &model.InformationConsentErrors{
			Field:    "Отсутствие изменений миокарда воспалительного или дистрофического характера по результатам ЭКГ на скрининге",
			Reasons:  subject.Screening.InclusionCriteria.NoMyocardialChangesCondition.Reason,
			Comments: subject.Screening.InclusionCriteria.NoMyocardialChangesCondition.Comment})
	}
	if subject.Screening.InclusionCriteria.NegativeTestResultForCOVIDCondition.Color == 3 || subject.Screening.InclusionCriteria.NegativeTestResultForCOVIDCondition.Color == 4 {
		errors = append(errors, &model.InformationConsentErrors{
			Field:    "Отрицательный результат исследования на COVID-2019, определяемый методом ПЦР или экспресс-методом за 3 дня до включения в исследование",
			Reasons:  subject.Screening.InclusionCriteria.NegativeTestResultForCOVIDCondition.Reason,
			Comments: subject.Screening.InclusionCriteria.NegativeTestResultForCOVIDCondition.Comment})
	}
	if subject.Screening.InclusionCriteria.NoContraindicationsToVaccinationCondition.Color == 3 || subject.Screening.InclusionCriteria.NoContraindicationsToVaccinationCondition.Color == 4 {
		errors = append(errors, &model.InformationConsentErrors{
			Field:    "Отсутствие противопоказаний к вакцинации",
			Reasons:  subject.Screening.InclusionCriteria.NoContraindicationsToVaccinationCondition.Reason,
			Comments: subject.Screening.InclusionCriteria.NoContraindicationsToVaccinationCondition.Comment})
	}
	return errors
}

func GetExclusionErrors(subject *model.Subject) []*model.InformationConsentErrors {

	var errors []*model.InformationConsentErrors
	if subject.Screening.ExclusionСriteria.LackOfSignedInformedConsentCondition.Color == 3 || subject.Screening.ExclusionСriteria.LackOfSignedInformedConsentCondition.Color == 4 {
		errors = append(errors, &model.InformationConsentErrors{
			Field:    "Отсутствие подписанного информированного согласия",
			Reasons:  subject.Screening.ExclusionСriteria.LackOfSignedInformedConsentCondition.Reason,
			Comments: subject.Screening.ExclusionСriteria.LackOfSignedInformedConsentCondition.Comment})
	}
	if subject.Screening.ExclusionСriteria.SteroidTherapyCondition.Color == 3 || subject.Screening.ExclusionСriteria.SteroidTherapyCondition.Color == 4 {
		errors = append(errors, &model.InformationConsentErrors{
			Field:    "Терапия стероидами (за исключением гормональных контрацептивных препаратов) и/или иммуноглобулинами или другими препаратами крови, не завершившаяся за 30 дней до включения в исследование",
			Reasons:  subject.Screening.ExclusionСriteria.SteroidTherapyCondition.Reason,
			Comments: subject.Screening.ExclusionСriteria.SteroidTherapyCondition.Comment})
	}

	if subject.Screening.ExclusionСriteria.TherapyWithImmunosuppressiveDrugsCondition.Color == 3 || subject.Screening.ExclusionСriteria.TherapyWithImmunosuppressiveDrugsCondition.Color == 4 {
		errors = append(errors, &model.InformationConsentErrors{
			Field:    "Терапия иммуносупрессивными препаратами, завершившаяся менее чем за 3 месяца до включения в исследование",
			Reasons:  subject.Screening.ExclusionСriteria.TherapyWithImmunosuppressiveDrugsCondition.Reason,
			Comments: subject.Screening.ExclusionСriteria.TherapyWithImmunosuppressiveDrugsCondition.Comment})
	}

	if subject.Screening.ExclusionСriteria.FemaleSubjectsDuringPregnancyCondition.Color == 3 || subject.Screening.ExclusionСriteria.FemaleSubjectsDuringPregnancyCondition.Color == 4 {
		errors = append(errors, &model.InformationConsentErrors{
			Field:    "Субъекты женского пола в период беременности или кормления грудью",
			Reasons:  subject.Screening.ExclusionСriteria.FemaleSubjectsDuringPregnancyCondition.Reason,
			Comments: subject.Screening.ExclusionСriteria.FemaleSubjectsDuringPregnancyCondition.Comment})
	}

	if subject.Screening.ExclusionСriteria.StrokeInLessThanOneYearCondition.Color == 3 || subject.Screening.ExclusionСriteria.StrokeInLessThanOneYearCondition.Color == 4 {
		errors = append(errors, &model.InformationConsentErrors{
			Field:    "Перенесенный менее чем за один год до включения в исследование острый коронарный синдром или инсульт",
			Reasons:  subject.Screening.ExclusionСriteria.StrokeInLessThanOneYearCondition.Reason,
			Comments: subject.Screening.ExclusionСriteria.StrokeInLessThanOneYearCondition.Comment})
	}

	if subject.Screening.ExclusionСriteria.ChronicSystemicInfectionsCondition.Color == 3 || subject.Screening.ExclusionСriteria.ChronicSystemicInfectionsCondition.Color == 4 {
		errors = append(errors, &model.InformationConsentErrors{
			Field:    "Туберкулез, хронические системные инфекции ",
			Reasons:  subject.Screening.ExclusionСriteria.ChronicSystemicInfectionsCondition.Reason,
			Comments: subject.Screening.ExclusionСriteria.ChronicSystemicInfectionsCondition.Comment})
	}

	if subject.Screening.ExclusionСriteria.AggravatedAllergicHistoryCondition.Color == 3 || subject.Screening.ExclusionСriteria.AggravatedAllergicHistoryCondition.Color == 4 {
		errors = append(errors, &model.InformationConsentErrors{
			Field:    "Отягощенный аллергологический анамнез (наличие в анамнезе сведений об анафилактическом шоке, отеке Квинке, полиморфной экссудативной экземе, сывороточной болезни), гиперчувствительность или аллергические реакции на введение иммунобиологических препаратов, известные аллергические реакции на компоненты препарата, обострение аллергических заболеваний на день включения в исследование ",
			Reasons:  subject.Screening.ExclusionСriteria.AggravatedAllergicHistoryCondition.Reason,
			Comments: subject.Screening.ExclusionСriteria.AggravatedAllergicHistoryCondition.Comment})
	}

	if subject.Screening.ExclusionСriteria.PresenceOfAHistoryOfNeoplasmsCondition.Color == 3 || subject.Screening.ExclusionСriteria.PresenceOfAHistoryOfNeoplasmsCondition.Color == 4 {
		errors = append(errors, &model.InformationConsentErrors{
			Field:    "Наличие в анамнезе новообразований (коды МКБ C00-D09)",
			Reasons:  subject.Screening.ExclusionСriteria.PresenceOfAHistoryOfNeoplasmsCondition.Reason,
			Comments: subject.Screening.ExclusionСriteria.PresenceOfAHistoryOfNeoplasmsCondition.Comment})
	}

	if subject.Screening.ExclusionСriteria.HistoryOfSplenectomyCondition.Color == 3 || subject.Screening.ExclusionСriteria.HistoryOfSplenectomyCondition.Color == 4 {
		errors = append(errors, &model.InformationConsentErrors{
			Field:    "Спленэктомия в анамнезе",
			Reasons:  subject.Screening.ExclusionСriteria.HistoryOfSplenectomyCondition.Reason,
			Comments: subject.Screening.ExclusionСriteria.HistoryOfSplenectomyCondition.Comment})
	}

	if subject.Screening.ExclusionСriteria.NeutropeniaCondition.Color == 3 || subject.Screening.ExclusionСriteria.NeutropeniaCondition.Color == 4 {
		errors = append(errors, &model.InformationConsentErrors{
			Field:    "Нейтропения (снижение абсолютного числа нейтрофилов менее 1000/мм3), агранулоцитоз, значительная кровопотеря, тяжелая анемия (гемоглобин менее 80 г/л), иммунодефицит в анамнезе в течение 6 месяцев до включения в исследование ",
			Reasons:  subject.Screening.ExclusionСriteria.NeutropeniaCondition.Reason,
			Comments: subject.Screening.ExclusionСriteria.NeutropeniaCondition.Comment})
	}

	if subject.Screening.ExclusionСriteria.SubjectsWithActiveSyphilisCondition.Color == 3 || subject.Screening.ExclusionСriteria.SubjectsWithActiveSyphilisCondition.Color == 4 {
		errors = append(errors, &model.InformationConsentErrors{
			Field:    "Субъекты с активной формой сифилиса, ВИЧ/СПИД, гепатиты B и C",
			Reasons:  subject.Screening.ExclusionСriteria.SubjectsWithActiveSyphilisCondition.Reason,
			Comments: subject.Screening.ExclusionСriteria.SubjectsWithActiveSyphilisCondition.Comment})
	}

	if subject.Screening.ExclusionСriteria.AnorexiaCondition.Color == 3 || subject.Screening.ExclusionСriteria.AnorexiaCondition.Color == 4 {
		errors = append(errors, &model.InformationConsentErrors{
			Field:    "Анорексия, белковый дефицит любого происхождения",
			Reasons:  subject.Screening.ExclusionСriteria.AnorexiaCondition.Reason,
			Comments: subject.Screening.ExclusionСriteria.AnorexiaCondition.Comment})
	}

	if subject.Screening.ExclusionСriteria.ExtensiveTattoosCondition.Color == 3 || subject.Screening.ExclusionСriteria.ExtensiveTattoosCondition.Color == 4 {
		errors = append(errors, &model.InformationConsentErrors{
			Field:    "Обширные татуировки на местах введения препарата (область дельтовидной мышцы), не позволяющие оценить местную реакцию на введение ИЛП",
			Reasons:  subject.Screening.ExclusionСriteria.ExtensiveTattoosCondition.Reason,
			Comments: subject.Screening.ExclusionСriteria.ExtensiveTattoosCondition.Comment})
	}

	if subject.Screening.ExclusionСriteria.TakingNarcoticAndPsychostimulantDrugsCondition.Color == 3 || subject.Screening.ExclusionСriteria.TakingNarcoticAndPsychostimulantDrugsCondition.Color == 4 {
		errors = append(errors, &model.InformationConsentErrors{
			Field:    "Прием наркотических и психостимулирующих препаратов в настоящее время или в анамнезе",
			Reasons:  subject.Screening.ExclusionСriteria.TakingNarcoticAndPsychostimulantDrugsCondition.Reason,
			Comments: subject.Screening.ExclusionСriteria.TakingNarcoticAndPsychostimulantDrugsCondition.Comment})
	}

	if subject.Screening.ExclusionСriteria.SmokingMoretThanTenCigarettesADayCondition.Color == 3 || subject.Screening.ExclusionСriteria.SmokingMoretThanTenCigarettesADayCondition.Color == 4 {
		errors = append(errors, &model.InformationConsentErrors{
			Field:    "Курение: более 10 сигарет в день",
			Reasons:  subject.Screening.ExclusionСriteria.SmokingMoretThanTenCigarettesADayCondition.Reason,
			Comments: subject.Screening.ExclusionСriteria.SmokingMoretThanTenCigarettesADayCondition.Comment})
	}
	if subject.Screening.ExclusionСriteria.AlcoholIntakeCondition.Color == 3 || subject.Screening.ExclusionСriteria.AlcoholIntakeCondition.Color == 4 {
		errors = append(errors, &model.InformationConsentErrors{
			Field:    "Прием алкоголя превышающий уровень низкого риска: не более 20 граммов чистого алкоголя в день, не более 5 дней в неделю, прием алкоголя в течение 48 часов до введения исследуемого препарата",
			Reasons:  subject.Screening.ExclusionСriteria.AlcoholIntakeCondition.Reason,
			Comments: subject.Screening.ExclusionСriteria.AlcoholIntakeCondition.Comment})
	}

	if subject.Screening.ExclusionСriteria.PlannedHospitalizationConditiont.Color == 3 || subject.Screening.ExclusionСriteria.PlannedHospitalizationConditiont.Color == 4 {
		errors = append(errors, &model.InformationConsentErrors{
			Field:    "Плановая госпитализация и/или хирургическое вмешательство в период участия в исследовании, а также за 4 недели до предполагаемой даты вакцинации",
			Reasons:  subject.Screening.ExclusionСriteria.PlannedHospitalizationConditiont.Reason,
			Comments: subject.Screening.ExclusionСriteria.PlannedHospitalizationConditiont.Comment})
	}
	if subject.Screening.ExclusionСriteria.DonorBloodDonationCondition.Color == 3 || subject.Screening.ExclusionСriteria.DonorBloodDonationCondition.Color == 4 {
		errors = append(errors, &model.InformationConsentErrors{
			Field:    "Донорская сдача крови (450 мл и более крови или плазмы) менее чем за 2 месяца до начала исследования ",
			Reasons:  subject.Screening.ExclusionСriteria.DonorBloodDonationCondition.Reason,
			Comments: subject.Screening.ExclusionСriteria.DonorBloodDonationCondition.Comment})
	}

	if subject.Screening.ExclusionСriteria.SubjectParticipationInAnyOtherStudyCondition.Color == 3 || subject.Screening.ExclusionСriteria.SubjectParticipationInAnyOtherStudyCondition.Color == 4 {
		errors = append(errors, &model.InformationConsentErrors{
			Field:    "Участие субъекта в любом другом интервенционном клиническом исследовании за последние 90 дней",
			Reasons:  subject.Screening.ExclusionСriteria.SubjectParticipationInAnyOtherStudyCondition.Reason,
			Comments: subject.Screening.ExclusionСriteria.SubjectParticipationInAnyOtherStudyCondition.Comment})
	}

	if subject.Screening.ExclusionСriteria.AnyVaccinationInTheLastMonthCondition.Color == 3 || subject.Screening.ExclusionСriteria.AnyVaccinationInTheLastMonthCondition.Color == 4 {
		errors = append(errors, &model.InformationConsentErrors{
			Field:    "Любая вакцинация за последние 30 дней",
			Reasons:  subject.Screening.ExclusionСriteria.AnyVaccinationInTheLastMonthCondition.Reason,
			Comments: subject.Screening.ExclusionСriteria.AnyVaccinationInTheLastMonthCondition.Comment})
	}

	if subject.Screening.ExclusionСriteria.InabilityToReadInRussianCondition.Color == 3 || subject.Screening.ExclusionСriteria.InabilityToReadInRussianCondition.Color == 4 {
		errors = append(errors, &model.InformationConsentErrors{
			Field:    "Неспособность читать на русском языке; невозможность или нежелание понять суть исследования. Любые другие состояния, которые ограничивают правомерность получения информированного согласия или могут повлиять на способность добровольца принять участие в исследовании повлиять на способность добровольца принять участие в исследовании ",
			Reasons:  subject.Screening.ExclusionСriteria.InabilityToReadInRussianCondition.Reason,
			Comments: subject.Screening.ExclusionСriteria.InabilityToReadInRussianCondition.Comment})
	}

	if subject.Screening.ExclusionСriteria.ResearchCenterStaffCondition.Color == 3 || subject.Screening.ExclusionСriteria.ResearchCenterStaffCondition.Color == 4 {
		errors = append(errors, &model.InformationConsentErrors{
			Field:    "Персонал исследовательских центров и другие сотрудники, непосредственно участвующие в проведении исследования и члены их семей (главный исследователь и члены исследовательской команды)",
			Reasons:  subject.Screening.ExclusionСriteria.ResearchCenterStaffCondition.Reason,
			Comments: subject.Screening.ExclusionСriteria.ResearchCenterStaffCondition.Comment})
	}

	if subject.Screening.ExclusionСriteria.LackOfSignedInformedConsentCondition.Color == 3 || subject.Screening.ExclusionСriteria.LackOfSignedInformedConsentCondition.Color == 4 {
		errors = append(errors, &model.InformationConsentErrors{
			Field:    "Любое иное состояние субъекта исследования, которое, по мнению-врача- исследователя, может препятствовать завершению исследования в соответствии с протоколом ",
			Reasons:  subject.Screening.ExclusionСriteria.AnyOtherStateOfTheSubjectOfTheStudyCondition.Reason,
			Comments: subject.Screening.ExclusionСriteria.AnyOtherStateOfTheSubjectOfTheStudyCondition.Comment})

	}

	return errors

}

func GetCompletionErrors(subject *model.Subject) []*model.InformationConsentErrors {
	var errors []*model.InformationConsentErrors
	if subject.Screening.CompletionOfScreening.VolunteerEligibleCondition.Color == 3 || subject.Screening.CompletionOfScreening.VolunteerEligibleCondition.Color == 4 {
		errors = append(errors, &model.InformationConsentErrors{
			Field:    "Доброволец соответствует критериям включения? ",
			Reasons:  subject.Screening.CompletionOfScreening.VolunteerEligibleCondition.Reason,
			Comments: subject.Screening.CompletionOfScreening.VolunteerEligibleCondition.Comment})
	}
	if subject.Screening.CompletionOfScreening.NoExclusionCriteriaCondition.Color == 3 || subject.Screening.CompletionOfScreening.NoExclusionCriteriaCondition.Color == 4 {
		errors = append(errors, &model.InformationConsentErrors{
			Field:    "Критерии невключения отсутствуют?",
			Reasons:  subject.Screening.CompletionOfScreening.NoExclusionCriteriaCondition.Reason,
			Comments: subject.Screening.CompletionOfScreening.NoExclusionCriteriaCondition.Comment})

	}
	if subject.Screening.CompletionOfScreening.InformedOfTheRestrictionsCondition.Color == 3 || subject.Screening.CompletionOfScreening.InformedOfTheRestrictionsCondition.Color == 4 {
		errors = append(errors, &model.InformationConsentErrors{
			Field:    "Доброволец был проинформирован об ограничениях, требуемых протоколом и согласен их соблюдать в течение исследования?",
			Reasons:  subject.Screening.CompletionOfScreening.InformedOfTheRestrictionsCondition.Reason,
			Comments: subject.Screening.CompletionOfScreening.InformedOfTheRestrictionsCondition.Comment})
	}
	if subject.Screening.CompletionOfScreening.VolunteerIncludedCondition.Color == 3 || subject.Screening.CompletionOfScreening.VolunteerIncludedCondition.Color == 4 {
		errors = append(errors, &model.InformationConsentErrors{
			Field:    "Доброволец включен в исследование?",
			Reasons:  subject.Screening.CompletionOfScreening.VolunteerIncludedCondition.Reason,
			Comments: subject.Screening.CompletionOfScreening.VolunteerIncludedCondition.Comment})
	}
	if subject.Screening.CompletionOfScreening.ReasonIfNotCondition.Color == 3 || subject.Screening.CompletionOfScreening.ReasonIfNotCondition.Color == 4 {
		errors = append(errors, &model.InformationConsentErrors{
			Field:    "Если «НЕТ», необходимо указать причину",
			Reasons:  subject.Screening.CompletionOfScreening.ReasonIfNotCondition.Reason,
			Comments: subject.Screening.CompletionOfScreening.ReasonIfNotCondition.Comment})
	}
	if subject.Screening.CompletionOfScreening.CommentCondition.Color == 3 || subject.Screening.CompletionOfScreening.CommentCondition.Color == 4 {
		errors = append(errors, &model.InformationConsentErrors{
			Field:    "Поле «Комментарий» заполняется по желанию в произвольном тексте.",
			Reasons:  subject.Screening.CompletionOfScreening.CommentCondition.Reason,
			Comments: subject.Screening.CompletionOfScreening.CommentCondition.Comment})
	}

	return errors
}
