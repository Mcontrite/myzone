package model

import (
	"fmt"
	"log"
	"myzone/package/setting"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var db *gorm.DB

type Model struct {
	gorm.Model
}

const PAGE_SIZE int = 10

func init() {
	var (
		err                                               error
		dbType, dbName, user, password, host, tablePrefix string
	)
	dbType = setting.DatabaseSetting.Type
	dbName = setting.DatabaseSetting.Name
	user = setting.DatabaseSetting.User
	password = setting.DatabaseSetting.Password
	host = setting.DatabaseSetting.Host
	tablePrefix = setting.DatabaseSetting.TablePrefix
	db, err = gorm.Open(dbType, fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
		user,
		password,
		host,
		dbName))
	if err != nil {
		log.Println(err)
	}
	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		return tablePrefix + defaultTableName
	}
	db.SingularTable(true)
	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(100)
	db.AutoMigrate(&Group{})
	db.AutoMigrate(&User{})
	db.AutoMigrate(&Article{})
	db.AutoMigrate(&Reply{})
	db.AutoMigrate(&Attach{})
	db.AutoMigrate(&Saying{})
	db.AutoMigrate(&Comment{})
	db.AutoMigrate(&MyArticle{})
	db.AutoMigrate(&MySaying{})
	db.AutoMigrate(&MyFavourite{})
}

func CloseDB() {
	defer db.Close()
}

func GetDb() *gorm.DB {
	return db
}
