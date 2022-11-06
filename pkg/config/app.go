package config

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
)

var (
	db *gorm.DB
)

func Connect() {
	DBMS := "mysql"
	USER := "go_test"
	PASS := "password"
	PROTOCOL := "tcp(127.0.0.2)"
	DBNAME := "go_database"

	CONNECT := USER + ":" + PASS + "@" + PROTOCOL + "/" + DBNAME + "?charset=utf8&parseTime=true&loc=Asia%2FTokyo"

	d, err := gorm.Open(DBMS, CONNECT)
	if err != nil {
		log.Println(err)
		panic(err)
	}
	db = d
}

func GetDB() *gorm.DB {
	return db
}
