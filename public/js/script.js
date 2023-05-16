function makeAjaxCall() {
    fetch('/api/fruits')
        .then(function (response) {
            return response.text();
        })
        .then(function (text) {
            alert(`we called out to the api via ajax and got this response => ${text}`);
        });
}



function getCondition() {
    let target = document.querySelector('#condition')
    target.insertAdjacentHTML('beforeend', `
{{if eq  .Subject.Screening.InformaionConsent.SignedCondition.Color 0}}
<span title="Автоматическая валидация" class="input-group-text"
    id="basic-addon1">⚪</span>
    {{else}}
<span title="Автоматическая валидация" class="input-group-text"
    id="basic-addon1">⚪</span>
<span title="Валидация монитором" class="input-group-text"
    id="basic-addon1">⚪</span>
<span title="Валидация дата-менеджером" class="input-group-text"
    id="basic-addon1">⚪</span>
    {{end}}

`)
}
window.onload = getCondition;