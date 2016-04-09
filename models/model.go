package models

import "github.com/siddontang/go/bson"

// Model main entity for models
type Model struct {
	ID bson.ObjectId `storm:"id"`
}
