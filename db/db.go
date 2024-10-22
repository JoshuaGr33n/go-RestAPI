package db

import (
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// var database *gorm.DB

func GetMySQLDB() *gorm.DB {
	dsn := "root:@tcp(localhost:3306)/go_crud?parseTime=true"
	database, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	return database

}

