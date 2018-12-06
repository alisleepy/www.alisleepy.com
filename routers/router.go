package routers

import (
	"github.com/gin-gonic/gin"
	. "www.alisleepy.com/controllers" //constroller部分
)

func InitRouter() *gin.Engine{
	router := gin.Default()

	//首页
	router.GET("/", Index)

	//渲染html页面
	router.LoadHTMLGlob("views/*")
	router.GET("/home/topBlog", GetTopBlog)

	//列表页面
	//router.GET("/home/list", ListHtml)
	//router.POST("/home/PageData", GetDataList)
	//router.POST("/home/PageNextData", PageNextData)
	//
	////新增页面
	//router.GET("home/add", AddHtml)
	//router.POST("home/saveadd", AddPersopnApi)
	//
	////编辑页面
	//router.GET("home/edit", EditHtml);
	//router.POST("home/saveedit", EditPersonApi)
	//
	////删除数据
	//router.GET("home/delete", DeletePersonApi)

	return router
}
