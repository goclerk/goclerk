package models

import "github.com/siddontang/go/bson"

// Contact organization contact
type Contact struct {
	Model          `storm:"inline"`
	OrganizationID bson.ObjectId `storm:"index"`
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
	ContactID bson.ObjectId `storm:"index"`
	Address
}
