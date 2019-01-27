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

	router.GET("/home/blogInfo",BlogInfo)            //单个文章页面
	router.GET("/home/getTopBlog", GetTopBlog)       //获取置顶文章列表
	router.GET("/home/getBlogInfo", GetBlogInfo)     //获取单个文章接口
	router.GET("/home/getCategorys", GetCategorys)   //获取文章分类列表
	router.GET("/home/ajaxGetBlogs", AjaxGetBlogs)   //获取普通文章列表
	router.GET("/home/ajaxGetBlogNum", AjaxGetBlogNum)  //获取博客总数
	router.GET("/home/getMyInfo", GetMyInfo)  //获取个人信息
	router.GET("/home/getTopViewBlog", GetTopViewBlogs)  //获取点击最高的几篇文章
	router.GET("/home/getFriendluUrl", GetFriendlyUrl);  //获取友链
	router.GET("/home/getBlogReplys", GetBlogReplys);    //获取文章评论列表
	router.POST("/home/postReply", PostReply)  //发表评论
	return router
}
