package models

type Customer struct {
	OrganizationId int
	Id             int
	DisplayName    string
	FirstName      string
	LastName       string
	CompanyName    string
	Addresses      []CustomerAddress `json:",omitempty"`
	Email          string
	PhoneNumber    string
	VATNumber      string
}

type CustomerAddress struct {
	Id         int64
	CustomerId int64
	Address
}
