package main

import (
	"github.com/ginx-contribs/dbx"
	"github.com/ginx-contribs/gormx"
	"log"
)

func main() {
	db, err := gormx.Open(gormx.Options{
		Driver:   dbx.Sqlite,
		Database: "test.db",
	})

	if err != nil {
		log.Fatal(err)
	}

	type User struct {
		Name string `gorm:"type:varchar(128);"`
	}

	err = db.Migrator().AutoMigrate(&User{})
	if err != nil {
		log.Fatal(err)
	}
}
