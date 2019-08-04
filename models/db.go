package models

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func Connect() *gorm.DB {
	db, err := gorm.Open("mysql", "root:adminmysql@tcp(127.0.0.1:3306)/passenger_feedback?parseTime=True")
	if err != nil {
		panic("failed to initiate connection" + err.Error())
	}
	return db
}
