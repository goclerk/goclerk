package models

import "github.com/siddontang/go/bson"

// Organization is the overarching entity over all the data
type Organization struct {
	Model `storm:"inline"`
	Name  string `storm:"unique"`
}

// OrganizationUsers links users to organizations
type OrganizationUsers struct {
	Model `storm:"inline"`
	OrganizationId bson.ObjectId `storm:"index"`
	UserID bson.ObjectId `storm:"index"`
}
