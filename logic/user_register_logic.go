package logic

import (
	"qqq_one_drive/controller"
	dao "qqq_one_drive/dao/mysql"
	"qqq_one_drive/pkg/jwt"
	snowflake "qqq_one_drive/pkg/snowflake"

	"go.uber.org/zap"
)

type UserRegisterLogic struct {
	UserName        string `form:"user_name" json:"user_name" binding:"required,min=3,max=30"`
	PassWord        string `form:"pass_word" json:"pass_word" binding:"required,min=6,max=30"`
	PassWordConfirm string `form:"pass_word_confirm" json:"pass_word_confirm" binding:"required,min=6,max=30"`
	Email           string `form:"email" json:"email" binding:"required,email"`
	Token           string `form:"token" json:"token" binding:"omitempty"`
}

func (logic *UserRegisterLogic) valid() *controller.ResponseData {
	if logic.PassWord != logic.PassWordConfirm {
		return &controller.ResponseData{
			Code: 200,
			Msg:  "两次密码输入的不一样哦",
		}
	}
	count := int64(0)
	dao.DB.Model(&dao.User{}).Where("user_name = ?", logic.UserName).Count(&count)
	if count > 0 {
		return &controller.ResponseData{
			Code: 200,
			Msg:  "用户名被注册",
		}
	}
	return nil
}

func (logic *UserRegisterLogic) Register() controller.ResponseData {
	userID := snowflake.GenID()
	user := dao.User{
		UserID:   userID,
		UserName: logic.UserName,
		Email:    logic.Email,
	}
	if err := logic.valid(); err != nil {
		return *err
	}

	if err := user.SetPassword(logic.PassWord); err != nil {
		zap.L().Fatal("用户注册时，密码加密的时候出错了", zap.Error(err))
		return controller.ResponseData{
			Code: 200,
			Msg:  "encypto password error",
		}
	}

	if err := dao.DB.Create(&user).Error; err != nil {
		zap.L().Fatal("用户注册的时候出现错误", zap.Error(err))
		return controller.ResponseData{
			Code: 200,
			Msg:  "注册失败",
		}
	}
	token, err := jwt.GenToken(userID, logic.UserName)
	if err != nil {
		return controller.ResponseData{
			Code: 200,
			Msg:  "注册失败",
		}
	}
	user.Token = token
	return controller.ResponseData{
		Msg:  "haojiahuo fanhui chenggongle ya ",
		Code: 200,
		Data: user,
	}
}
