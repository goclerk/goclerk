package models

// Contact organization contact
type Contact struct {
	OrganizationID int64
	ID             int64
	DisplayName    string
	FirstName      string
	LastName       string
	CompanyName    string
	Addresses      []ContactAddress `json:",omitempty"`
	Email          string
	PhoneNumber    string
	VATNumber      string
}

// ContactAddress address of a Customer
type ContactAddress struct {
	ID        int64
	ContactID int64
	Address
}
