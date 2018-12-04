package models

import (
	"fmt"
	db "ginweb/database"
	//db "GinLearn/GinLearn/database"
	"log"
)

//表结构
type Person struct {
	Id        int  `json:"id" form:"id"`
	FirstName string `json:"first_name" form:"first_name"`
	LastName  string `json:"last_name" form:"last_name"`
}

//获取数据列表
func GetPersonList(pageno, pagesize int, search string)(persons []Person){
	fmt.Println("搜索参数："+search)
	persons = make([]Person, 0) //定义一个切片

	//sql查询分页语句
	if search != "" {
		rows, err := db.SqlDB.Query("SELECT id, first_name, last_name FROM person where 1=1 and last_name like '%"+search+"%' or first_name like '%"+search+"%' limit ?,?", (pageno-1)*pagesize, pagesize)
		if err != nil {
			return nil
		}
		defer rows.Close()

		//数据添加到数据集中
		for rows.Next() {
			var person Person
			rows.Scan(&person.Id, &person.FirstName, &person.LastName)
			persons = append(persons, person)
		}
		if err = rows.Err(); err != nil {
			return nil
		}
	}else{
		rows, err := db.SqlDB.Query("select id, first_name, last_name from person where 1=1 limit ?,?",(pageno-1)*pagesize, pagesize)
		if err != nil{
			return nil
		}
		defer rows.Close()

		//数据添加到数据集中
		for rows.Next(){
			var person Person
			rows.Scan(&person.Id, &person.FirstName, &person.LastName)
			persons = append(persons, person)
		}
		if err = rows.Err(); err != nil{
			return nil
		}
	}
	return persons
}
//获取数据条数
func GetRecordNum(search string) int {
	num := 0
	//sql分页查询语句
	if search != ""{
		rows, err := db.SqlDB.Query("SELECT id, first_name, last_name FROM person where 1=1 and first_name like '%?%' or last_name '%?%'",search,search)
		if err != nil{
			return 0
		}
		defer rows.Close()

		//数据添加到数据集中
		for rows.Next(){
			num++
		}
	}else{
		rows, err := db.SqlDB.Query("select id, first_name, last_name form person where 1=1")
		if err != nil{
			return 0
		}
		//数据添加到数据集中
		for rows.Next(){
			num++
		}
	}
	return num
}
//添加数据
func (p *Person) AddPersonData() bool {
	rs, err := db.SqlDB.Exec("INSERT INTO person(first_name, last_name) VALUES (?, ?)", p.FirstName, p.LastName)
	if err != nil{
		return false
	}
	id, err := rs.LastInsertId()
	fmt.Println(id)
	if err != nil{
		return false
	}else{
		return true
	}
}
//获取单个用户数据
func GetPersonData(id int)(p *Person){
	var person Person
	err := db.SqlDB.QueryRow("SELECT * FROM person WHERE id = ?", id).Scan(
		&person.Id, &person.FirstName, &person.LastName,
	)
	if err != nil{
		log.Println(err)
		return
	}
	return &person
}
//更新单个用户的数据
func (p *Person) EditPersonData() bool {
	rs, err := db.SqlDB.Exec("update person set first_name = ?,last_name = ? where id = ?", p.FirstName, p.LastName, p.Id)
	if err != nil {
		return false
	}
	pid, err := rs.RowsAffected()
	fmt.Println(pid)
	if err != nil {
		return false
	}else{
		return true
	}
}
//删除用户
func DelPersonData(id int)bool{
	rs, err := db.SqlDB.Exec("delete from person where id = ?", id)
	if err != nil {
		return false
	}
	pid, err := rs.RowsAffected()
	fmt.Println(pid)
	if err != nil {
		return false
	}else{
		return true
	}
}