package models

import "github.com/siddontang/go/bson"

// User user to authenticate with
type User struct {
	Model           `storm:"inline"`
	OrganizationIDs []bson.ObjectId `storm:"index"`
	Username        string `storm:"unique"`
	Email           string
	Password        string `json:"-"`
}
