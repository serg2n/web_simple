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
	log.Printf("Create a new Contact...")
	newContact := new(simplewebapp.Contact)

	if err := json.NewDecoder(req.Body).Decode(newContact); err != nil {
		log.Printf("Error creating a new Contact: %v", err)
	}

	newContact, err := cc.ContactService.CreateContact(newContact)

	if err != nil {
		InternalServerErrorResponse(res)
		return
	}

	data, err := json.Marshal(newContact)
	if err != nil {
		log.Printf("Cannot marshal data: %v", err)
		InternalServerErrorResponse(res)
		return
	}
	res.Header().Set("content-type", "application/json")
	_, err = res.Write(data)
	if err != nil {
		log.Printf("Cannot write data to response:%v", err)
		InternalServerErrorResponse(res)
	}
}

func (cc *ContactController) Contacts(res http.ResponseWriter, req *http.Request) {
	query := req.URL.Query()
	pageSizeStr := query.Get("length")
	var pageSize, offset int
	if pageSizeStr == "" {
		pageSize = constants.PAGE_SIZE
	} else {
		pageSize = convertHttpParam2Int(pageSizeStr, constants.PAGE_SIZE)
	}
	offsetStr := query.Get("start")
	if offsetStr == "" {
		offset = constants.OFFSET
	} else {
		offset = convertHttpParam2Int(offsetStr, constants.OFFSET)
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

	id, err := IdFromRequest(req, 2)
	if err != nil {
		log.Printf("Can not get Id from the url: %s, %v", req.URL.Path, err)
		log.Printf("Contact Id must be specified for Contact Update")
		BadRequestResponse(res)
		return
	}

	log.Printf("Update Contact %d", id)
	contactToUpdate := new(simplewebapp.Contact)
	if err := json.NewDecoder(req.Body).Decode(contactToUpdate); err != nil {
		log.Printf("Error updating Contact: %v", err)
	}
	contactToUpdate.Id = id

	contactToUpdate, err, updatedCount := cc.ContactService.UpdateContact(contactToUpdate)
	if err != nil {
		InternalServerErrorResponse(res)
		return
	}
	if updatedCount < 1 {
		log.Printf("Contact not found for update: %d", id)
		http.NotFound(res, req)
		return
	}

	data, err := json.Marshal(contactToUpdate)
	if err != nil {
		log.Printf("Can not marshall data: %v", err)
		InternalServerErrorResponse(res)
		return
	}

	res.Header().Set("content-type", "application/json")
	_, err = res.Write(data)
	if err != nil {
		log.Printf("Can not write data to response: %v", err)
		InternalServerErrorResponse(res)
	}
}

func (cc *ContactController) DeleteContact(res http.ResponseWriter, req *http.Request) {
	id, err := IdFromRequest(req, 2)
	if err != nil {
		log.Printf("Cannot get Id from the url: %s, %v", req.URL.Path, err)
		BadRequestResponse(res)
		return
	}
	log.Printf("Delete Contact %d", id)

	deletedContact, err := cc.ContactService.DeleteContact(id)
	if err != nil {
		InternalServerErrorResponse(res)
		return
	}
	if deletedContact == nil {
		http.NotFound(res, req)
		return
	}

	resultContact, err := json.Marshal(deletedContact)
	if err != nil {
		log.Printf("Cannot return deleted Contact %d", id)
	}

	res.Header().Set("content-type", "application/json")
	_, err = res.Write(resultContact)

	if err != nil {
		log.Printf("Can not write data to response: %v", err)
	}
}

func NewContactController(cs simplewebapp.ContactService) *ContactController {
	return &ContactController{
		ContactService: cs,
	}
}

func convertHttpParam2Int(srcStr string, defVal int) int {
	res, err := strconv.Atoi(srcStr)
	if err != nil {
		log.Printf("Cannot convert http param %s to int: %v, using default value %d", srcStr, err, defVal)
		return defVal
	}
	return res
}
