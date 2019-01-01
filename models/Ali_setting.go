package models

import (
	"log"
	db "www.alisleepy.com/database"
	//"fmt"
)

type Ali_setting struct {
	Id int `json:"sId" form:"sId"`
	Key string `json:"key" form:"key"`
	Value string `json:"value" form:"value"`
}

//获取个人信息
func GetMyInfosAndViewNum()(infos []Ali_setting){
	datas, err := db.SqlDB.Query("select `key`,`value` from ali_setting where `key`='qq' or `key` = 'email' or `key` = 'viewNum'")
	if err != nil{
		log.Println(err)
	}
	defer datas.Close()
	for datas.Next(){
		var uInfo Ali_setting
		datas.Scan(&uInfo.Key, &uInfo.Value)
		infos  = append(infos,uInfo)
	}
	if err = datas.Err(); err != nil {
		return nil
	}
	return infos
}
//获取浏览次数
func GetViewsNum()(view *Ali_setting){
	var viewNum Ali_setting
	err := db.SqlDB.QueryRow("select `value` from ali_setting where `key` = 'viewNum'").Scan(&viewNum)
	if err != nil{
		log.Fatalln(err)
	}
	return &viewNum
}
//浏览次数+1
/*func IncreaseViewNum(){
	rs, err := db.SqlDB.Exec("update ali_setting set first_name = ?,last_name = ? where id = ?", p.FirstName, p.LastName, p.Id)
	if err != nil {

	}
	pid, err := rs.RowsAffected()
	fmt.Println(pid)
}*/
//获取友链
func GetFriendlyUrls()(urlinfo *Ali_setting){
	var friendlyurls Ali_setting
	err := db.SqlDB.QueryRow("select `value` from ali_setting where `key` = 'friendlyurl'").Scan(&friendlyurls.Value)
	if err != nil{
		log.Fatalln(err)
	}
	return &friendlyurls
}