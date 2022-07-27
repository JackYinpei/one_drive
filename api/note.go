package api

import (
	"qqq_one_drive/controller"
	"qqq_one_drive/logic"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func PostNote(c *gin.Context) {
	var postNote logic.Note
	// TODO 给postNote 加上用户信息
	noteID, yes := c.Get(controller.CtxUserIDKey)
	if !yes {
		zap.L().Fatal("Something wrong, cannot get userID after JWT")
	}
	if err := c.ShouldBind(&postNote); err == nil {
		postNote.PostNote(noteID.(int64))
	} else {
		zap.L().Error("绑定前端信息失败，错误信息是", zap.Error(err))
		c.JSON(200, controller.ResponseData{
			Code: 200,
			Msg:  "上传Note失败",
		})
	}
}

func GetNote(c *gin.Context) {
	var getNote logic.GetNote
	resData := getNote.GetNoteLogic()
	c.JSON(200, resData.Data)
}
