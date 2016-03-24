package models

type User struct {
	OrganizationId int64
	Organization   *Organization
	Id             int64
	Username       string
	Email          string
	Password       string
}
