package config

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	db *gorm.DB
)

func Connect() {
	dbConnection, err := gorm.Open("mysql", "abhiram:abc@123/simplerest?charset=utf8&parseTime=Trueloc=Local")
	if err != nil {
		panic(err)
	}
	db = dbConnection
}

func GetDBConnection() *gorm.DB {
	return db
}
