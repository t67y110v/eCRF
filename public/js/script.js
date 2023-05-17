function makeAjaxCall() {
    fetch('/api/fruits')
        .then(function (response) {
            return response.text();
        })
        .then(function (text) {
            alert(`we called out to the api via ajax and got this response => ${text}`);
        });
}


function updateColorWithComment(field, id, color) {
    let subjectNum = '{{.Subject.Number}}'
    let protocolId = '{{.Subject.ProtocolId}}'

    fetch('/subject/screening/updatecolorwithcomment', {

        method: 'POST',
        headers: {
            'Accept': 'application/json',
            'Content-Type': 'application/json'
        },
        body: JSON.stringify({
            "field_name": field.name,
            "value": color,
            "subject_number": subjectNum,
            "protocol_id": protocolId,
            "reason": document.getElementById('reason' + id).value,
            "comment": document.getElementById('comment' + id).value,
        })
    })

    $("#condition" + id).load(location.href + " #condition" + id + ">*", "");

}

function updateSecondColor(field, id, color) {
    let subjectNum = '{{.Subject.Number}}'
    let protocolId = '{{.Subject.ProtocolId}}'

    fetch('/subject/screening/updatecolor', {

        method: 'POST',
        headers: {
            'Accept': 'application/json',
            'Content-Type': 'application/json'
        },
        body: JSON.stringify({ "field_name": field.name, "value": color, "subject_number": subjectNum, "protocol_id": protocolId })
    })


    $("#condition" + id).load(window.location.href + " #condition" + id);

}


function gotopage(selval) {
    var value = selval.options[selval.selectedIndex].value;
    window.location.href = value;
}

