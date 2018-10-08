package main

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

func connectDb(uri string) (*gorm.DB, error) {
	db, err := gorm.Open("sqlite3", uri)
	if err != nil {
		return nil, err
	}

	return db, nil
}
