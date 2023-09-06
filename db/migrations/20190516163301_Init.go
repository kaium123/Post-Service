package main

import (
	"gorm.io/gorm"
	"log"
)

type User struct {
	gorm.Model
	Name string
}

// Up is executed when this migration is applied
func Up_20190516163301(txn *gorm.DB) {
	if err := txn.Migrator().CreateTable(&User{}); err != nil {
		log.Fatal(err)
	}
}

// Down is executed when this migration is rolled back
func Down_20190516163301(txn *gorm.DB) {
	if err := txn.Migrator().DropTable(&User{}); err != nil {
		log.Fatal(err)
	}
}
