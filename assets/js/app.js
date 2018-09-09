
$(document).ready(function () {
    var contactsTable = $("#contactsTable").DataTable({
        "paging": true,
        "ajax": {
            "url" : "contact/",
            "type": "GET",
            "dataType": "json",
            "contentType": 'application/json; charset=utf-8',
            "dataSrc": ""
        }
        ,
        "columns": [
            {data: "Id"},
            {data: "FirstName"},
            {data: "LastName"},
            {data: "Phone"},
            {data: "Email"}
        ]
        ,
        "columnDefs": [
            {
                "targets" : [0],
                "visible" : false,
                "searchable": false
            },
            {
                "targets" : [5],
                "data": null,
                "defaultContent": "<button id='editBtn'>Edit</button>"
            },
            {
                "targets" : [6],
                "data": null,
                "defaultContent": "<button id='remBtn'>Remove</button>"
            }
        ]
    });

    $("#contactsTable tbody").on('click', '#editBtn', function () {
        var id2Edit = contactsTable.row( $(this).parents('tr') ).data()["Id"];
    });

    $("#contactsTable tbody").on('click', '#remBtn', function () {
        var id2Remove = contactsTable.row( $(this).parents('tr') ).data()["Id"];

        $.ajax({
            url: "contact/" + id2Remove,
            type: "DELETE",
            success: function (data) {
                contactsTable.ajax.reload();

            },
            error: function (data) {
                alert("Error while removing the record. Please check the logs.");
            }

        });
    });
});