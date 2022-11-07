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
	PROTOCOL := "tcp(go-bookstoreDB:3306)"
	DBNAME := "go_database"

	CONNECT := USER + ":" + PASS + "@" + PROTOCOL + "/" + DBNAME + "?charset=utf8&parseTime=true&loc=Asia%2FTokyo"

	d, err := gorm.Open(DBMS, CONNECT)
	if err != nil {
		panic(err)
	} else {
		log.Println("DB CONNECTED")
	}
	log.Println("db = d")
	db = d
}

func GetDB() *gorm.DB {
	return db
}
