package main

import (
	db "www.alisleepy.com/database"
	"www.alisleepy.com/routers"
)

//入口部分
func main(){
	//1、执行数据库部分
	defer db.SqlDB.Close()

	//2、执行路由部分
	router := routers.InitRouter()

	//3、设置静态资源
	router.Static("/static", "./static")

	//4、设置模板文件目录
	router.LoadHTMLGlob("views/*")

	//运行端口
	router.Run(":8888")
}