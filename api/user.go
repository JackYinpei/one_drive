package api

import (
	"qqq_one_drive/controller"
	"qqq_one_drive/logic"

	"github.com/gin-gonic/gin"
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
		c.JSON(200, controller.ResponseData{
			Code: 200,
			Msg:  "登录成功",
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
