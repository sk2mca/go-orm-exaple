package database

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"local.com/rest/db/orm/modal"
)

const DB_USERNAME = "root"
const DB_PASSWORD = "Temp!123"
const DB_NAME = "gotest"
const DB_HOST = "mysql"
const DB_PORT = "3306"

var Db *gorm.DB

func InitDb() *gorm.DB {
	Db = connectDB()
	return Db
}

func connectDB() *gorm.DB {
	var err error
	dsn := DB_USERNAME + ":" + DB_PASSWORD + "@tcp" + "(" + DB_HOST + ":" + DB_PORT + ")/" + DB_NAME + "?" + "parseTime=true&loc=Local"

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Error connecting to database")
	}
	db.AutoMigrate(&modal.Order{}, &modal.Item{})
	return db
}
