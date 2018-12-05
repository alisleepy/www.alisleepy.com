package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"html/template"
	"log"
	"net/http"
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