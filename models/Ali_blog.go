/**
 * Created by Goland.
 * User: wangkaikai
 * Date: 2018/12/05
 * Time: 21:03
 */
package models

import (
	"fmt"
	"log"
	"strconv"
	db "www.alisleepy.com/database"
)

//定义每页数目
const pagesize  = "5"

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
	CatName string `json:"catName" form:"catName"`
	LName string `json:"lName" form:"lName"`
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
//获取博客列表
func GetBlogList(page int, cId int, lId int, keywords string)(blogs []Ali_blog){
	blogs = make([]Ali_blog,0) //定义一个切片存放数据
	var start int
	var offset = pagesize
	start = (page-1)*2
	newStart := strconv.Itoa(start) //int转string（真tm烦，不转的话下边拼接sql语句报错，string不能和int拼接）
	if cId > 0 && lId == 0{ //存在cId ,不存在lID
		rows, err := db.SqlDB.Query("SELECT blog.bId,blog.catId,blog.bTitle,blog.bInfo,blog.bPic,blog.bContent," +
			"blog.lId,blog.add_time,blog.vViews,blog.vReply_num,blog.allowReply,cat.*,lab.* FROM ali_blog AS blog " +
			"INNER JOIN ali_category AS cat ON blog.catId = cat.catId " +
			"INNER JOIN ali_label AS lab ON blog.lId = lab.lId " +
			"WHERE blog.bStatus = 1  " +
			"WHERE blog.catId = cat.catId  " +
			"ORDER BY is_top DESC,vViews DESC LIMIT "+newStart+","+offset)
		defer rows.Close()
		for rows.Next(){
			var blog Ali_blog   //定义一个结构体类型的
			rows.Scan(&blog.BId, &blog.CatId, &blog.BTitle, &blog.BInfo, &blog.BPic, &blog.BContent, &blog.LId,
				&blog.Add_time, &blog.VViews, &blog.VReply_num, &blog.AllowReply, &blog.CatId, &blog.CatName, &blog.LId, &blog.LName,)
			blogs = append(blogs, blog)
		}
		if err = rows.Err(); err != nil {
			return nil
		}
	}else if lId >0 && cId==0{ //存在lId，不存在cId
		rows, err := db.SqlDB.Query("SELECT blog.bId,blog.catId,blog.bTitle,blog.bInfo,blog.bPic,blog.bContent," +
			"blog.lId,blog.add_time,blog.vViews,blog.vReply_num,blog.allowReply,cat.*,lab.* FROM ali_blog AS blog " +
			"INNER JOIN ali_category AS cat ON blog.catId = cat.catId " +
			"INNER JOIN ali_label AS lab ON blog.lId = lab.lId " +
			"WHERE blog.bStatus = 1  " +
			"WHERE blog.lId = lab.lId  " +
			"ORDER BY is_top DESC,vViews DESC LIMIT "+newStart+","+offset)
		defer rows.Close()
		for rows.Next(){
			var blog Ali_blog   //定义一个结构体类型的
			rows.Scan(&blog.BId, &blog.CatId, &blog.BTitle, &blog.BInfo, &blog.BPic, &blog.BContent, &blog.LId,
				&blog.Add_time, &blog.VViews, &blog.VReply_num, &blog.AllowReply, &blog.CatId, &blog.CatName, &blog.LId, &blog.LName,)
			blogs = append(blogs, blog)
		}
		if err = rows.Err(); err != nil {
			return nil
		}
	}else if keywords != "" && lId == 0 && cId ==0{ //只有keywords
		rows, err := db.SqlDB.Query("SELECT blog.bId,blog.catId,blog.bTitle,blog.bInfo,blog.bPic,blog.bContent," +
			"blog.lId,blog.add_time,blog.vViews,blog.vReply_num,blog.allowReply,cat.*,lab.* FROM ali_blog AS blog " +
			"INNER JOIN ali_category AS cat ON blog.catId = cat.catId " +
			"INNER JOIN ali_label AS lab ON blog.lId = lab.lId " +
			"WHERE blog.bStatus = 1  " +
			"AND (blog.bTitle LIKE %?% OR blog.bInfo LIKE %?% OR blog.bContent LIKE %?%)" +
			"ORDER BY is_top DESC,vViews DESC LIMIT "+newStart+","+offset, keywords)
		defer rows.Close()
		for rows.Next(){
			var blog Ali_blog   //定义一个结构体类型的
			rows.Scan(&blog.BId, &blog.CatId, &blog.BTitle, &blog.BInfo, &blog.BPic, &blog.BContent, &blog.LId,
				&blog.Add_time, &blog.VViews, &blog.VReply_num, &blog.AllowReply, &blog.CatId, &blog.CatName, &blog.LId, &blog.LName,)
			blogs = append(blogs, blog)
		}
		if err = rows.Err(); err != nil {
			return nil
		}
	}else{
		//初始情况走到这儿
		rows, err := db.SqlDB.Query("SELECT blog.bId,blog.catId,blog.bTitle,blog.bInfo,blog.bPic,blog.bContent," +
			"blog.lId,blog.add_time,blog.vViews,blog.vReply_num,blog.allowReply,cat.*,lab.* FROM ali_blog AS blog " +
			"INNER JOIN ali_category AS cat ON blog.catId = cat.catId " +
			"INNER JOIN ali_label AS lab ON blog.lId = lab.lId " +
			"WHERE blog.bStatus = 1  ORDER BY is_top DESC,vViews DESC LIMIT "+newStart+","+offset)
		if err != nil{
			return nil
		}
		defer rows.Close()
		for rows.Next(){
			var blog Ali_blog   //定义一个结构体类型的
			rows.Scan(&blog.BId, &blog.CatId, &blog.BTitle, &blog.BInfo, &blog.BPic, &blog.BContent, &blog.LId,
				&blog.Add_time, &blog.VViews, &blog.VReply_num, &blog.AllowReply, &blog.CatId, &blog.CatName, &blog.LId, &blog.LName,)
			blogs = append(blogs, blog)
		}
		if err = rows.Err(); err != nil {
			return nil
		}
	}
	//循环修改添加时间为日期格式
	//for k, v := range blogs{
	//	blogs[k]["add_time"] = time.Unix(v["add_time"], 0)
	//}
	return blogs
}
//获取博客总数
func GetBlogNum()int{
	blogNum, err := db.SqlDB.Query("select count(bId) as num from ali_blog where bStatus = 1")
	if err != nil{
		log.Fatalln(err)
	}
	fmt.Println(blogNum)
	defer blogNum.Close()
	return 1
}
