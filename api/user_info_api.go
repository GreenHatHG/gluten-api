package api

//
import (
	"github.com/gin-gonic/gin"
)

func InitUserInfoRouter(Router *gin.RouterGroup) {
	GlutenInfoGroup := Router.Group("user_info")
	GlutenInfoGroup.POST("/action/login", AddUserInfo)
}

func AddUserInfo(c *gin.Context) {
	//var userInfo model.UserInfo
	//_ = c.ShouldBindJSON(&userInfo)
	//fmt.Printf("%+v\n", userInfo)
	//
	//res, err := weapp.Login(global.MINI.AppId, global.MINI.Secret, userInfo.Code)
	//if err == nil {
	//	var query model.UserInfo
	//	global.DB.Where(model.UserInfo{OpenId: res.OpenID}).Assign(userInfo).FirstOrCreate(&query)
	//	global.OkWithData(gin.H{
	//		"id":        query.ID,
	//		"avatarUrl": query.AvatarUrl,
	//		"nickName":  query.NickName,
	//	}, c)
	//} else {
	//	fmt.Println(err)
	//	global.FailWithMessage(err.Error(), c)
	//}

}
