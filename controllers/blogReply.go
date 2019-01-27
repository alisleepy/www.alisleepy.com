package controllers

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
	. "www.alisleepy.com/models"
)
//获取文章评论列表
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
//发表评论
func PostReply(this *gin.Context){
	bId := this.PostForm("bId")
	uName := this.PostForm("uName")
	uEmail := this.PostForm("uEmail")
	rContent := this.PostForm("rContent")
	id, err := strconv.Atoi(bId)
	if err != nil{
		log.Fatalln(err)
	}
	var isUser int64
	//先添加用户
	isUser = AddUser(uName, uEmail)
	if isUser == 0{
		this.JSON(http.StatusOK, gin.H{
			"code" : -1,
			"msg" : "发表评论失败1",
		})
	}else{
		//添加评论内容
		isReply := AddReply(id, isUser, rContent)
		if isReply == true{
			this.JSON(http.StatusOK, gin.H{
				"code" : 0,
				"msg" : "发表评论成功",
			})
		}else{
			this.JSON(http.StatusOK, gin.H{
				"code" : -1,
				"msg" : "发表评论失败2",
			})
		}
	}
}