package setting

import (
	"fmt"
	"github.com/go-ini/ini"
	"os"
)

var (
	Connection *connection
)

type connection struct {
	Host     string
	Username string
	Password string
	Database string
	Schema   string
}

func GetConnectionSettings(section *ini.Section) *connection {
	c := &connection{
		Host:     "localhost",
		Schema:   "public",
		Database: "goclerk",
	}

	section.MapTo(c)

	return c
}

func LoadSettings() {
	cfg, err := ini.Load("settings.ini")

	if err != nil {
		fmt.Fprintf(os.Stderr, err.Error()+"\n")
		os.Exit(1)
	}

	Connection = GetConnectionSettings(cfg.Section("database"))
}
