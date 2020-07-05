package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/medivhzhan/weapp/v2"
	"gluten/global"
	"gluten/model"
)

func InitUserInfoRouter(Router *gin.RouterGroup) {
	GlutenInfoGroup := Router.Group("user_info")
	GlutenInfoGroup.POST("/action/login", AddUserInfo)
}

func AddUserInfo(c *gin.Context) {
	var userInfo model.UserInfo
	_ = c.ShouldBindJSON(&userInfo)
	fmt.Printf("%+v\n", userInfo)

	res, _ := weapp.Login(global.MINI.AppId, global.MINI.Secret, userInfo.Code)
	var query model.UserInfo
	notFound := global.DB.Where(model.UserInfo{OpenId: res.OpenID}).First(&query).RecordNotFound()
	if notFound {
		userInfo.OpenId = res.OpenID
		err := model.AddUserInfo(userInfo)
		if err != nil {
			fmt.Println(err)
			global.FailWithMessage(fmt.Sprintf("插入数据失败，%v", err), c)
		} else {
			data, _ := model.GetUserByOpenId(userInfo.OpenId)
			global.OkWithData(data, c)
		}
	} else {
		global.OkWithData(query, c)
	}
}
