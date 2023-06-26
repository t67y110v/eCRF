package utils

import (
	"strings"
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

func GetFieldNameForUpdate(field string) string {
	s := strings.Split(field, ".")
	last := s[len(s)-2]
	res, _ := strings.CutSuffix(last, "condition")
	return res
}

func GetFieldName(field string) string {
	m := make(map[string]string)

	m["dateofstart"] = "screening.startofscreening.datestartofscreeningcondition."
	m["timeofstart"] = "screening.startofscreening.timestartofscreeningcondition."

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

	m["presenceofaninformationpanel"] = "screening.inclusioncriteria.presenceofaninformationpanelcondition."
	m["aged18to55years"] = "screening.inclusioncriteria.aged18to55yearscondition."
	m["negativehivtestresult"] = "screening.inclusioncriteria.negativehivtestresultcondition."
	m["bodymassindex"] = "screening.inclusioncriteria.bodymassindexcondition."
	m["absenceofacuteinfectiousdiseases"] = "screening.inclusioncriteria.absenceofacuteinfectiousdiseasescondition."
	m["consenttouseeffectivemethodsofcontraception"] = "screening.inclusioncriteria.consenttouseeffectivemethodsofcontraceptioncondition."
	m["negativepregnancytest"] = "screening.inclusioncriteria.negativepregnancytestcondition."
	m["negativedrugtest"] = "screening.inclusioncriteria.negativedrugtestcondition."
	m["negativealcoholtest"] = "screening.inclusioncriteria.negativealcoholtestcondition."
	m["nohistoryofseverepostvaccinationreactions"] = "screening.inclusioncriteria.nohistoryofseverepostvaccinationreactionscondition."
	m["indicatorsbloodtestsatscreeningwithin"] = "screening.inclusioncriteria.indicatorsbloodtestsatscreeningwithincondition."
	m["nomyocardialchanges"] = "screening.inclusioncriteria.nomyocardialchangescondition."
	m["negativetestresultforcovid"] = "screening.inclusioncriteria.negativetestresultforcovidcondition."
	m["nocontraindicationstovaccination"] = "screening.inclusioncriteria.nocontraindicationstovaccinationcondition."

	m["lackofsignedinformedconsent"] = "screening.exclusionсriteria.lackofsignedinformedconsentcondition."
	m["steroidtherapy"] = "screening.exclusionсriteria.steroidtherapycondition."
	m["therapywithimmunosuppressivedrugs"] = "screening.exclusionсriteria.therapywithimmunosuppressivedrugscondition."
	m["femalesubjectsduringpregnancy"] = "screening.exclusionсriteria.femalesubjectsduringpregnancycondition."
	m["strokeinlessthanoneyear"] = "screening.exclusionсriteria.strokeinlessthanoneyearcondition."
	m["chronicsystemicinfections"] = "screening.exclusionсriteria.chronicsystemicinfectionscondition."
	m["aggravatedallergichistory"] = "screening.exclusionсriteria.aggravatedallergichistorycondition."
	m["presenceofahistoryofneoplasms"] = "screening.exclusionсriteria.presenceofahistoryofneoplasmscondition."
	m["historyofsplenectomy"] = "screening.exclusionсriteria.historyofsplenectomycondition."
	m["neutropenia"] = "screening.exclusionсriteria.neutropeniacondition."
	m["subjectswithactivesyphilis"] = "screening.exclusionсriteria.subjectswithactivesyphiliscondition."
	m["anorexia"] = "screening.exclusionсriteria.anorexiacondition."
	m["extensivetattoos"] = "screening.exclusionсriteria.extensivetattooscondition."
	m["takingnarcoticandpsychostimulantdrugs"] = "screening.exclusionсriteria.takingnarcoticandpsychostimulantdrugscondition."
	m["smokingmoretthantencigarettesaday"] = "screening.exclusionсriteria.smokingmoretthantencigarettesadaycondition."
	m["alcoholintake"] = "screening.exclusionсriteria.alcoholintakecondition."
	m["plannedhospitalization"] = "screening.exclu	sionсriteria.plannedhospitalizationconditiont."
	m["donorblooddonation"] = "screening.exclusionсriteria.donorblooddonationcondition."
	m["subjectparticipationinanyotherstudy"] = "screening.exclusionсriteria.subjectparticipationinanyotherstudycondition."
	m["anyvaccinationinthelastmonth"] = "screening.exclusionсriteria.anyvaccinationinthelastmonthcondition."
	m["inabilitytoreadinrussian"] = "screening.exclusionсriteria.inabilitytoreadinrussiancondition."
	m["researchcenterstaff"] = "screening.exclusionсriteria.researchcenterstaffcondition."
	m["anyotherstateofthesubjectofthestudy"] = "screening.exclusionсriteria.anyotherstateofthesubjectofthestudycondition."

	m["volunteereligible"] = "screening.completionofscreening.volunteereligiblecondition."
	m["noexclusioncriteria"] = "screening.completionofscreening.noexclusioncriteriacondition."
	m["informedoftherestrictions"] = "screening.completionofscreening.informedoftherestrictionscondition."
	m["volunteerincluded"] = "screening.completionofscreening.volunteerincludedcondition."
	m["reasonifnot"] = "screening.completionofscreening.reasonifnotcondition."
	m["commentvalue"] = "screening.completionofscreening.commentcondition."
	return m[field]
}
