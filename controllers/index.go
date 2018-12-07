/**
 * Created by Goland.
 * User: wangkaikai
 * Date: 2018/12/05
 * Time: 21:03
 */
package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"html/template"
	"log"
	"net/http"
	. "www.alisleepy.com/models"
	"strconv"
)

//博客首页
func Index(this *gin.Context){
	//模板文件的拼接
	t, err := template.ParseFiles(
		"views/header.html",
		"views/head.html",
		"views/footer.html",
	)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(t)
	//渲染html文件
	this.HTML(http.StatusOK,"index.html", gin.H{
		"title": "布局页面",
	})
}
//置顶3篇文章
func GetTopBlog(this *gin.Context){
	datas := GetTopBlogList()
	num := len(datas)
	var code int
	if num > 0 {
		code = 200
	} else {
		code = 0
	}
	//获取置顶文章
	this.JSON(http.StatusOK, gin.H{
		"code":code,
		"data":datas,
	})
}
//获取单篇文章
func GetBlogInfo(this *gin.Context){
	bId := this.Query("bId")
	id, err := strconv.Atoi(bId)
	if err != nil{
		log.Fatalln(err)
	}
	var code int
	data := GetBlogInfoData(id)
	if data == nil{
		code = 0
		fmt.Println("blog data is empty")
	}else{
		code = 200
	}
	this.JSON(http.StatusOK, gin.H{
		"code" : code,
		"data" : data,
	})
}
//获取文章分类列表
func GetCategorys(this *gin.Context){
	datas := GetCategoryList()
	num := len(datas)
	fmt.Println(num)
	var code int
	if num >0 {
		code = 200
	}else{
		code = 0
	}
	this.JSON(http.StatusOK, gin.H{
		"code" : code,
		"data" : datas,
	})
}