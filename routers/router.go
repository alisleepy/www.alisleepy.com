package routers

import (
	"github.com/gin-gonic/gin"
	. "ginweb/apis" //api部分
	//. "ginweb/controllers" //constroller部分
)

func InitRouter() *gin.Engine{
	router := gin.Default()

	//hellow alisleepy页面
	router.GET("/", IndexApi)

	//渲染html页面
	router.LoadHTMLGlob("views/*")
	router.GET("/home/index", ShowHtmlPage)

	//列表页面
	router.GET("/home/list", ListHtml)
	router.POST("/home/PageData", GetDataList)
	router.POST("/home/PageNextData", PageNextData)

	//新增页面
	router.GET("home/add", AddHtml)
	router.POST("home/saveadd", AddPersopnApi)

	//编辑页面
	router.GET("home/edit", EditHtml);
	router.POST("home/saveedit", EditPersonApi)

	//删除数据
	router.GET("home/delete", DeletePersonApi)

	return router
}
