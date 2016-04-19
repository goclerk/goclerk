package models

// User user to authenticate with
type User struct {
	Model           `storm:"inline"`
	Username        string `storm:"unique"`
	Email           string
	Password        string `json:"-"`
}
