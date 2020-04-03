package db

import (
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
)

// Db gorm db
var Db *gorm.DB

func init() {
	var err error
	godotenv.Load("../../.env")

	myEnv, err := godotenv.Read()
	if err != nil {
		log.Panicln("err:", err.Error())
	}

	connString := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		myEnv["DATABASE_USER"],
		myEnv["DATABASE_PASSWORD"],
		myEnv["DATABASE_HOST"],
		myEnv["DATABASE_PORT"],
		myEnv["DATABASE_DB"],
	)

	Db, err = gorm.Open("mysql", connString)
	if err != nil {
		log.Panicln("err:", err.Error())
	}

	Db.SingularTable(true)
	Db.DB().SetMaxOpenConns(10)
	Db.DB().SetMaxIdleConns(10)
}
