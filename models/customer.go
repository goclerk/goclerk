package models

type Customer struct {
	OrganizationId int64
	Id             int64
	DisplayName    string
	FirstName      string
	LastName       string
	CompanyName    string
	Address
	Email       string
	PhoneNumber string
	VATNumber   string
}

type Address struct {
	Address    string
	PostalCode string
	City       string
	Country    string
}
