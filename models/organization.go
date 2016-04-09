package models

import ()

// Organization is the overarching entity over all the data
type Organization struct {
	Model `storm:"inline"`
	Name  string `storm:"unique"`
	Users []User `json:",omitempty"`
}
