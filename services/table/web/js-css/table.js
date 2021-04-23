/* Get JSON data from server */
const url = document.location.protocol + "//" + document.location.hostname + ":"+ document.location.port +"/users";
function ajaxRequest(params) {

    let obj
    $.getJSON(url, function (data, textStatus, jqXHR) {
        obj = data
        if (jqXHR.status != 200) {
            console.log("getJSON status", jqXHR.status)
        };
    }).done(function () {
        console.log("getJSON success");
        params.success({
            "rows": obj,
            "total": obj.length
        }, null, {});
    }).fail(function (er) {
        console.log("getJSON error");
        params.error(er);
    })
}

let $table = $('#table');
let dataRow
$table.bootstrapTable({
    onClickRow: function (row) {
        dataRow = row
    }
})

// Уведомления
function alerts(classAlert, text) {
    return `<div class="alert `+classAlert+`" role="alert" data-dismiss="alert">
            `+text+`
            <button type="button" class="close" data-dismiss="alert" aria-label="Close">
            <span aria-hidden="true">&times;</span>
            </button></div>`}


/* // Детальная информация */
function detailFormatter(index, row) {
    var html = []
    $.each(row, function (key, value) {
        html.push('<p class="detail_info"><b>' + key + ':</b> ' + value + '</p>')
    });
    return html.join('')
};

/* // Меню со списком в каждой строке */
function dropdown_row() {
    return `<div class="dropdown">
            <button class="btn btn-primary" type="button" id="dropdownMenuButton"
            data-toggle="dropdown" aria-haspopup="true" aria-expanded="false"><i class="fas fa-angle-down"></i></button>
            <div class="dropdown-menu" aria-labelledby="dropdownMenuButton">
            <button id="button-update-row" type="button" class="dropdown-item" data-toggle="modal" data-target="#row-table">Изменить</button>
            <button type="button" class="dropdown-item" data-toggle="modal" data-target="#delete-confirm">Удалить</button>
            </div>
            </div>`
};

let forms = $("#form-row");
// Изменение кнопок и заполнение полей формы для update
$table.on('click', "#button-update-row", function () {
    $("#button-add-row-submit").hide();
    $("#button-update-row-submit").show();
    $(".modal-body").find('#username').val(dataRow.username)
    $(".modal-body").find('#pc_name').val(dataRow.pc_name)
    $(".modal-body").find('#user_group').val(dataRow.user_group)
    $(".modal-body").find('#phone_number').val(dataRow.phone_number)
    $(".modal-body").find('#cabinet').val(dataRow.cabinet)
    $(".modal-body").find('#discription').val(dataRow.discription)
});
// Отправка данных с формы при клике
$('body').on("click", "#button-update-row-submit", function (event) {
    // Проверка полей
    var validation = Array.prototype.filter.call(forms, function (form) {
        if (form.checkValidity() === false) {
            event.preventDefault();
            event.stopPropagation();
            $(".modal #for-alerts").append(alerts("alert-warning", "Незаполнены поля"))
        } else {
            var data = forms.serializeArray();

            function getFormData(data) {
                var unindexed_array = data;
                var indexed_array = {};

                $.map(unindexed_array, function (n, i) {
                    indexed_array[n['name']] = n['value'];
                });
                indexed_array.user_id = dataRow.user_id;
                return indexed_array;
                
            }
            $.ajax({
                url: url,
                type: 'PUT',
                data: JSON.stringify(getFormData(data)),
                contentType: 'application/json; charset=utf-8',
                dataType: 'json',
                async: false,
                success: function (data) {
                    $table.bootstrapTable('refresh')
                    $("#for-alerts").append(alerts("alert-success", "Запись изменена: "+dataRow.username))
                    $("#row-table").modal('hide');
                    
                },
                error: function (data) {
                    $(".modal #for-alerts").append(alerts("alert-danger", "Произошла ошибка"))
                }
            });
        }
    });
})


/* // ADD ROW */
$('body').on("click", "#button-add-row", function () {
    $("#button-add-row-submit").show();
    $("#button-update-row-submit").hide();
});
// Очистка полей формы когда hidden
$("#row-table").on('hidden.bs.modal', function () {
    $(this).find('form')[0].reset();
});
$('body').on("click", "#button-add-row-submit", function (event) {
    // Проверка полей
    var validation = Array.prototype.filter.call(forms, function (form) {
        if (form.checkValidity() === false) {
            event.preventDefault();
            event.stopPropagation();
            $(".modal #for-alerts").append(alerts("alert-warning", "Незаполнены поля"))
        } else {
            var data = forms.serializeArray();

            function getFormData(data) {
                var unindexed_array = data;
                var indexed_array = {};

                $.map(unindexed_array, function (n, i) {
                    indexed_array[n['name']] = n['value'];
                });

                return indexed_array;
            }
            $.ajax({
                url: url,
                type: 'POST',
                data: JSON.stringify(getFormData(data)),
                contentType: 'application/json; charset=utf-8',
                dataType: 'json',
                async: false,
                success: function (data) {
                    $("#for-alerts").append(alerts("alert-success","Запись создана: "+ data.username))
                    $("#row-table").modal('hide');
                    $table.bootstrapTable('refresh')
                },
                error: function (data) {
                    $(".modal #for-alerts").append(alerts("alert-danger", "Произошла ошибка"))
                }
            });
        }
    })

})

// Delete row
$('body').on("click", "#button-delete-row", function (event) {
    var data = forms.serializeArray();

    function getFormData(data) {
        var unindexed_array = data;
        var indexed_array = {};

        $.map(unindexed_array, function (n, i) {
            indexed_array[n['name']] = n['value'];
        });
        indexed_array.user_id = dataRow.user_id;
        return indexed_array;
    }
    $.ajax({
        url: url,
        type: 'DELETE',
        data: JSON.stringify(getFormData(data)),
        contentType: 'application/json; charset=utf-8',
        dataType: 'json',
        async: false,
        success: function (data) {
            $("#for-alerts").append(alerts("alert-success", "Запись удалена: "+dataRow.username))
            $("#delete-confirm").modal('hide');
            $table.bootstrapTable('refresh')
        },
        error: function (data) {
            $(".modal #for-alerts").append(alerts("alert-danger", "Произошла ошибка"))
        }
    });
})

$('#table').on('load-success.bs.table', function () {
    $("#countRows").text("Всего строк: "+ $('#table').bootstrapTable('getOptions').totalRows)
    
})