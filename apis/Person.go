package apis

import (
	. "ginweb/models"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
	"fmt"
)

//站点首页
func IndexApi(c *gin.Context){
	//输出字符串
	c.String(http.StatusOK, "你好，世界！这是站点首页，哇哈哈")
	//输出json数据
	//c.JSON(http.StatusOK, gin.H{
	//	"status" :200,
	//	"error": nil,
	//})
}
//渲染html页面
func ShowHtmlPage(c *gin.Context)  {
	//输出html
	c.HTML(http.StatusOK, "index.html", gin.H{
		"title" :"GIN:HTML页面",
	})
}
//用户列表页
func ListHtml(c *gin.Context)  {
	c.HTML(http.StatusOK, "list.html", gin.H{
		"title" : "GIN:用户列表页面",
	})
}
//得到用户列表数据
func GetDataList(c *gin.Context){
	//得到请求的参数
	search := c.PostForm("search")  //post提交的数据
	num := c.PostForm("pageno") //数据长度
	pageno, err := strconv.Atoi(num) //类型转换
	if err != nil{
		log.Fatal(err)
	}
	//得到数据集合
	datalist := GetPersonList(pageno, 10, search) //从module中获取数据

	//得到数据的总数
	count := GetRecordNum(search)

	//返回结果
	c.JSON(http.StatusOK, gin.H{
		"datalist" : datalist,
		"count" : count,
		"pagesize" : 10,
		"pageno" : pageno,
	})
}
//得到下一页数据
func PageNextData(c *gin.Context){
	//得到请求参数
	search := c.PostForm("search")
	num := c.PostForm("pageno")
	pageno,err:= strconv.Atoi(num)
	if err!=nil{
		log.Fatalln(err)
	}

	//得到数据集
	datalist := GetPersonList(pageno, 10, search)
	//得到个数
	count := GetRecordNum(search)
	//返回结果
	c.JSON(http.StatusOK, gin.H{
		"datalist" : datalist,
		"count" : count,
		"pagesize" : 10,
		"pageno" : pageno,
	})
}
//新增用户的页面
func AddHtml(c *gin.Context){
	c.HTML(http.StatusOK, "add.html", gin.H{
		"title" :"GIN:HTML添加页面",
	})
}
//新增用户操作
func AddPersopnApi(c *gin.Context){
	//得到前端post值
	first_name := c.PostForm("first_name")
	last_name := c.PostForm("last_name")

	//给结构体赋值
	p := Person{FirstName:first_name, LastName:last_name}
	//写入数据库
	result := p.AddPersonData()
	c.JSON(http.StatusOK, gin.H{
		"success":result,
	})
	//页面跳转
	//c.Redirect(http.StatusOK, "/home/list")
}
//编辑用户页面
func EditHtml(c *gin.Context){
	//获取get参数
	pid := c.Query("id")
	id, err := strconv.Atoi(pid)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(id)
	//获取用户数据
	p := GetPersonData(id)
	if p == nil{
		fmt.Println("未得到数据")
	}else{
		fmt.Println("得到正确数据")
	}
	//渲染到html模板中
	c.HTML(http.StatusOK, "edit.html", gin.H{
		"title" : "GIN：编辑用户页面",
		"id" : p.Id,
		"first_name" : p.FirstName,
		"last_name" : p.LastName,
	})
	//json返回数据
	//c.JSON(http.StatusOK,gin.H{
	//	"id" : p.Id,
	//	"first_name" : p.FirstName,
	//	"last_name" : p.LastName,
	//})
}
//编辑用户信息
func EditPersonApi(c *gin.Context){
	//post参数
	pid := c.PostForm("id")
	first_name := c.PostForm("first_name")
	last_name := c.PostForm("last_name")
	id, err := strconv.Atoi(pid)
	if err != nil {
		log.Fatalln(err)
	}
	data := Person{Id:id,FirstName:first_name, LastName:last_name}
	result := data.EditPersonData()
	c.JSON(http.StatusOK, gin.H{
		"code" : result,
	})
}
//删除用户
func DeletePersonApi(c *gin.Context){
	pid := c.Query("id")
	id, err := strconv.Atoi(pid)
	if err != nil {
		log.Fatalln(err)
	}
	result := DelPersonData(id)
	c.JSON(http.StatusOK,gin.H{
		"code" : result,
	})
}