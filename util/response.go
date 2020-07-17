package util

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Response struct {
	Code int         `json:"code"`
	Data interface{} `json:"data"`
	Msg  string      `json:"msg"`
}

const (
	ERROR   = 7
	SUCCESS = 0
)

func Result(httpCode int, code int, data interface{}, msg string, c *gin.Context) {
	c.JSON(httpCode, Response{
		code,
		data,
		msg,
	})
}

func Ok(c *gin.Context) {
	Result(http.StatusOK, SUCCESS, map[string]interface{}{}, "操作成功", c)
}

func OkWithMessage(message string, c *gin.Context) {
	Result(http.StatusOK, SUCCESS, map[string]interface{}{}, message, c)
}

func OkWithData(data interface{}, c *gin.Context) {
	Result(http.StatusOK, SUCCESS, data, "操作成功", c)
}

func OkDetailed(data interface{}, message string, c *gin.Context) {
	Result(http.StatusOK, SUCCESS, data, message, c)
}

func Fail(c *gin.Context) {
	Result(http.StatusBadRequest, ERROR, map[string]interface{}{}, "操作失败", c)
}

func FailWithMessage(message string, c *gin.Context) {
	Result(http.StatusBadRequest, ERROR, map[string]interface{}{}, message, c)
}

func FailWithDetailed(code int, data interface{}, message string, c *gin.Context) {
	Result(http.StatusBadRequest, code, data, message, c)
}

func StatusUnauthorized(c *gin.Context) {
	Result(http.StatusUnauthorized, ERROR, map[string]interface{}{}, "token验证不通过", c)
}

func IncorrectParameters(err string, c *gin.Context) {
	Result(http.StatusBadRequest, 12121, map[string]interface{}{}, "参数错误："+err, c)
}

func DBUpdateFailed(err error, c *gin.Context) {
	Result(http.StatusBadRequest, 13334, map[string]interface{}{}, "数据库更新失败："+err.Error(), c)
}
