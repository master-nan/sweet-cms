/**
 * @Author: Nan
 * @Date: 2023/3/29 23:28
 */

package server

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"sweet-cms/form/request"
	"sweet-cms/global"
	"sweet-cms/model"
	"time"
)

type QuestionnaireServer struct {
	ctx *gin.Context
}

func NewQuestionnaireServer(ctx *gin.Context) *QuestionnaireServer {
	return &QuestionnaireServer{ctx: ctx}
}

func (q *QuestionnaireServer) GetQuestionnaireContent(query request.QuestionnaireContentQueryReq) ([]model.QuestionnaireContent, int, error) {
	var qc []model.QuestionnaireContent
	var count int64 = 0
	db := global.DB.Where("gmt_delete is null")
	if query.Title != nil {
		db = db.Where("title like ?", "%"+*query.Title+"%")
	}
	if query.Status != nil {
		if *query.Status != "" {
			db = db.Where("status = ?", query.Status)
		}
	}
	if query.Creator != nil {
		db = db.Where("creator = ?", query.Creator)
	}
	if query.Page > 0 && query.Num > 0 {
		db = db.Limit(query.Page).Offset(query.Page * (query.Num - 1))
	}
	err := db.Find(&qc).Limit(-1).Offset(-1).Count(&count).Error
	if err != nil {
		return []model.QuestionnaireContent{}, 0, err
	}
	return qc, int(count), nil
}

func (q *QuestionnaireServer) GetQuestionnaireContentById(id string) (model.QuestionnaireContent, error) {
	var qc model.QuestionnaireContent
	err := global.DB.Where("gmt_delete is null").First(&qc, "uuid = ?", id).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return qc, err
	}
	return qc, nil
}

func (q *QuestionnaireServer) CreateQuestionnaireContent(qc model.QuestionnaireContent) (int, error) {
	err := global.DB.Omit("gmt_delete").Create(&qc).Error
	if err != nil {
		return 0, err
	}
	return qc.ID, nil
}

func (q *QuestionnaireServer) UpdateQuestionnaireContent(id uuid.UUID, data request.QuestionnaireContentUpdateReq) error {
	db := global.DB.Model(model.QuestionnaireContent{})
	err := db.Where("uuid = ?", id.String()).Updates(data).Error
	if err != nil {
		return err
	}
	return nil
}

func (q *QuestionnaireServer) DeleteQuestionnaireContent(id string) error {
	err := global.DB.Model(model.QuestionnaireContent{}).Where("uuid = ?", id).Updates(map[string]interface{}{"gmt_delete": time.Now(), "state": false}).Error
	if err != nil {
		return err
	}
	return nil
}

func (q *QuestionnaireServer) GetQuestionnaireAnswer(query request.QuestionnaireAnswerQueryReq) ([]model.QuestionnaireAnswer, int, error) {
	var qa []model.QuestionnaireAnswer
	var count int64 = 0
	db := global.DB.Where("gmt_delete is null")

	if query.QuestionUUID != nil {
		db = db.Where("question_uuid =?", query.QuestionUUID)
	}
	if query.Creator != nil {
		db = db.Where("creator = ?", query.Creator)
	}
	if query.Page > 0 && query.Num > 0 {
		db = db.Limit(query.Page).Offset(query.Page * (query.Num - 1))
	}
	err := db.Find(&qa).Limit(-1).Offset(-1).Count(&count).Error
	if err != nil {
		return []model.QuestionnaireAnswer{}, 0, err
	}
	return qa, int(count), nil
}

func (q *QuestionnaireServer) GetQuestionnaireAnswerById(id int) ([]model.QuestionnaireAnswer, error) {
	var qa []model.QuestionnaireAnswer
	err := global.DB.Where("gmt_delete is null").Find(&qa, "question_uuid = ?", id).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return qa, err
	}
	return qa, nil
}

func (q *QuestionnaireServer) CreateQuestionnaireAnswer(qa model.QuestionnaireAnswer) (int, error) {
	err := global.DB.Omit("gmt_delete").Create(&qa).Error
	if err != nil {
		return 0, err
	}
	return qa.ID, nil
}

func (q *QuestionnaireServer) UpdateQuestionnaireAnswer(id int, data request.QuestionnaireAnswerUpdateReq) error {
	db := global.DB.Model(model.ArticleContent{})
	err := db.Where("id = ?", id).Updates(data).Error
	if err != nil {
		return err
	}
	return nil
}

func (q *QuestionnaireServer) DeleteQuestionnaireAnswer(id int) error {
	var qa = model.QuestionnaireAnswer{
		Basic: model.Basic{
			ID: id,
		},
	}
	err := global.DB.Model(&qa).Updates(map[string]interface{}{"gmt_delete": time.Now(), "state": false}).Error
	if err != nil {
		return err
	}
	return nil
}
