package db

import (
	"github.com/jinzhu/gorm"
	_ "github.com/mattn/go-sqlite3"
)

var DBTYPE string
var DBNAME string

func Init() {
	// ぶっちゃけ現状ではなくてもいい気がする
	DBTYPE = "sqlite3"
	DBNAME = "test.sqlite3"
}

func GetDB() *gorm.DB {
	db, err := gorm.Open(DBTYPE, DBNAME)
	if err != nil {
		panic("failed to connect database\n")
	}
	return db
}
