package models

type Organization struct {
	Id    int64
	Name  string
	Users []User `json:",omitempty"`
}
