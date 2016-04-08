package models

// User user to authenticate with
type User struct {
	Model
	OrganizationIDs []int64
	Username        string
	Email           string
	Password        string
}
