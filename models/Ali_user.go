package models

import (
	"log"
	"time"
	db "www.alisleepy.com/database"
)

type Ali_user struct {
	UId int `json:"uId" form:"uId"`
	UName string `json:"uName" form:"uName"`
	UPassword string `json:"uPassword" form:"uPassword"`
	UEmail string `json:"uEmail" form:"uEmail"`
	Add_time int `json:"add_time" form:"add_time"`
	UStatus int `json:"uStatus" form:"uStatus"`
}
//添加用户
func AddUser(uName string, uEmail string)int64{
	add_time := time.Now().Unix()
	user := Ali_user{UName:uName, UEmail:uEmail}
	re, err := db.SqlDB.Exec("INSERT INTO ali_user(uName,uEmail,add_time) VALUE (?,?,?)", user.UName, user.UEmail,add_time)
	if err != nil{
		//添加失败
		log.Fatalln("add user fail")
		return 0
	}
	id, err := re.LastInsertId()
	if err != nil{
		log.Fatalln("add user fail")
		return 0
	}
	return id
}
