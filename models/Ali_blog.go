package models

import (
	db "www.alisleepy.com/database"
)

type Ali_blog struct {
	BId int `json:"bId" form:"bId"`
	AId int `json:"aId" form:"aId"`
	CatId int `json:"catId" form:"catId"`
	BTitle string `json:"bTitle" form:"bTitle"`
	BInfo string `json:"bInfo" form:"bInfo"`
	BPic string `json:"bPic" form:"bPic"`
	BContent string `json:"bContent" form:"bContent"`
	LId int `json:"lId" form:"lId"`
	Is_top int `json:"is_top" form:"is_top"`
	Add_time int `json:"add_time" form:"add_time"`
	Update_time int `json:"update_time" form:"update_time"`
	VViews int `json:"vViews" form:"vViews"`
	VReply_num int `json:"vReply_num" form:"vReply_num"`
	BStatus int `json:"bStatus" form:"bStatus"`
	AllowReply int `json:"allowReply" form:"allowReply"`
}

//获取推荐文章，3条
func GetTopBlogs()(blogs []Ali_blog){
	blogs = make([]Ali_blog,0) //定义一个切片存放数据
	//查询推荐博客
	rows, err := db.SqlDB.Query("select bId,bTitle,bInfo,vViews,vReply_num from ali_blog where bStatus = 1 and is_top = 1 limit 3")
	if err != nil{
		return nil
	}
	//这一句没明白
	defer rows.Close()
	for rows.Next(){
		var blog Ali_blog   //定义一个结构体类型的
		rows.Scan(&blog.BId,&blog.BTitle,&blog.BInfo,&blog.VViews,&blog.VReply_num)
		blogs = append(blogs, blog)
	}
	if err = rows.Err(); err != nil {
		return nil
	}
	return blogs
}
