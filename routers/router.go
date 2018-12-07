/**
 * Created by Goland.
 * User: wangkaikai
 * Date: 2018/12/05
 * Time: 21:03
 */
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

	//获取置顶文章列表
	router.GET("/home/getTopBlog", GetTopBlog)
	//获取单个文章
	router.GET("/home/getBlogInfo", GetBlogInfo)
	//获取文章分类列表
	router.GET("/home/getCategorys", GetCategorys)
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
