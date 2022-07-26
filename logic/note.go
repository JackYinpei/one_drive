package logic

import (
	"qqq_one_drive/controller"
	dao "qqq_one_drive/dao/mysql"

	"go.uber.org/zap"
)

type Note struct {
	Topic   string `form:"topic" json:"topic"`
	Article string `form:"article" json:"article" binding:"required"`
}

func (Note *Note) PostNote(noteID int64) controller.ResCode {
	note := dao.Note{
		Context: Note.Article,
		Tag:     Note.Topic,
		NoteID:  noteID,
	}

	if err := dao.DB.Create(&note).Error; err != nil {
		zap.L().Fatal("上传文档的时候出错拉", zap.Error(err))
		return controller.ResCode(500)
	}
	return controller.ResCode(200)
}
