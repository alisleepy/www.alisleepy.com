package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	. "www.alisleepy.com/models"
)

func GetMyInfo(this *gin.Context){
	myInfo := GetMyInfos()
	this.JSON(http.StatusOK, gin.H{
		"code":200,
		"data":myInfo,
	})
}
