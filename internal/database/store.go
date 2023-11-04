package database

import (
	"log"
	"project/internal/model"
)

func CreateTable() {
	// Drop the table student if it exists
	db, err := Open()
	if err != nil {
		log.Fatalln(err)
	}
	err = db.Migrator().DropTable(&model.User{})
	if err != nil {
		log.Fatalln(err)
	}
	err = db.Migrator().AutoMigrate(&model.User{})
	if err != nil {
		log.Fatalln(err)
	}
}
