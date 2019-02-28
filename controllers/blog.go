package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	. "www.alisleepy.com/models"
	"strconv"
)
//单篇文章页面
func BlogInfo(this *gin.Context){
	bId := this.Query("bId")
	this.HTML(http.StatusOK, "info.html", gin.H{
		"title" : "文章详情",
		"bId" : bId,
	})
}
//获取单篇文章
func GetBlogInfo(this *gin.Context){
	bId := this.Query("bId")
	fmt.Println(bId)
	id, err := strconv.Atoi(bId)
	if err != nil{
		log.Fatalln(err)
	}
	var code int
	fmt.Println(id)
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
		"catId":catId,
		"lId":lableId,
		"page":curpage,
		"keywords":keywords,
	})
}
//获取博客总数
func AjaxGetBlogNum(this *gin.Context){
	//获取博客总数
	blogNum := GetBlogNum()
	this.JSON(http.StatusOK, gin.H{
		"code":200,
		"data":blogNum,
	})
}
//获取点击排行前5的文章
func GetTopViewBlogs(this *gin.Context){
	datas := GetTopViewsBlogs()
	var code int
	num := len(datas)
	if num > 0{
		code = 200
	}else{
		code = 0
	}
	this.JSON(http.StatusOK, gin.H{
		"code" : code,
		"data" : datas,
	});
}
//添加文章
func AddBlog(this *gin.Context){
	this.HTML(http.StatusOK, "addblog.html", gin.H{
		"title" : "添加文章",
	})
}