package http

import (
	"net/http"
	"simple-web-app"
	"simple-web-app/constants"
)

type ContactController struct {
	ContactService simplewebapp.ContactService
}

func (cc *ContactController) Contact(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("content-type", "application/json")
	res.Write([]byte("{ \"result\":\"OK\"}"))
}

func (cc *ContactController) CreateContact(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("content-type", "application/json")
	res.Write([]byte("{ \"result\":\"OK\"}"))
	//ContactSer
}

func (cc *ContactController) Contacts(res http.ResponseWriter, req *http.Request) {
	query := req.URL.Query()
	pageSize := query.Get("pageSize")
	if pageSize == "" {
		pageSize = constants.PageSize
	}
	offset := query.Get("offset")
	if offset == "" {
		offset = constants.Offset
	}
	cc.ContactService.Contacts(0, 0)
}

func (cc *ContactController) UpdateContact(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("content-type", "application/json")
	res.Write([]byte("{ \"result\":\"OK\"}"))
}

func (cc *ContactController) DeleteContact(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("content-type", "application/json")
	res.Write([]byte("{ \"result\":\"OK\"}"))
}

func NewContactController(cs simplewebapp.ContactService) (*ContactController) {
	return &ContactController{
		ContactService: cs,
	}
}
