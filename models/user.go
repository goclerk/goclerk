package models

// User user to authenticate with
type User struct {
	Model           `storm:"inline"`
	OrganizationIDs []int64
	Username        string `storm:"unique"`
	Email           string
	Password        string
}
