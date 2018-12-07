/**
 * Created by Goland.
 * User: wangkaikai
 * Date: 2018/12/05
 * Time: 22:01
 */
package models

import (
	db "www.alisleepy.com/database"
	"log"
)
type Ali_category struct {
	CatId int `json:"catId" form:"catId"`
	CatName string `json:"catName" form:"catName"`
}

func GetCategoryList()(categorys[] *Ali_category){
	datas, err := db.SqlDB.Query("select * from ali_catogory")
	if err != nil{
		log.Println(err)
	}
	defer datas.Close()
	for datas.Next(){
		var category Ali_category
		datas.Scan(&category.CatId, &category.CatName)
		categorys  = append(categorys, category)
	}
	if err = datas.Err(); err != nil {
		return nil
	}
	return categorys
}