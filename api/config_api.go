package api

import (
	"github.com/gin-gonic/gin"
	"gluten/global"
)

func InitConfigRouter(Router *gin.RouterGroup) {
	GlutenInfoGroup := Router.Group("config")
	GlutenInfoGroup.GET("/github", GetGithubConfig)
}

func GetGithubConfig(c *gin.Context) {
	global.OkWithData(gin.H{
		"clientId": global.GITHUB.ClientID,
	}, c)
}
