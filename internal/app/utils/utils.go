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
