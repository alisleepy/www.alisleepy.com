package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	. "www.alisleepy.com/models"
	"strconv"
)

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
//获取文章列表
func AjaxGetBlogs(this *gin.Context){
	curpage := this.Query("curpage") //当前页码
	catId := this.Query("catId")  //分类ID
	lableId := this.Query("lId")  //标签ID
	keywords := this.Query("keywords")
	page, err := strconv.Atoi(curpage)
	if err != nil{
		log.Println("page is wrong")
	}
	cId, err := strconv.Atoi(catId)
	if err != nil{
		log.Fatal(err)
	}
	lId, err := strconv.Atoi(lableId)
	if err != nil{
		log.Fatal(err)
	}
	//得到博客列表
	data := GetBlogList(page, cId, lId, keywords)
	var code int
	num := len(data)
	if num >0{
		code = 200
	}else{
		code = 0
	}
	this.JSON(http.StatusOK, gin.H{
		"code":code,
		"data":data,
	})
}