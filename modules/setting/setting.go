// Package setting manage settings for the app
package setting

import (
	"fmt"
	"github.com/go-ini/ini"
	"github.com/unrolled/render"
	"os"
)

var (
	// Settings contains all the settings
	Settings *settings

	// Renderer is the output renderer being used
	Renderer *render.Render
)

type settings struct {
	// Database is the database filename
	Database string
	// Folder to export the database to
	ExportFolder string
	//Folder to import data into the database
	ImportFolder string
}

// LoadSettings load all the settings from the ini file
func LoadSettings() {
	s := &settings{
		Database:     "database.db",
		ExportFolder: "export/",
	}

	Renderer = render.New(render.Options{IndentJSON: true})

	cfg, err := ini.LooseLoad("settings.ini")

	cfg.MapTo(s)

	Settings = s

	if err != nil {
		fmt.Fprintf(os.Stderr, err.Error()+"\n")
		os.Exit(1)
	}
}

// SaveSettings Save the current settings to the ini file
func SaveSettings() {
	cfg := ini.Empty()
	err := ini.ReflectFrom(cfg, Settings)

	if err != nil {
		fmt.Fprintf(os.Stderr, err.Error()+"\n")
		os.Exit(1)
	}

	cfg.SaveTo("settings.ini")
}
