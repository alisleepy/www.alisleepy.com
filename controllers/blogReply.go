package controllers

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
	. "www.alisleepy.com/models"
)

func GetBlogReplys(this *gin.Context){
	bId := this.Query("bId")
	id, err := strconv.Atoi(bId)
	if err != nil{
		log.Fatalln(err)
	}
	var code int
	data := GetBlogReplysBybId(id)
	if data == nil{
		code = 0
		log.Println("blogReplys is empty")
	}else{
		code = 200
	}
	this.JSON(http.StatusOK, gin.H{
		"code" : code,
		"data" : data,
	})
}
