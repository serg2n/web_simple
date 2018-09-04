package postgres

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"simple-web-app"
)

type ContactServiceImpl struct {
	DB *sql.DB
}

const (
	contactsListSql  = "SELECT * FROM contacts LIMIT $1 OFFSET $2"
	contactSql       = "SELECT * FROM contacts WHERE id = $1"
	createContactSql = "INSERT INTO contacts(id, first_name, last_name, phone, email) VALUES ($1, $2, $3, $4, $5)"
	updateContactSql = "UPDATE contacts SET first_name = $1, last_name = $2, phone = $3, email = $4"
	delectContactSql = "DELETE FROM contacts WHERE id = $1"
)

//ContactService implementation
func (cs *ContactServiceImpl) Contact(id int) (*simplewebapp.Contact, error) {
	log.Printf("Get Contact by ID: %d", id)

	contact := new(simplewebapp.Contact)

	err := cs.DB.QueryRow(contactSql, id).Scan(&contact.Id, &contact.FirstName, &contact.LastName, &contact.Phone,
		&contact.Email)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Printf("No contact found with id = %d", id)
			return nil, nil
		}
		log.Printf("Cannot get contact by id %d, error: %v", id, err)
		return nil, errors.New(fmt.Sprintf("Cannot get Contact with id %d from the storage: %v",
			id, err))
	}
	return contact, nil
}

func (cs *ContactServiceImpl) Contacts(pageSize int, offset int) ([]*simplewebapp.Contact, error) {
	log.Printf("List of Contacts, pageSize %d, offset %d", pageSize, offset)
	rows, err := cs.DB.Query(contactsListSql, pageSize, offset)
	if err != nil {
		log.Printf("Can not get list of contacts: %v", err)
		return nil, errors.New(fmt.Sprintf("Can not get list of contacts: %v", err))
	}
	defer rows.Close()

	var resultContacts []*simplewebapp.Contact
	for rows.Next() {
		contact := new(simplewebapp.Contact)
		err := rows.Scan(&contact.Id, &contact.FirstName, &contact.LastName, &contact.Phone, &contact.Email)
		if err != nil {
			log.Printf("Cannot read contact from the storage: %v", err)
			return nil, errors.New(fmt.Sprintf("Cannot read contact from the storage: %v", err))
		}
		resultContacts = append(resultContacts, contact)
	}
	err = rows.Err()
	if err != nil {
		log.Printf("Cannot read contact from the storage: %v", err)
		return nil, errors.New(fmt.Sprintf("Cannot read contact from the storage: %v", err))

	}
	return resultContacts, nil
}

func (cs *ContactServiceImpl) CreateContact(newContact *simplewebapp.Contact) (*simplewebapp.Contact, error) {
	seqVal, err := NextSeqVal()
	if err != nil {
		log.Printf("Cannot create a new Contact, error while getting a new ID: %v", err)
		return nil, errors.New(fmt.Sprintf("Cannot create a new contact: %v", err))
	}

	newContact.Id = seqVal
	stmt, err := cs.DB.Prepare(createContactSql)
	if err != nil {
		msg := fmt.Sprintf("Cannot create a new Contact: %v", err)
		log.Printf(msg)
		return nil, errors.New(msg)
	}

	result, err := stmt.Exec(newContact.Id, newContact.FirstName, newContact.LastName,
		newContact.Phone, newContact.Email)
	if err != nil {
		msg := fmt.Sprintf("Cannot create a new Contact: %v", err)
		log.Printf(msg)
		return nil, errors.New(msg)
	}
	rowCnt, _ := result.RowsAffected()
	log.Printf("A new contact created, %d rows affected", rowCnt)

	return newContact, nil
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
