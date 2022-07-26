package logic

import (
	"fmt"
	"qqq_one_drive/controller"
	dao "qqq_one_drive/dao/mysql"
	"qqq_one_drive/pkg/jwt"

	"go.uber.org/zap"
)

// UserLoginService 管理用户登录的服务
type UserLoginService struct {
	UserName string `form:"user_name" json:"user_name" binding:"required,min=5,max=30"`
	Password string `form:"pass_word" json:"pass_word" binding:"required,min=8,max=40"`
}

// Login 用户登录函数
func (service *UserLoginService) Login() controller.ResponseData {
	var user dao.User

	// find user_name
	if err := dao.DB.Where("user_name = ?", service.UserName).First(&user).Error; err != nil {
		return controller.ResponseData{
			Code: 200,
			Msg:  "账号或密码错误",
		}
	}
	fmt.Println(user, "user_login_logic")

	if !user.CheckPassword(service.Password) {
		return controller.ResponseData{
			Code: 200,
			Msg:  "账号或密码错误",
		}
	}

	token, err := jwt.GenToken(user.UserID, user.UserName)
	if err != nil {
		zap.L().Fatal("生成Token 错误", zap.Error(err))
	}
	user.Token = token
	return controller.ResponseData{
		Code: 200,
		Msg:  "登录成功",
		Data: user,
	}
}
