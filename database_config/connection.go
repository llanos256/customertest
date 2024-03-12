package database_config

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)
var secret = "root:maq2006@tcp(localhost:3306)/issuedb?charset=utf8mb4&parseTime=True&loc=Local"
func GetConnection() (db *gorm.DB) {
	if db, err := gorm.Open(mysql.Open(secret), &gorm.Config{}); err != nil {
		fmt.Println(err)
		panic(err)
	} else {
		fmt.Println("succcessful")
		return db
	}
}