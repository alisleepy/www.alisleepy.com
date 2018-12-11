package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	. "www.alisleepy.com/models"
)
//获取个人信息和浏览次数
func GetMyInfo(this *gin.Context){
	myInfo := GetMyInfosAndViewNum()
	this.JSON(http.StatusOK, gin.H{
		"code":200,
		"data":myInfo,
	})
}
