package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

//博客首页
func Index(this *gin.Context){
	this.HTML(http.StatusOK, "index.html", gin.H{
		"title" : "alisleepy首页",
	})
}