package models

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
	"os"
)

var db *gorm.DB

type Model struct {
	ID int `gorm:"primary_key" json:"id"`
	CreatedOn int `json:"created_on"`
	ModifiedOn int `json:"modified_on"`
}


func init()  {

	var (
		err error
		dbType, dbName, user, password, host, tablePrefix string
	)

	dbType = os.Getenv("DB_CONNECTION")
	dbName = os.Getenv("DB_DATABASE")
	user = os.Getenv("DB_USERNAME")
	password = os.Getenv("DB_PASSWORD")
	host = os.Getenv("DB_HOST")
	tablePrefix = os.Getenv("TABLE_PREFIX")


	db, err = gorm.Open(dbType, fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
		user,
		password,
		host,
		dbName))

	if err != nil {
		log.Println(err)
	}


	gorm.DefaultTableNameHandler = func (db *gorm.DB, defaultTableName string) string  {
		return tablePrefix + defaultTableName
	}

	db.SingularTable(true)
	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(100)
}