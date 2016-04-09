package store

import "github.com/asdine/storm"

var (
	DB *storm.DB
)

func InitializeStore() {
	DB, _ = storm.Open("my.db")
	//defer DB.Close()
}
