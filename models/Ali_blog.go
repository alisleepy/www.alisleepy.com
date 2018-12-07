/**
 * Created by Goland.
 * User: wangkaikai
 * Date: 2018/12/05
 * Time: 21:03
 */
package models

import (
	db "www.alisleepy.com/database"
	"log"
	"fmt"
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
func GetTopBlogList()(blogs []Ali_blog){
	blogs = make([]Ali_blog,0) //定义一个切片存放数据
	//查询推荐博客
	rows, err := db.SqlDB.Query("select bId,bTitle,bInfo,vViews,vReply_num from ali_blog where bStatus = 1 and is_top = 1 order by add_time desc limit 1")
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
//获取单篇文章
func GetBlogInfoData(id int)(b *Ali_blog){
	fmt.Println(id)
	var blog Ali_blog
	err := db.SqlDB.QueryRow("select * from ali_blog where bId = ?", id).Scan(
		&blog.BId,
		&blog.AId,
		&blog.CatId,
		&blog.BTitle,
		&blog.BInfo,
		&blog.BPic,
		&blog.BContent,
		&blog.LId,
		&blog.Is_top,
		&blog.Add_time,
		&blog.Update_time,
		&blog.VViews,
		&blog.VReply_num,
		&blog.BStatus,
		&blog.AllowReply,
	)
	if err != nil{
		log.Println(err)
		return
	}
	return &blog
}
