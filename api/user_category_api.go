package api

import (
	"github.com/gin-gonic/gin"
	"gluten/middleware"
	"gluten/model"
	"gluten/service"
	"gluten/util"
)

func InitUserCategoryRouter(Router *gin.RouterGroup) {
	CategoryGroup := Router.Group("user_category").Use(middleware.Auth())
	CategoryGroup.PUT("", UpdateUserCategory)
	CategoryGroup.GET("/actions/get", SelectUserCategoryById)
}

func UpdateUserCategory(c *gin.Context) {
	var userCategory model.UserCategory
	if err := c.ShouldBindJSON(&userCategory); err != nil {
		util.IncorrectParameters(err.Error(), c)
	}
	if err := service.CreateOrUpdateUserCategory(userCategory); err != nil {
		util.DBUpdateFailed(err, c)
	} else {
		util.Ok(c)
	}
}

func SelectUserCategoryById(c *gin.Context) {
	query, err := service.SelectUserCategoryById(util.GetJwtId(c))
	var company []string
	var category []string
	var post []string
	if err == nil {
		company = query.Company
		category = query.Category
		post = query.Post
	}
	util.OkWithData(gin.H{
		"company":  company,
		"category": category,
		"post":     post,
	}, c)
}
