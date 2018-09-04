package http

import (
	"encoding/json"
	"log"
	"net/http"
	"simple-web-app"
	"simple-web-app/constants"
	"strconv"
)

type ContactController struct {
	ContactService simplewebapp.ContactService
}

func (cc *ContactController) Contact(res http.ResponseWriter, req *http.Request) {
	id, err := IdFromRequest(req, 2)
	if err != nil {
		log.Printf("Cannot get ID from the url: %s, %v", req.URL.Path, err)
		BadRequestResponse(res)
		return
	}

	contact, err := cc.ContactService.Contact(id)
	if err != nil {
		InternalServerErrorResponse(res)
		return
	}
	if contact == nil {
		log.Printf("Contact not found: %d", id)
		http.NotFound(res, req)
		return
	}

	resultContact, err := json.Marshal(contact)
	if err != nil {
		log.Printf("Cannot get contact by id %d, error: %v", id, err)
		return
	}
	res.Header().Set("content-type", "application/json")
	_, err = res.Write(resultContact)
	if err != nil {
		log.Printf("Cannot write data to HTTP response: %v", err)
		InternalServerErrorResponse(res)
	}

}

func (cc *ContactController) CreateContact(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("content-type", "application/json")
	res.Write([]byte("{ \"result\":\"OK\"}"))
	//ContactSer
}

func (cc *ContactController) Contacts(res http.ResponseWriter, req *http.Request) {
	query := req.URL.Query()
	pageSizeStr := query.Get("pageSize")
	var pageSize, offset int
	if pageSizeStr == "" {
		pageSize = constants.PageSize
	} else {
		pageSize = convertHttpParam2Int(pageSizeStr, constants.PageSize)
	}
	offsetStr := query.Get("offset")
	if offsetStr == "" {
		offset = constants.Offset
	} else {
		offset = convertHttpParam2Int(offsetStr, constants.Offset)
	}

	resultContacts, err := cc.ContactService.Contacts(pageSize, offset)
	if err != nil {
		InternalServerErrorResponse(res)
		return
	}

	resultData, err := json.Marshal(resultContacts)
	if err != nil {
		log.Printf("Cannot marshal data (list of contacts): %v", err)
		return
	}
	res.Header().Set("content-type", "application/json")

	_, err = res.Write(resultData)
	if err != nil {
		log.Printf("Cannot write data to HTTP response: %v", err)
		InternalServerErrorResponse(res)
	}

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

func convertHttpParam2Int(srcStr string, defVal int) (int) {
	res, err := strconv.Atoi(srcStr)
	if err != nil {
		log.Printf("Cannot convert http param %s to int: %v, using default value %d", srcStr, err, defVal)
		return defVal
	}
	return res
}
