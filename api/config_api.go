package api

import (
	"github.com/gin-gonic/gin"
	"gluten/global"
	"gluten/util"
)

func InitConfigRouter(Router *gin.RouterGroup) {
	ConfigGroup := Router.Group("config")
	ConfigGroup.GET("/github", GetGithubConfig)
}

func GetGithubConfig(c *gin.Context) {
	util.OkWithData(gin.H{
		"clientId": global.GITHUB.ClientID,
	}, c)
}
