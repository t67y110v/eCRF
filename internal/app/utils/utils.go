package utils

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
	m["signed"] = "screening.informaionconsent.signedcondition.color"
	m["date_of_sign"] = "screening.informaionconsent.dateofsigncondition.color"
	m["time_of_sign"] = "screening.informaionconsent.timeofsigncondition.color"
	m["original"] = "screening.informaionconsent.receivedaninsurancepolicycondition.color"
	m["consent"] = "screening.informaionconsent.receivedaninformaionconsentcondition.color"
	return m[field]
}
