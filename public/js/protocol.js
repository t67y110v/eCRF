

function updateColorWithComment(id, color) {
    let protocolId = document.getElementById("protocol_id_for_js").getAttribute("value")
    let subjectNum = document.getElementById("subject_id_for_js").getAttribute("value")
    let t = document.getElementById("condition" + id)
    let field = t.getAttribute("field-name")
    console.log(document.getElementById('reason' + id).value)
    fetch('/subject/screening/updatecolorwithcomment', {

        method: 'POST',
        headers: {
            'Accept': 'application/json',
            'Content-Type': 'application/json'
        },
        body: JSON.stringify({
            "field_name": field,
            "value": color,
            "subject_number": subjectNum,
            "protocol_id": protocolId,
            "reason": document.getElementById('reason' + id).value,
            "comment": document.getElementById('comment' + id).value,
        })
    })

    $("#condition" + id).load(location.href + " #condition" + id + ">*", "");

}

function updateField(id, color) {
    let protocolId = document.getElementById("protocol_id_for_js").getAttribute("value")
    let subjectNum  = document.getElementById("subject_id_for_js").getAttribute("value")
    let t = document.getElementById("condition" + id)
    let field = t.getAttribute("field-name")

    console.log(subjectNum, protocolId, field ,document.getElementById('new_value_' + id).value )
    fetch('/subject/screening/updatefield', {
        method: 'POST',
        headers: {

            'Accept': 'application/json',
            'Content-Type': 'application/json'
        },
        body: JSON.stringify({
            "field_name": field,
            "value": document.getElementById('new_value_' + id).value,
            "color": color,
            "subject_number": subjectNum, 
            "protocol_id": protocolId,
        })


    })



}





function updateSecondColor(id, color) {
    let protocolId = document.getElementById("protocol_id_for_js").getAttribute("value")
    let subjectNum = document.getElementById("subject_id_for_js").getAttribute("value")

    let t = document.getElementById("condition" + id)

    let field = t.getAttribute("field-name")

    console.log(field, color, subjectNum, protocolId)
    fetch('/subject/screening/updatecolor', {

        method: 'POST',
        headers: {
            'Accept': 'application/json',
            'Content-Type': 'application/json'
        },
        body: JSON.stringify({
             "field_name": field, "value": color, "subject_number": subjectNum, "protocol_id": protocolId })
    })


    $("#condition" + id).load(window.location.href + " #condition" + id);

}




(function () {
    let protocolId = document.getElementById("protocol_id_for_js").getAttribute("value")
    let subjectNum = document.getElementById("subject_id_for_js").getAttribute("value")

    for (let i = 1; i <= 57; i++) {
        let t = document.getElementById("condition" + i)
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

                                <button type="button" onclick="updateSecondColor(${i},${2})"  class="btn  btn-outline-success">Принять</button>

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
                                                                
                                                                    <input class="form-control" type="hidden" placeholder="${protocolId}"
                                                                        value="${protocolId}" name="protocol_id" id="protocol_id"
                                                                        readonly>
                                                                </div>
                                                                <div class="mb-2">
                                                               
                                                                    <input class="form-control" type="hidden" placeholder="${subjectNum}"
                                                                        value="${subjectNum}" name="center_id" id="center_id" readonly>
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
                                                                <button type="button" onclick="updateColorWithComment(${i},${3})"  class="btn  btn-outline-success " data-bs-dismiss="modal">Принять</button>
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

                                <button type="button" onclick="updateSecondColor(${i},${5})"  class="btn btn-outline-success">Принять</button>

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
                                                                
                                                                    <input class="form-control" type="hidden" placeholder="${protocolId}"
                                                                        value="${protocolId}" name="protocol_id" id="protocol_id"
                                                                        readonly>
                                                                </div>
                                                                <div class="mb-2">
                                                               
                                                                    <input class="form-control" type="hidden" placeholder="${subjectNum}"
                                                                        value="${subjectNum}" name="center_id" id="center_id" readonly>
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
                                                               
                                                            </div>
                                                        </div>
                                            
                                                </div>
                                                <div class="modal-footer">
                                                    <button type="button" onclick="updateColorWithComment(${i},${4})"  class="btn  btn-outline-success " data-bs-dismiss="modal">Принять</button>
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
                            
                           

                                 <li > 
                                 <div class="d-grid gap-2">
                                    <!-- Button trigger modal -->
                                    <button type="button" class="btn btn-outline-warning" data-bs-toggle="modal" data-bs-target="#new_value${i}">
                                     Изменить
                                     </button>
                                    
                            </div>
                            
                            </li>

                         
                           
                
                         </ul>
                    </div>
                
                
                </span>
              <span title="Валидация монитором" class="input-group-text"
                      id="basic-addon1">⚪</span>
            `)

                t.insertAdjacentHTML('beforeend', `
                                      <!-- Modal -->
                                    <div class="modal fade" id="new_value${i}" data-bs-backdrop="static" data-bs-keyboard="false" tabindex="-1"
                                        aria-labelledby="new_value${i}" aria-hidden="true">
                                        <div class="modal-dialog modal-xl">
                                            <div class="modal-content">
                                                <div class="modal-header">
                                                    <h5 class="modal-title" id="new_value${i}">Изменение значения
                                                    </h5>
                                                    <button type="button" class="btn-close" data-bs-dismiss="modal" aria-label="Close"></button>
                                                </div>
                                                <div class="modal-body">
                                                    <div class="row justify-content-center mt-3">
                                                
                                                        <div class="m-3">
                                                            
                                                            <div class="mb-2">
                                                                <label class="form-label">Новое значение:</label>
                                                                <input id="new_value_${i}" type="text" class="form-control" name="new_value">
                                                            </div>

                                                            <div class="d-grid gap-2">
                                                                <br>
                                                                <button type="button" onclick="updateField(${i},${1})"  class="btn  btn-outline-success " data-bs-dismiss="modal">Сохранить</button>
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

                                <button type="button" onclick="updateSecondColor(${i},${5})"  class="btn btn-outline-success">Принять</button>

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
                                                                
                                                                    <input class="form-control" type="hidden" placeholder="${protocolId}"
                                                                        value="${protocolId}" name="protocol_id" id="protocol_id"
                                                                        readonly>
                                                                </div>
                                                                <div class="mb-2">
                                                               
                                                                    <input class="form-control" type="hidden" placeholder="${subjectNum}"
                                                                        value="${subjectNum}" name="center_id" id="center_id" readonly>
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
                                                                <button type="button" onclick="updateColorWithComment(${i},${4})"  class="btn  btn-outline-success " data-bs-dismiss="modal">Принять</button>
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

})();




