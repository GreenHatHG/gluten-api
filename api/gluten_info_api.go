package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gluten/middleware"
	"gluten/model"
	"gluten/util"
)

func InitGlutenInfoRouter(Router *gin.RouterGroup) {
	GlutenInfoGroup := Router.Group("gluten_info")
	GlutenInfoGroup.GET("/list", SelectAllGlutenInfo).Use(middleware.Auth())
}

func SelectAllGlutenInfo(c *gin.Context) {
	err, data := model.SelectAllGlutenInfo()
	if err != nil {
		fmt.Println(err)
		util.FailWithMessage(fmt.Sprintf("获取数据失败，%v", err), c)
	} else {
		util.OkWithData(data, c)
	}
}
