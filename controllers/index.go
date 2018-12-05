package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"html/template"
	"log"
	"fmt"
)

//博客首页
func Index(this *gin.Context){
	//模板文件的拼接
	t, err := template.ParseFiles(
		"views/public/head.html",
		"views/public/header.html",
		"views/public/footer.html",
	)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(t)
	//渲染html文件
	this.HTML(http.StatusOK,"index.html", gin.H{
		"title": "布局页面",
	})
	//this.String(http.StatusOK, "alisleepy")
}