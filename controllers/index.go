/**
 * Created by Goland.
 * User: wangkaikai
 * Date: 2018/12/05
 * Time: 21:03
 */
package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"html/template"
	"log"
	"net/http"
	. "www.alisleepy.com/models"
	//"strconv"
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
		"title" : "alisleepy小站",
		"curpage" : 1, //当前页数
		"catId" : 0, //分类ID
		"lId" : 0, //标签ID
		"keywords" : "", //搜索关键字
	})
}
//置顶3篇文章
func GetTopBlog(this *gin.Context){
	datas := GetTopBlogList()
	num := len(datas)
	var code int
	if num > 0 {
		code = 200
	} else {
		code = 0
	}
	//获取置顶文章
	this.JSON(http.StatusOK, gin.H{
		"code":code,
		"data":datas,
	})
}