package models

import (
	db "www.alisleepy.com/database"
)

//定义一个结构体
type Ali_reply struct {
	RId int `json:"rId" form:"rId"`
	UId int `json:"uId" form:"uId"`
	BId int `json:"bId" form:"bId"`
	RContent string `json:"rContent" form:"rContent"`
	Add_time int `json:"add_time" form:"add_time"`
	RStatus int `json:"rStatus" form:"rStatus"`
	UName string `json:"uName"`
}
func GetBlogReplysBybId(id int)(replys []Ali_reply){
	replys = make([]Ali_reply, 0) //切片
	rows, err := db.SqlDB.Query("select r.*,u.uName from ali_reply as r " +
								"INNER JOIN ali_blog AS b ON b.bId = r.bId " +
								"INNER JOIN ali_user AS u ON r.uId = u.uId " +
								"WHERE b.bId = ? limit 5", id)
	defer rows.Close()
	for rows.Next(){
		var reply Ali_reply   //定义一个结构体类型的
		rows.Scan(&reply.RId, &reply.UId, &reply.BId, &reply.RContent, &reply.Add_time, &reply.RStatus, &reply.UName)
		replys = append(replys, reply)
	}
	if err = rows.Err(); err != nil {
		return nil
	}
	return replys
}