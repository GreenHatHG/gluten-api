package initialize

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gluten/api"
	"gluten/middleware"
	"net/http"
)

// 初始化总路由
func Routers() *gin.Engine {
	fmt.Println("router register begin")
	var Router = gin.Default()
	Router.Use(middleware.GinLoggerToFile())
	Router.Use(Cors())
	// 方便统一添加路由组前缀
	ApiGroup := Router.Group("")
	api.InitGlutenInfoRouter(ApiGroup)
	//api.InitUserInfoRouter(ApiGroup)
	api.InitConfigRouter(ApiGroup)
	api.InitOauthRouter(ApiGroup)
	api.InitUserCategoryRouter(ApiGroup)
	fmt.Println("router register success")
	return Router
}

// 处理跨域请求,支持options访问
func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Headers", "Content-Type,AccessToken,X-CSRF-Token, Authorization, Token")
		c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type")
		c.Header("Access-Control-Allow-Credentials", "true")

		// 放行所有OPTIONS方法
		if method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
		}
		// 处理请求
		c.Next()
	}
}
