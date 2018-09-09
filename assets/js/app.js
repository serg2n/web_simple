var contact2UpdateId;

$(document).ready(function () {
    var contactsTable = $("#contactsTable").DataTable({
        "paging": true,
        "ajax": {
            "url": "contact/",
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
                "targets": [0],
                "visible": false,
                "searchable": false
            },
            {
                "targets": [5],
                "data": null,
                "defaultContent": "<button id='editBtn'>Edit</button>"
            },
            {
                "targets": [6],
                "data": null,
                "defaultContent": "<button id='remBtn'>Remove</button>"
            }
        ]
    });

    $("#contactsTable tbody").on('click', '#editBtn', function () {
        var id2Edit = contactsTable.row($(this).parents('tr')).data()["Id"];
        contact2UpdateId = id2Edit;
        $.ajax({
            url: "contact/" + id2Edit,
            type: "GET",
            success: function (data) {
                clearContactInputs();
                $("#createContactBtn").hide();
                $("#updateContactBtn").show();
                $("#firstName").val(data["FirstName"]);
                $("#lastName").val(data["LastName"]);
                $("#phone").val(data["Phone"]);
                $("#email").val(data["Email"]);
                $("#addNewContactModal").modal('show');


            },
            error: function (data) {
                alert("Error while editing the record. Please check the logs.");
            }

        });
    });

    $("#contactsTable tbody").on('click', '#remBtn', function () {
        var id2Remove = contactsTable.row($(this).parents('tr')).data()["Id"];

        bootbox.confirm("Delete the record?", function (result) {
            if (result === true) {
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
            }
        });
    });
});

function showAddNewContactModal() {
    $("#createContactBtn").show();
    $("#updateContactBtn").hide();
    clearContactInputs();
    $("#addNewContactModal").modal('show');
}

function saveContact() {
    var newContact = {
        FirstName: $("#firstName").val().trim(),
        LastName: $("#lastName").val().trim(),
        Phone: $("#phone").val().trim(),
        Email: $("#email").val().trim()
    };

    $.ajax({
        url: "contact/",
        type: "POST",
        dataType: "json",
        contentType: "application/json",
        data : JSON.stringify(newContact),
        success: function(data) {
            $("#addNewContactModal").modal('hide');
            $("#contactsTable").DataTable().ajax.reload();
            clearContactInputs();
        },
        error: function(data) {
            bootbox.alert("Error when saving data. Please check the logs.");
        }
    });
}

function updateContact() {
    var contact2Update = {
        FirstName: $("#firstName").val().trim(),
        LastName: $("#lastName").val().trim(),
        Phone: $("#phone").val().trim(),
        Email: $("#email").val().trim()
    };

    $.ajax({
        url: "contact/" + contact2UpdateId,
        type: "PUT",
        dataType: "json",
        contentType: "application/json",
        data : JSON.stringify(contact2Update),
        success: function(data) {
            $("#addNewContactModal").modal('hide');
            $("#contactsTable").DataTable().ajax.reload();
            clearContactInputs();
        },
        error: function(data) {
            bootbox.alert("Error when saving data. Please check the logs.");
        }
    });
}

function clearContactInputs() {
    $("#firstName").val('');
    $("#lastName").val('');
    $("#phone").val('');
    $("#email").val('');
}