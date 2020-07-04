package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gluten/global"
	"gluten/model"
)

func InitGlutenInfoRouter(Router *gin.RouterGroup) {
	GlutenInfoGroup := Router.Group("gluten_info")
	GlutenInfoGroup.GET("/list", SelectAllGlutenInfo)
}

func SelectAllGlutenInfo(c *gin.Context) {
	err, data := model.SelectAllGlutenInfo()
	if err != nil {
		fmt.Println(err)
		global.FailWithMessage(fmt.Sprintf("获取数据失败，%v", err), c)
	} else {
		global.OkWithData(data, c)
	}
}