document.addEventListener('DOMContentLoaded',( function () {
    for (let i = 1; i <= 5; i++) {
        let t = document.getElementById("condition" + i)

        let field = t.getAttribute("field-name")

        let color = t.getAttribute("color-value")
        switch (color) {
            case "1":
                t.insertAdjacentHTML('beforeend', `
              
                <span title="Автоматическая валидация" class="input-group-text"
                  id="basic-addon1">🟢</span>
                  
                 <span title="Автоматическая валидация" class="input-group-text"
                     id="basic-addon1">
                     
                     <div class="dropdown">
                        <button  class="btn btn-outline-secondary dropdown-toggle btn-sm" type="button" data-bs-toggle="dropdown" aria-expanded="false">
                            ⚪
                             </button>
                        <ul class="dropdown-menu">
                            
                           
                                 <li ><div class="d-grid gap-2">

                                <button type="button" onclick="updateSecondColor(${field} ,${i},${2})"  class="btn  btn-outline-success">Принять</button>

                            </div></li>

                                 <li > <div class="d-grid gap-2">
                                    <!-- Button trigger modal -->
                                    <button type="button" class="btn btn-outline-danger" data-bs-toggle="modal" data-bs-target="#denied_${i}">
                                     Отклонить
                                     </button>
                                    

                                        

                            </div></li>

                         
                           
                
                         </ul>
                    </div>
                
                
                </span>
                 <span title="Валидация монитором" class="input-group-text"
                       id="basic-addon1">⚪</span>
               `)

                t.insertAdjacentHTML('beforeend', `
                                      <!-- Modal -->
                                    <div class="modal fade" id="denied_${i}" data-bs-backdrop="static" data-bs-keyboard="false" tabindex="-1"
                                        aria-labelledby="denied_${i}" aria-hidden="true">
                                        <div class="modal-dialog modal-xl">
                                            <div class="modal-content">
                                                <div class="modal-header">
                                                    <h5 class="modal-title" id="denied_${i}">Причина отклонения
                                                    </h5>
                                                    <button type="button" class="btn-close" data-bs-dismiss="modal" aria-label="Close"></button>
                                                </div>
                                                <div class="modal-body">
                                                    <div class="row justify-content-center mt-3">
                                                
                                                        <div class="m-3">
                                                            <div class="m-3">
                                                                <div class="mb-2">
                                                                
                                                                    <input class="form-control" type="hidden" placeholder="{{.Protocol.Id}}"
                                                                        value="{{.Protocol.Id}}" name="protocol_id" id="protocol_id"
                                                                        readonly>
                                                                </div>
                                                                <div class="mb-2">
                                                               
                                                                    <input class="form-control" type="hidden" placeholder="{{.ClinicId}}"
                                                                        value="{{.ClinicId}}" name="center_id" id="center_id" readonly>
                                                                </div>
                                                                <div class="mb-2">
                                                                  
                                                                    <input class="form-control" type="hidden" placeholder={{.Subject.Number}}
                                                                        value={{.Subject.Number}} readonly>
                                                                </div>
                                                            </div>
                                                            <div class="mb-2">
                                                                <label class="form-label">Пожалуйста, укажите причину отказа в верификации:</label>
                                                                <input id="reason${i}" type="text" class="form-control" name="reason">
                                                            </div>
                                                            <div class="mb-2">
                                                                <label class="form-label">Комментарий:</label>
                                                                <input id="comment${i}" type="text" class="form-control" name="comment">
                                                            </div>

                                                            <div class="d-grid gap-2">
                                                                <br>
                                                                <button type="button" onclick="updateColorWithComment(${field} ,${i},${3})"  class="btn  btn-outline-success " data-bs-dismiss="modal">Принять</button>
                                                            </div>
                                                        </div>
                                            
                                                </div>
                                                <div class="modal-footer">
                                                    <button type="button" class="btn btn-danger" data-bs-dismiss="modal">Закрыть</button>

                                                </div>
                                            </div>
                                        </div>
                                    </div>`)


                continue;
            case "2":
                t.insertAdjacentHTML('beforeend', `
                 <span title="Автоматическая валидация" class="input-group-text"
                   id="basic-addon1">🟢</span>
            
              <span title="Автоматическая валидация" class="input-group-text"
                  id="basic-addon1">🟢</span>
                  <span title="Автоматическая валидация" class="input-group-text"
                     id="basic-addon1">
                     
                     <div class="dropdown">
                        <button  class="btn btn-sm btn-outline-secondary dropdown-toggle" type="button" data-bs-toggle="dropdown" aria-expanded="false">
                            ⚪
                             </button>
                        <ul class="dropdown-menu">
                            
                           
                                 <li ><div class="d-grid gap-2">

                                <button type="button" onclick="updateSecondColor(${field} ,${i},${5})"  class="btn btn-outline-success">Принять</button>

                            </div></li>

                                 <li > <div class="d-grid gap-2">
                                    <!-- Button trigger modal -->
                                    <button type="button" class="btn btn-outline-danger" data-bs-toggle="modal" data-bs-target="#denied_${i}">
                                     Отклонить
                                     </button>
                                    

                                        

                            </div>
                            </li>
                
                         </ul>
                    </div>
                
                
                </span>
            `)
            t.insertAdjacentHTML('beforeend', `
                                      <!-- Modal -->
                                    <div class="modal fade" id="denied_${i}" data-bs-backdrop="static" data-bs-keyboard="false" tabindex="-1"
                                        aria-labelledby="denied_${i}" aria-hidden="true">
                                        <div class="modal-dialog modal-xl">
                                            <div class="modal-content">
                                                <div class="modal-header">
                                                    <h5 class="modal-title" id="denied_${i}">Причина отклонения
                                                    </h5>
                                                    <button type="button" class="btn-close" data-bs-dismiss="modal" aria-label="Close"></button>
                                                </div>
                                                <div class="modal-body">
                                                    <div class="row justify-content-center mt-3">
                                                
                                                        <div class="m-3">
                                                            <div class="m-3">
                                                                <div class="mb-2">
                                                                
                                                                    <input class="form-control" type="hidden" placeholder="{{.Protocol.Id}}"
                                                                        value="{{.Protocol.Id}}" name="protocol_id" id="protocol_id"
                                                                        readonly>
                                                                </div>
                                                                <div class="mb-2">
                                                               
                                                                    <input class="form-control" type="hidden" placeholder="{{.ClinicId}}"
                                                                        value="{{.ClinicId}}" name="center_id" id="center_id" readonly>
                                                                </div>
                                                                <div class="mb-2">
                                                                  
                                                                    <input class="form-control" type="hidden" placeholder={{.Subject.Number}}
                                                                        value={{.Subject.Number}} readonly>
                                                                </div>
                                                            </div>
                                                            <div class="mb-2">
                                                                <label class="form-label">Пожалуйста, укажите причину отказа в верификации:</label>
                                                                <input id="reason${i}" type="text" class="form-control" name="reason">
                                                            </div>
                                                            <div class="mb-2">
                                                                <label class="form-label">Комментарий:</label>
                                                                <input id="comment${i}" type="text" class="form-control" name="comment">
                                                            </div>

                                                            <div class="d-grid gap-2">
                                                                <br>
                                                                <button type="button" onclick="updateColorWithComment(${field} ,${i},${4})"  class="btn  btn-outline-success " data-bs-dismiss="modal">Принять</button>
                                                            </div>
                                                        </div>
                                            
                                                </div>
                                                <div class="modal-footer">
                                                    <button type="button" class="btn btn-danger" data-bs-dismiss="modal">Закрыть</button>

                                                </div>
                                            </div>
                                        </div>
                                    </div>`)
                continue;
            case "3":
                t.insertAdjacentHTML('beforeend', `
              <span title="Автоматическая валидация" class="input-group-text"
                 id="basic-addon1">🟢</span>
            
                 <span title="Автоматическая валидация" class="input-group-text"
                     id="basic-addon1">
                     
                     <div class="dropdown">
                        <button  class="btn btn-outline-secondary dropdown-toggle btn-sm" type="button" data-bs-toggle="dropdown" aria-expanded="false">
                            🔴
                             </button>
                        <ul class="dropdown-menu">
                            
                           
                                 <li ><div class="d-grid gap-2">

                                <button type="button" onclick="updateSecondColor(${field} ,${i},${2})"  class="btn  btn-outline-success">Принять</button>

                            </div></li>

                                 <li > <div class="d-grid gap-2">
                                    <!-- Button trigger modal -->
                                    <button type="button" class="btn btn-outline-danger" data-bs-toggle="modal" data-bs-target="#denied_${i}">
                                     Отклонить
                                     </button>
                                    

                                        

                            </div></li>

                         
                           
                
                         </ul>
                    </div>
                
                
                </span>
              <span title="Валидация монитором" class="input-group-text"
                      id="basic-addon1">⚪</span>
            `)

                t.insertAdjacentHTML('beforeend', `
                                      <!-- Modal -->
                                    <div class="modal fade" id="denied_${i}" data-bs-backdrop="static" data-bs-keyboard="false" tabindex="-1"
                                        aria-labelledby="denied_${i}" aria-hidden="true">
                                        <div class="modal-dialog modal-xl">
                                            <div class="modal-content">
                                                <div class="modal-header">
                                                    <h5 class="modal-title" id="denied_${i}">Причина отклонения
                                                    </h5>
                                                    <button type="button" class="btn-close" data-bs-dismiss="modal" aria-label="Close"></button>
                                                </div>
                                                <div class="modal-body">
                                                    <div class="row justify-content-center mt-3">
                                                
                                                        <div class="m-3">
                                                            <div class="m-3">
                                                                <div class="mb-2">
                                                                
                                                                    <input class="form-control" type="hidden" placeholder="{{.Protocol.Id}}"
                                                                        value="{{.Protocol.Id}}" name="protocol_id" id="protocol_id"
                                                                        readonly>
                                                                </div>
                                                                <div class="mb-2">
                                                               
                                                                    <input class="form-control" type="hidden" placeholder="{{.ClinicId}}"
                                                                        value="{{.ClinicId}}" name="center_id" id="center_id" readonly>
                                                                </div>
                                                                <div class="mb-2">
                                                                  
                                                                    <input class="form-control" type="hidden" placeholder={{.Subject.Number}}
                                                                        value={{.Subject.Number}} readonly>
                                                                </div>
                                                            </div>
                                                            <div class="mb-2">
                                                                <label class="form-label">Пожалуйста, укажите причину отказа в верификации:</label>
                                                                <input id="reason${i}" type="text" class="form-control" name="reason">
                                                            </div>
                                                            <div class="mb-2">
                                                                <label class="form-label">Комментарий:</label>
                                                                <input id="comment${i}" type="text" class="form-control" name="comment">
                                                            </div>

                                                            <div class="d-grid gap-2">
                                                                <br>
                                                                <button type="button" onclick="updateColorWithComment(${field} ,${i},${3})"  class="btn  btn-outline-success " data-bs-dismiss="modal">Принять</button>
                                                            </div>
                                                        </div>
                                            
                                                </div>
                                                <div class="modal-footer">
                                                    <button type="button" class="btn btn-danger" data-bs-dismiss="modal">Закрыть</button>

                                                </div>
                                            </div>
                                        </div>
                                    </div>`)

                continue;
            case "4":
                t.insertAdjacentHTML('beforeend', `
              <span title="Автоматическая валидация" class="input-group-text"
                 id="basic-addon1">🟢</span>
            
              <span title="Автоматическая валидация" class="input-group-text"
                  id="basic-addon1">🟢</span>
              <span title="Валидация монитором" class="input-group-text"
                      id="basic-addon1">
                      
                      
                      <div class="dropdown">
                        <button  class="btn btn-sm btn-outline-secondary dropdown-toggle" type="button" data-bs-toggle="dropdown" aria-expanded="false">
                            🔴
                             </button>
                        <ul class="dropdown-menu">
                            
                           
                                 <li ><div class="d-grid gap-2">

                                <button type="button" onclick="updateSecondColor(${field} ,${i},${5})"  class="btn btn-outline-success">Принять</button>

                            </div></li>

                                 <li > <div class="d-grid gap-2">
                                    <!-- Button trigger modal -->
                                    <button type="button" class="btn btn-outline-danger" data-bs-toggle="modal" data-bs-target="#denied_${i}">
                                     Отклонить
                                     </button>
                                    

                                        

                            </div>
                            </li>
                
                         </ul>
                    </div>
                      
                      
                      </span>
            `)
            t.insertAdjacentHTML('beforeend', `
                                      <!-- Modal -->
                                    <div class="modal fade" id="denied_${i}" data-bs-backdrop="static" data-bs-keyboard="false" tabindex="-1"
                                        aria-labelledby="denied_${i}" aria-hidden="true">
                                        <div class="modal-dialog modal-xl">
                                            <div class="modal-content">
                                                <div class="modal-header">
                                                    <h5 class="modal-title" id="denied_${i}">Причина отклонения
                                                    </h5>
                                                    <button type="button" class="btn-close" data-bs-dismiss="modal" aria-label="Close"></button>
                                                </div>
                                                <div class="modal-body">
                                                    <div class="row justify-content-center mt-3">
                                                
                                                        <div class="m-3">
                                                            <div class="m-3">
                                                                <div class="mb-2">
                                                                
                                                                    <input class="form-control" type="hidden" placeholder="{{.Protocol.Id}}"
                                                                        value="{{.Protocol.Id}}" name="protocol_id" id="protocol_id"
                                                                        readonly>
                                                                </div>
                                                                <div class="mb-2">
                                                               
                                                                    <input class="form-control" type="hidden" placeholder="{{.ClinicId}}"
                                                                        value="{{.ClinicId}}" name="center_id" id="center_id" readonly>
                                                                </div>
                                                                <div class="mb-2">
                                                                  
                                                                    <input class="form-control" type="hidden" placeholder={{.Subject.Number}}
                                                                        value={{.Subject.Number}} readonly>
                                                                </div>
                                                            </div>
                                                            <div class="mb-2">
                                                                <label class="form-label">Пожалуйста, укажите причину отказа в верификации:</label>
                                                                <input id="reason${i}" type="text" class="form-control" name="reason">
                                                            </div>
                                                            <div class="mb-2">
                                                                <label class="form-label">Комментарий:</label>
                                                                <input id="comment${i}" type="text" class="form-control" name="comment">
                                                            </div>

                                                            <div class="d-grid gap-2">
                                                                <br>
                                                                <button type="button" onclick="updateColorWithComment(${field} ,${i},${4})"  class="btn  btn-outline-success " data-bs-dismiss="modal">Принять</button>
                                                            </div>
                                                        </div>
                                            
                                                </div>
                                                <div class="modal-footer">
                                                    <button type="button" class="btn btn-danger" data-bs-dismiss="modal">Закрыть</button>

                                                </div>
                                            </div>
                                        </div>
                                    </div>`)

                continue;
            case "5":
                t.insertAdjacentHTML('beforeend', `
              <span title="Автоматическая валидация" class="input-group-text"
                 id="basic-addon1">🟢</span>
            
              <span title="Автоматическая валидация" class="input-group-text"
                  id="basic-addon1">🟢</span>
              <span title="Валидация монитором" class="input-group-text"
                      id="basic-addon1">🟢</span>
            `)
                continue;




            default:
                t.insertAdjacentHTML('beforeend', `
               <span title="Автоматическая валидация" class="input-group-text"
                      id="basic-addon1">⚪</span>
            
              <span title="Автоматическая валидация" class="input-group-text"
                      id="basic-addon1">⚪</span>
                 <span title="Валидация монитором" class="input-group-text"
                  id="basic-addon1">⚪</span>
            `)
                continue;
        }

    }

}), false);
;




