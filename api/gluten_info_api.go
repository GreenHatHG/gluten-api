package api

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"gluten/middleware"
	"gluten/model"
	"gluten/service"
	"gluten/util"
	"strings"
)

func InitGlutenInfoRouter(Router *gin.RouterGroup) {
	GlutenInfoGroup := Router.Group("gluten_info").Use(middleware.Auth())
	GlutenInfoGroup.POST("/actions/add", AddGlutenInfo)
	GlutenInfoGroup.GET("", SelectAllGlutenInfoById)
}

func SelectAllGlutenInfoById(c *gin.Context) {
	if err, data := service.SelectAllGlutenInfoById(util.GetJwtId(c)); err != nil {
		util.Logger.Error(err)
		util.FailWithMessage("获取数据失败", c)
	} else {
		util.OkWithData(data, c)
	}
}

func AddGlutenInfo(c *gin.Context) {
	var body struct {
		Content         string `binding:"required,min=1,max=30"`
		Category        string `binding:"required"`
		Company         string `binding:"required"`
		Post            string `binding:"required"`
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

	if _, err := service.CreateOrUpdateUserCategory(model.UserCategory{
		ID:       util.GetJwtId(c),
		Category: body.Category,
		Company:  body.Company,
		Post:     body.Post,
	}); err != nil {
		util.DBUpdateFailed(err, c)
	}

	companyJson := GetInitMap(body.ContentCompany)
	postJson := GetInitMap(body.ContentPost)
	if err := service.AddGlutenInfo(model.GlutenInfo{
		Title:    body.Content,
		Star:     1,
		Category: strings.Join(body.ContentCategory, ","),
		Company:  companyJson,
		Post:     postJson,
		UserId:   util.GetJwtId(c),
	}); err != nil {
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
