package api

import (
	"github.com/gin-gonic/gin"
	"gluten/global"
	"gluten/util"
)

func InitConfigRouter(Router *gin.RouterGroup) {
	GlutenInfoGroup := Router.Group("config")
	GlutenInfoGroup.GET("/github", GetGithubConfig)
}

func GetGithubConfig(c *gin.Context) {
	util.OkWithData(gin.H{
		"clientId": global.GITHUB.ClientID,
	}, c)
}
