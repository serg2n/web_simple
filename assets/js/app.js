
$(document).ready(function () {
    $("#contactsTable").DataTable({
        "paging": true,
        "ajax": {
            "url" : "contact/",
            "type": "GET",
            "dataType": "json",
            "contentType": 'application/json; charset=utf-8',
            "dataSrc": ""
        }
        ,
        columns: [
            {data: "Id"},
            {data: "FirstName"},
            {data: "LastName"},
            {data: "Phone"},
            {data: "Email"},
            {
                "data": null,
                "defaultContent": "<button>Edit</button>"
            },
            {
                "data": null,
                "defaultContent": "<button>Remove</button>"
            }
        ]
    });
});