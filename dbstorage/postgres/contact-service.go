package postgres

import (
	"database/sql"
	"log"
	"simple-web-app"
)

type ContactServiceImpl struct {
	DB *sql.DB
}

//ContactService implementation
func (cs *ContactServiceImpl) Contact(id int) (*simplewebapp.Contact, error) {
	log.Printf("'Contact' method invoked.")
	return nil, nil
}

func (cs *ContactServiceImpl) Contacts(start int, offset int) ([]*simplewebapp.Contact, error) {
	log.Printf("'List of Contacts' method invoked.")
	return nil, nil
}

func (cs *ContactServiceImpl) CreateContact(c *simplewebapp.Contact) (*simplewebapp.Contact, error) {
	panic("implement me")
}

func (cs *ContactServiceImpl) UpdateContact(c *simplewebapp.Contact) (*simplewebapp.Contact, error) {
	panic("implement me")
}

func (cs *ContactServiceImpl) DeleteContact(id int) (*simplewebapp.Contact, error) {
	panic("implement me")
}

func NewContactServiceImpl(db *sql.DB) simplewebapp.ContactService {
	return &ContactServiceImpl{
		DB: db,
	}
}
