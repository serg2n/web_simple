package simple_web_app

type Contact struct {
	ID        int
	firstName string
	lastName  string
	phone     string
	email     string
}

type ContactService interface {
	Contact(id int) (*Contact, error)
	Contacts(start int, offset int) ([]*Contact, error)
	CreateContact(c *Contact) (*Contact, error)
	UpdateContact(c *Contact) (*Contact, error)
	DeleteContact(id int) (*Contact, error)
}
