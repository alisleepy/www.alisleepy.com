package models

import (
	"log"
	db "www.alisleepy.com/database"
)

type Ali_setting struct {
	Id int `json:"sId" form:"sId"`
	Key int `json:"key" form:"key"`
	Value int `json:"value" form:"value"`
}

//获取个人信息
func GetMyInfos()(info *Ali_setting){
	var userInfo Ali_setting
	err := db.SqlDB.QueryRow("select `key`,`value` from ali_setting where `key` = 'qq' or `key` = 'email'").Scan(&userInfo.Key,&userInfo.Value)
	if err != nil{
		log.Fatalln("userInfo is empty")
		return
	}
	return &userInfo
}