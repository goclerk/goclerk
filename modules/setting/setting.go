package setting

import (
	"fmt"
	"github.com/go-ini/ini"
	"os"
)

var (
	Connection *ConnectionDetails
)

type settings struct {
	*ConnectionDetails
}

type ConnectionDetails struct {
	Host     string
	Username string
	Password string
	Database string
	Schema   string
}

func GetConnectionSettings(section *ini.Section) *ConnectionDetails {
	c := &ConnectionDetails{
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

func SaveSettings() {
	s := &settings{
		Connection,
	}
	cfg := ini.Empty()
	err := ini.ReflectFrom(cfg, s)

	if err != nil {
		fmt.Fprintf(os.Stderr, err.Error()+"\n")
		os.Exit(1)
	}

	cfg.SaveTo("settings.ini")
}
