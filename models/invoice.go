package models

import (
	"time"
)

type Invoice struct {
	OrganizationId int
	Id             int
	Number         string
	CustomerId     int
	Customer       *Customer
	Address
	Amount      int
	CreatedAt   time.Time
	UpdatedAt   time.Time
	InvoiceDate time.Time
	DueDate     time.Time
	PaidDate    time.Time
	note        string
	Details     []InvoiceDetail
	Status      string
}

type InvoiceDetail struct {
	Id          int
	InvoiceId   int
	Invoice     *Invoice
	Quantity    int
	Description string
	Amount      int
	Vat         int
}
