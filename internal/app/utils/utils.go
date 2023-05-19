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
