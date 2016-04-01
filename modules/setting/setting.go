package setting

import "github.com/go-ini/ini"

var (
	Connection *connection
)

type connection struct {
	Host string
	Username string
	Password string
	Database string
	Schema string
}

func GetConnectionSettings(section *ini.Section) *connection {
	c := &connection{
		Host: "localhost",
		Schema: "public",
		Database: "goclerk",
	}

	section.MapTo(c)

	return c
}