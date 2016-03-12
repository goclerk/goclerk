package models

import (
	"time"
)

type Invoice struct {
	OrganizationId int64
	Id             int64
	Number         string
	CustomerId     int64
	Customer       *Customer
	AddressId      int64
	Address        *InvoiceAddress
	Amount         int64
	CreatedAt      time.Time
	UpdatedAt      time.Time
	InvoiceDate    time.Time
	DueDate        time.Time
	PaidDate       time.Time
	note           string
	Rows           []InvoiceRow
	Status         string
}

type InvoiceAddress struct {
	Id int64
	Address
}

type InvoiceRow struct {
	Id          int64
	InvoiceId   int64
	Invoice     *Invoice
	Quantity    int64
	Description string
	Amount      int64
	Vat         int64
}
