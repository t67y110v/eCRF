package utils

import model "github.com/t67y110v/web/internal/app/model/subject"

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
