package request

import (
	"github.com/google/uuid"
	"time"
)

type QuestionnaireContentQueryReq struct {
	Page     int        `form:"page"`
	Num      int        `form:"num"`
	Title    *string    `form:"title"`
	Status   *string    `form:"status"`
	Creator  *int       `form:"creator"`
	GmtStart *time.Time `form:"gmt_start"`
	GmtEnd   *time.Time `form:"gmt_end"`
}

type QuestionnaireContentCreateReq struct {
	Title    string     `json:"title" binding:"required"`
	Status   string     `json:"status"`
	Creator  int        `json:"creator"`
	GmtStart *time.Time `json:"gmtStart"`
	GmtEnd   *time.Time `json:"gmtEnd"`
}

type QuestionnaireContentUpdateReq struct {
	Title    string `json:"title" `
	Status   string `json:"status"`
	Content  string `json:"content"`
	Creator  int    `json:"creator"`
	GmtStart string `json:"gmtStart"`
	GmtEnd   string `json:"gmtEnd"`
}

type QuestionnaireAnswerQueryReq struct {
	Page         int     `form:"page"`
	Num          int     `form:"num"`
	QuestionUUID *string `form:"questionUUID"`
	Creator      *string `form:"creator"`
}

type QuestionnaireAnswerCreateReq struct {
	QuestionUUID uuid.UUID `json:"questionUUID"`
	Creator      int       `json:"creator"`
	Content      string    `json:"content"`
}

type QuestionnaireAnswerUpdateReq struct {
	QuestionUUID string `json:"questionUUID"`
	Creator      int    `json:"creator"`
	Content      string `json:"content"`
}
