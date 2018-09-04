package simplewebapp

type Contact struct {
	Id        int
	FirstName string
	LastName  string
	Phone     string
	Email     string
}

type ContactService interface {
	Contact(id int) (*Contact, error)
	Contacts(pageSize int, offset int) ([]*Contact, error)
	CreateContact(newContact *Contact) (*Contact, error)
	UpdateContact(contact *Contact) (*Contact, error, int)
	DeleteContact(id int) (*Contact, error)
}
