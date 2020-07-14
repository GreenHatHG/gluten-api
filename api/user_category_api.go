package api

import (
	"github.com/gin-gonic/gin"
	"gluten/model"
	"gluten/util"
	"strings"
)

func InitUserCategoryRouter(Router *gin.RouterGroup) {
	GlutenInfoGroup := Router.Group("user_category")
	GlutenInfoGroup.PUT("", UpdateUserCategory)
	GlutenInfoGroup.GET("/actions/get", SelectUserCategoryById)
}

func UpdateUserCategory(c *gin.Context) {
	var userCategory model.UserCategory
	if err := c.ShouldBindJSON(&userCategory); err != nil {
		util.FailWithDetailed(12121, err.Error(), "参数错误", c)
	}
	query := userCategory.CreateOrUpdateUserCategory()
	util.OkWithData(query, c)
}

func SelectUserCategoryById(c *gin.Context) {
	id, _ := c.Get("id")
	category := model.SelectUserCategoryById(id.(uint))
	util.OkWithData(gin.H{
		"id":       category.ID,
		"company":  strings.Split(category.Company, "/"),
		"category": strings.Split(category.Category, "/"),
		"post":     strings.Split(category.Post, "/"),
	}, c)
}
