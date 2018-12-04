package main

import (
	db "ginweb/database"
	"ginweb/routers"
)

func main(){
	//数据库部分
	defer db.SqlDB.Close()

	//路由部分
	router := routers.InitRouter()

	//静态资源
	router.Static("/static", "./static")

	//运行端口
	router.Run(":8000")
}