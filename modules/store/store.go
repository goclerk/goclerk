package store

import (
	"fmt"
	"github.com/asdine/storm"
	"github.com/jonaswouters/goclerk/modules/setting"
)

var (
	db *storm.DB
)

// GetDB returns a storm database instance
func GetDB() *storm.DB {
	if db == nil {
		InitializeStore()
	}

	return db
}

// InitializeStore creates a storm instance
func InitializeStore() {
	var err error
	db, err = storm.Open(setting.Settings.Database)

	if err != nil {
		fmt.Printf("Error opening database file %s", setting.Settings.Database)
		return
	}
}
