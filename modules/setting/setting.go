// Package setting manage settings for the app
package setting

import (
	"fmt"
	"github.com/go-ini/ini"
	"os"
)

var (
	// Connection contains database details
	Connection *ConnectionDetails
)

type settings struct {
	*ConnectionDetails
}

// ConnectionDetails type to store connection details
type ConnectionDetails struct {
	Host     string
	Username string
	Password string
	Database string
	Schema   string
}

// GetConnectionSettings to map connection settings from ini section
func GetConnectionSettings(section *ini.Section) *ConnectionDetails {
	c := &ConnectionDetails{
		Host:     "localhost",
		Schema:   "public",
		Database: "goclerk",
	}

	section.MapTo(c)

	return c
}

// LoadSettings load all the settings from the ini file
func LoadSettings() {
	cfg, err := ini.Load("settings.ini")

	if err != nil {
		fmt.Fprintf(os.Stderr, err.Error()+"\n")
		os.Exit(1)
	}

	Connection = GetConnectionSettings(cfg.Section("database"))
}

// SaveSettings Save the current settings to the ini file
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
