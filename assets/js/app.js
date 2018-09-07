
$(document).ready(function () {
    $("#contactsTable").DataTable({
        // "processing" :true,
        // "serverSide": true,
        "ajax": {
            "url" : "contact/",
            "dataType": "json",
            "dataSrc": ""
        },
        columns: [
            {data: "Id"},
            {data: "FirstName"},
            {data: "LastName"},
            {data: "Phone"},
            {data: "Email"}
        ]
    });
});