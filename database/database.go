package database

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

var SqlDB * sql.DB

func init()  {
	var err error
	SqlDB, err = sql.Open("mysql", "root:root@tcp(127.0.0.1)/www.wangkaikai.cn?charset=utf8")
	if err != nil{
		log.Fatal(err.Error())
	}
	err = SqlDB.Ping()
	if err != nil{
		log.Fatal(err.Error())
	}
}
