package api

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"gluten/middleware"
	"gluten/model"
	"gluten/service"
	"gluten/util"
	"gorm.io/datatypes"
	"strconv"
	"strings"
	"time"
)

func InitGlutenInfoRouter(Router *gin.RouterGroup) {
	GlutenInfoGroup := Router.Group("gluten_info").Use(middleware.Auth())
	GlutenInfoGroup.POST("/actions/add", AddGlutenInfo)
	GlutenInfoGroup.GET("", SelectAllGlutenInfoByIdOrCategory)
	GlutenInfoGroup.GET("/count", CountGlutenInfo)
	GlutenInfoGroup.PUT("/title", UpdateGlutenTitle)
	GlutenInfoGroup.PUT("/value", UpdateGlutenValue)
	GlutenInfoGroup.PUT("/star", UpdateGlutenStar)
	GlutenInfoGroup.DELETE("/id", DeleteGlutenById)
}

func SelectAllGlutenInfoByIdOrCategory(c *gin.Context) {
	currentPage := c.Query("current_page")
	pageSize := c.Query("page_size")
	sort := c.Query("sort")
	category := c.Query("category")

	pageSizeInt, pageSizeErr := strconv.ParseInt(pageSize, 10, 64)
	if pageSizeErr != nil {
		util.IncorrectParameters(pageSizeErr.Error(), c)
		return
	}
	currentPageInt, currentPageErr := strconv.ParseInt(currentPage, 10, 64)
	if currentPageErr != nil {
		util.IncorrectParameters(currentPageErr.Error(), c)
		return
	}
	var data []model.GlutenInfo
	var err error
	if category != "" {
		err, data = service.SelectGlutenInfoByCategory(util.GetJwtId(c), currentPageInt, pageSizeInt, sort, category)
	} else {
		err, data = service.SelectAllGlutenInfoById(util.GetJwtId(c), currentPageInt, pageSizeInt, sort)
	}
	if err != nil {
		util.Logger.Error(err)
		util.FailWithMessage("获取数据失败", c)
	} else {
		util.OkWithData(data, c)
	}
}

func CountGlutenInfo(c *gin.Context) {
	category := c.Query("category")
	if count, err := service.CountGlutenInfo(util.GetJwtId(c), category); err != nil {
		util.Logger.Error(err)
		util.FailWithMessage("获取数据失败", c)
	} else {
		util.OkWithData(count, c)
	}
}

func AddGlutenInfo(c *gin.Context) {
	var body struct {
		Content         string `binding:"required"`
		Category        []string
		Company         []string
		Post            []string
		ContentCategory []string
		ContentPost     []string
		ContentCompany  []string
	}
	if err := c.ShouldBindJSON(&body); err != nil {
		util.IncorrectParameters(err.Error(), c)
		return
	}

	if strings.Trim(body.Content, " ") == "" {
		util.IncorrectParameters("content不能全为空格", c)
		return
	}

	if err := service.CreateOrUpdateUserCategory(model.UserCategory{
		UserId:   util.GetJwtId(c),
		Category: body.Category,
		Company:  body.Company,
		Post:     body.Post,
	}); err != nil {
		util.Logger.Error(err)
		util.DBUpdateFailed(err, c)
	}

	if err := service.AddGlutenInfo(model.GlutenInfo{
		CreatedAt: time.Now().Local(),
		Title:     body.Content,
		Star:      1,
		Category:  body.ContentCategory,
		Company:   GetInitMap(body.ContentCompany),
		Post:      GetInitMap(body.ContentPost),
		UserId:    util.GetJwtId(c),
	}); err != nil {
		util.Logger.Error(err)
		util.DBUpdateFailed(err, c)
	}
}

//根据string数组生成kv json
func GetInitMap(data []string) []byte {
	var dataMap []map[string]interface{}
	for _, item := range data {
		t := make(map[string]interface{})
		t["title"] = item
		t["value"] = 1
		dataMap = append(dataMap, t)
	}
	mar, _ := json.Marshal(dataMap)
	return mar
}

func UpdateGlutenTitle(c *gin.Context) {
	var body struct {
		Id    string `binding:"required"`
		Title string `binding:"required"`
	}
	if err := c.ShouldBindJSON(&body); err != nil {
		util.IncorrectParameters(err.Error(), c)
		return
	}
	if err := service.UpdateGlutenInfoTitle(body.Id, body.Title, util.GetJwtId(c)); err != nil {
		util.Logger.Error(err)
		util.DBUpdateFailed(err, c)
	}
}

func UpdateGlutenValue(c *gin.Context) {
	var body struct {
		Id    string         `binding:"required"`
		Key   string         `binding:"required"`
		Value datatypes.JSON `binding:"required"`
	}
	if err := c.ShouldBindJSON(&body); err != nil {
		util.IncorrectParameters(err.Error(), c)
		return
	}
	if err := service.UpdateGlutenInfoValue(body.Id, util.GetJwtId(c), body.Key, body.Value); err != nil {
		util.Logger.Error(err)
		util.DBUpdateFailed(err, c)
	}
}

func UpdateGlutenStar(c *gin.Context) {
	var body struct {
		Id   string `binding:"required"`
		Star int    `binding:"required"`
	}
	if err := c.ShouldBindJSON(&body); err != nil {
		util.IncorrectParameters(err.Error(), c)
		return
	}
	if err := service.UpdateGlutenInfoStar(body.Id, util.GetJwtId(c), body.Star); err != nil {
		util.Logger.Error(err)
		util.DBUpdateFailed(err, c)
	}
}

func DeleteGlutenById(c *gin.Context) {
	var body struct {
		Id string `binding:"required"`
	}
	if err := c.ShouldBindJSON(&body); err != nil {
		util.IncorrectParameters(err.Error(), c)
		return
	}
	if err := service.DeleteGlutenById(body.Id, util.GetJwtId(c)); err != nil {
		util.Logger.Error(err)
		util.DBUpdateFailed(err, c)
	}
}
