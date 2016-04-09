package store

import (
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
	db, _ = storm.Open(setting.Settings.Database)
}
