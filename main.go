package main

import (
	"fmt"
	_ "github.com/gin-gonic/gin"
	"gluten/global"
	"gluten/initialize"
)

func init() {
	fmt.Println("init mysql")
	initialize.Mysql()
}

func main() {
	//var data map[string]json.RawMessage
	//a := datatypes.JSON(`{1:2}`)
	//_ = json.Unmarshal(a, &data)
	//fmt.Printf("%s\n", data)
	//info := model.GlutenInfo{Title:"Hash索引和B-Tree索引对比",Star: 0, Post:datatypes.JSON(`{"测试开发": 1}`),
	//	Category: 0, Company:datatypes.JSON(`{"测试开发": 1}`), UserId: 1}
	//model.AddGlutenInfo(info)

	//r.GET("/", func(context *gin.Context) {
	//	type gluten struct {
	//		Title string `json:"title"`
	//		Star int `json:"star"`
	//		Category string `json:"category"`
	//		Company string `json:"company"`
	//	}
	//	data := []gluten{{ "Hash索引和B-Tree索引对比", 1, "MySQL","腾讯"},
	//		{"讲一下线程间的通信", 1, "操作系统", "腾讯"},
	//		{"synchronized 和 Lock 有什么区别", 1, "Java", "阿里"}}
	//
	//	context.JSON(200, data)
	//})
	//_ = r.Run(":8080")
	r := initialize.Routers()
	_ = r.Run(":8090")
	// 程序结束前关闭数据库链接
	defer global.DB.Close()
}
