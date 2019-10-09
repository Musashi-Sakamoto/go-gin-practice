package models

import (
	"fmt"
	"log"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"app/pkg/setting"
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
		dbType, dbName, user, host, dbPassword, port, tablePrefix string
	)

	dbType = setting.Config.DBType
	dbName = setting.Config.DBName
	dbPassword = setting.Config.DBPassword
	user = setting.Config.DBUser
	host = setting.Config.DBHost
	port = setting.Config.DBPort
	tablePrefix = setting.Config.TablePrefix

	fmt.Println(host, user, dbName, dbPassword, port)

	db, err = gorm.Open(dbType, fmt.Sprintf("host=%s user=%s dbname=%s password=%s port=%s sslmode=disable", host, user, dbName, dbPassword, port))

	fmt.Println(db.DB().Ping())

	if err != nil {
		log.Println(err)
	}

	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		return tablePrefix + defaultTableName
	}

	db.SingularTable(true)
	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(100)
}

func CloseDB() {
	defer db.Close()
}