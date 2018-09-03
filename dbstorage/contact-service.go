package dbstorage

import (
	"database/sql"
	"simple-web-app"
)

type ContactService struct {
	DB *sql.DB
}


//ContactService implementation
func (ContactService) Contact(id int) (*simplewebapp.Contact, error) {
	panic("implement me")
}

func (ContactService) Contacts(start int, offset int) ([]*simplewebapp.Contact, error) {
	panic("implement me")
}

func (ContactService) CreateContact(c *simplewebapp.Contact) (*simplewebapp.Contact, error) {
	panic("implement me")
}

func (ContactService) UpdateContact(c *simplewebapp.Contact) (*simplewebapp.Contact, error) {
	panic("implement me")
}

func (ContactService) DeleteContact(id int) (*simplewebapp.Contact, error) {
	panic("implement me")
}


