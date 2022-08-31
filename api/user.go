package api

import (
	"fmt"
	"net/http"
	"qqq_one_drive/controller"
	dao "qqq_one_drive/dao/mysql"
	"qqq_one_drive/logic"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func Register(c *gin.Context) {
	var userRegister logic.UserRegisterLogic
	if err := c.ShouldBind(&userRegister); err == nil {
		res := userRegister.Register()
		c.JSON(200, controller.ResponseData{
			Code: 200,
			Msg:  "成功",
			Data: res,
		})
	} else {
		c.JSON(200, controller.ResponseData{
			Code: 200,
			Msg:  "失败",
			Data: err,
		})
	}
}

func Login(c *gin.Context) {
	var userLogin logic.UserLoginService
	if err := c.ShouldBind(&userLogin); err == nil {
		res := userLogin.Login()
		daoUser := res.Data
		userInDao, OK := daoUser.(dao.User)
		fmt.Println(userInDao, "user in dao shigesha ")
		if OK {
			// TODO 我想让登录成功后重定向到一个新的界面
			c.HTML(http.StatusOK, "index.html", gin.H{
				"data": userInDao,
			})
		} else {
			zap.L().Error("登陆后返回的值有问题")
			c.JSON(200, gin.H{
				"MSG": "阿西吧",
			})
		}
	} else {
		c.JSON(200, controller.ResponseData{
			Code: 200,
			Msg:  "失败",
			Data: err,
		})
	}
}
