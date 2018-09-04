package simplewebapp

type Contact struct {
	ID        int
	FirstName string
	LastName  string
	Phone     string
	Email     string
}

type ContactService interface {
	Contact(id int) (*Contact, error)
	Contacts(pageSize int, offset int) ([]*Contact, error)
	CreateContact(c *Contact) (*Contact, error)
	UpdateContact(c *Contact) (*Contact, error)
	DeleteContact(id int) (*Contact, error)
}
