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

type GetNote struct {
	Bottom int32 `form:"bottom" json:"bottom"`
	Top    int32 `from:"top" json:"top"`
}

func (GetNote *GetNote) GetNoteLogic() controller.ResponseData {
	notes := make([]dao.Note, 20)
	if err := dao.DB.Order("updated_at").Find(&notes).Error; err != nil {
		zap.L().Error("根据updated时间查询note 失败了")
		return controller.ResponseData{
			Code: 200,
			Msg:  "查询失败",
			Data: err,
		}
	}
	return controller.ResponseData{
		Code: 200,
		Msg:  "OK",
		Data: notes,
	}
}
