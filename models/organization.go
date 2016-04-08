package models

// Organization is the overarching entity over all the data
type Organization struct {
	Model
	Name  string
	Users []User `json:",omitempty"`
}
