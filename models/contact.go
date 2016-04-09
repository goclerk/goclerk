package models

// Contact organization contact
type Contact struct {
	Model          `storm:"inline"`
	OrganizationID int64 `storm:"index"`
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
	Model     `storm:"inline"`
	ContactID int64 `storm:"index"`
	Address
}
