package model

import (
	"github.com/jinzhu/gorm"
)

func connect() *gorm.DB{
	DBMS := "mysql"
	USER := "root"
	PASS := "123456"
	PROTOCOL := "tcp(localhost:3306)"
	DBNAME := "test"

	CONNECT := USER + ":" +PASS+"@"+PROTOCOL+"/"+DBNAME

	db,err := 	gorm.Open(DBMS,CONNECT)

	if err == nil{
		panic(err.Error())
	}
	return db;
}

