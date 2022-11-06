package config

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
)

var (
	db *gorm.DB
)

func Connect() {
	USER := "matthew"
	PASS := "password1234"
	HOST := "mysql-db"
	PORT := "3306"
	NAME := "bookstore-db"

	CONNECT := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=True", USER, PASS, HOST, PORT, NAME)
	log.Println(CONNECT)
	d, err := gorm.Open("mysql", CONNECT)

	if err != nil {
		log.Println("sibal what is this 9")
		log.Println(err)
		panic(err)
	} else {
		fmt.Sprintf("DB Connected.")
	}
	db = d
}

func GetDB() *gorm.DB {
	return db
}
