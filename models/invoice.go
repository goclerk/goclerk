package models

import (
	"github.com/siddontang/go/bson"
	"time"
)

// Invoice Invoice model
type Invoice struct {
	Model          `storm:"inline"`
	OrganizationID bson.ObjectId `storm:"index"`
	Number         string
	ContactID      bson.ObjectId `storm:"index"`
	Contact        *Contact
	Address
	VATNumber   string
	Amount      int64
	CreatedAt   time.Time
	UpdatedAt   time.Time
	InvoiceDate time.Time `storm:"index"`
	DueDate     time.Time `storm:"index"`
	PaidDate    time.Time
	note        string
	Details     []InvoiceDetail
	Status      string
}

// InvoiceDetail detail line of an Invoice
type InvoiceDetail struct {
	Quantity    int
	Description string
	Amount      int
	Vat         int
}
