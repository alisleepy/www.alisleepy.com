package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"html/template"
	"log"
	"net/http"
	"strconv"
	. "www.alisleepy.com/models"
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
	datas := GetTopBlogs()
	num := len(datas)
	if num > 0 {
		code := 200
	} else {
		code := 0
	}
	//获取置顶文章
	this.JSON(http.StatusOK, gin.H{
		"code":code,
		"data":datas,
	})
}