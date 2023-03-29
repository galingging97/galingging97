package config

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	server := "anakrimba.org"
	port := "3306"
	dbname := "anakrimb_sales"
	username := "anakrimb_demo"
	password := "@ARflutter2022"

	dsn := username + ":" + password + "@tcp(" + server + ":" + port + ")/" + dbname + "?charset=utf8&parseTime=true&loc=Local"
	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{SkipDefaultTransaction: true})
	if err != nil {
		panic("cannot connect database")
	}
}
