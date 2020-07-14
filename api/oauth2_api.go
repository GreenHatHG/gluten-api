package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/tidwall/gjson"
	"gluten/global"
	"gluten/model"
	"gluten/util"
)

func InitOauth2Router(Router *gin.RouterGroup) {
	GlutenInfoGroup := Router.Group("oauth2")
	GlutenInfoGroup.POST("/github", GithubOauth2)
}

func GithubOauth2(c *gin.Context) {
	code := c.Query("code")

	//根据code获取token
	body := fmt.Sprintf(`{"client_id":"%s","client_secret":"%s","code":"%s"}`,
		global.GITHUB.ClientID, global.GITHUB.ClientSecret, code)
	content, err := util.Post("https://github.com/login/oauth/access_token", body)
	if err != nil {
		util.Logger.Error(err)
		util.FailWithMessage("请求github token失败", c)
	}
	token := gjson.Get(string(content), "access_token").String()

	//根据token获取用户信息
	content, err = util.Get("https://api.github.com/user", token)
	if err != nil {
		util.Logger.Error(err)
		util.FailWithMessage("获取用户信息失败", c)
	}
	data := gjson.Parse(string(content)).Map()

	//保存到数据库
	info := model.UserInfo{AvatarUrl: data["avatar_url"].Str, Username: data["login"].Str, Email: data["email"].Str,
		Location: data["location"].Str, OauthId: int(data["id"].Int())}
	userInfo := model.UserInfo.CreateOrUpdateUserInfo(info)
	if token, err := util.GetJWTToken(int(userInfo.ID)); err == nil {
		util.OkWithData(gin.H{
			"avatarUrl": userInfo.AvatarUrl,
			"username":  userInfo.Username,
			"id":        userInfo.ID,
			"token":     token,
		}, c)
	} else {
		util.Logger.Error(err)
		util.FailWithMessage("token获取失败", c)
	}
}
