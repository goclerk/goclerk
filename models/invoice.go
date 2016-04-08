package models

import (
	"time"
)

// Invoice Invoice model
type Invoice struct {
	Model
	OrganizationID int64
	Number         string
	ContactID      int64
	Contact        *Contact
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

// InvoiceDetail detail line of an Invoice
type InvoiceDetail struct {
	ID          int64
	InvoiceID   int64
	Invoice     *Invoice
	Quantity    int
	Description string
	Amount      int
	Vat         int
}
