package server

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"strconv"
	"sweet-cms/form/request"
	"sweet-cms/global"
	"sweet-cms/model"
	"time"
)

type ArticleServer struct {
	ctx *gin.Context
}

func NewArticleServer(ctx *gin.Context) *ArticleServer {
	return &ArticleServer{
		ctx: ctx,
	}
}

func (as *ArticleServer) GetArticleBasic(query request.ArticleBasicQueryReq) ([]model.ArticleBasic, int, error) {
	var ab []model.ArticleBasic
	var count int64 = 0
	db := global.DB.Where("gmt_delete is null")
	if query.Title != nil {
		db = db.Where("title like ?", "%"+*query.Title+"%")
	}
	if query.Introduction != nil {
		db = db.Where("introduction like ?", "%"+*query.Introduction+"%")
	}
	if query.IsAd != nil {
		db = db.Where("is_ad = ?", query.IsAd)
	}
	if query.Status != nil {
		if *query.Status != "" {
			db = db.Where("status = ?", query.Status)
		}
	}
	if query.IsComment != nil {
		db = db.Where("is_comment = ?", query.IsComment)
	}
	if query.ChannelId != nil {
		db = db.Where("channel_id = ?", query.ChannelId)
	}
	if query.Page > 0 && query.Num > 0 {
		db = db.Limit(query.Page).Offset(query.Page * (query.Num - 1))
	}
	err := db.Find(&ab).Limit(-1).Offset(-1).Count(&count).Error
	if err != nil {
		return []model.ArticleBasic{}, 0, err
	}
	return ab, int(count), nil
}

func (as *ArticleServer) GetArticleBasicByID(uid string) (model.ArticleBasic, error) {
	var ab model.ArticleBasic
	err := global.DB.Where("uuid = ? and gmt_delete is null", uid).First(&ab).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return model.ArticleBasic{}, err
	}
	return ab, nil
}

func (as *ArticleServer) CreateArticleBasic(ab model.ArticleBasic) (string, error) {
	ab.UUID = uuid.New()
	err := global.DB.Omit("gmt_delete").Create(&ab).Error
	if err != nil {
		return "", err
	}
	return ab.UUID.String(), nil
}

func (as *ArticleServer) DeleteArticleBasic(uid string) error {
	err := global.DB.Model(model.ArticleBasic{}).Where("uuid = ?", uid).Updates(map[string]interface{}{"gmt_delete": time.Now(), "state": false}).Error
	if err != nil {
		return err
	}
	return nil
}

func (as *ArticleServer) UpdateArticleBasic(uid uuid.UUID, data request.ArticleBasicUpdateReq) error {
	str := uid.String()
	db := global.DB.Model(model.ArticleBasic{})
	err := db.Where("uuid = ?", str).Updates(data).Error
	if err != nil {
		return err
	}
	return nil
}

func (as *ArticleServer) CreateArticleChannel(ac model.ArticleChannel) (int, error) {
	err := global.DB.Omit("gmt_delete").Create(&ac).Error
	if err != nil {
		return 0, err
	}
	return ac.ID, nil
}

func (as *ArticleServer) DeleteArticleChannel(id int) error {
	var ac = model.ArticleChannel{
		Basic: model.Basic{
			ID: id,
		},
	}
	err := global.DB.Model(&ac).Updates(map[string]interface{}{"gmt_delete": time.Now(), "state": false}).Error
	if err != nil {
		return err
	}
	return nil
}

func (as *ArticleServer) UpdateArticleChannel(data request.ArticleChannelUpdateReq) error {
	id, err := strconv.Atoi(as.ctx.Param("id"))
	if err != nil {
		return err
	}
	db := global.DB.Model(model.ArticleChannel{})
	err = db.Where("id = ?", id).Updates(data).Error
	if err != nil {
		return err
	}
	return nil
}

func (as *ArticleServer) GetArticleChannelByID(id int) (model.ArticleChannel, error) {
	var acModel model.ArticleChannel
	err := global.DB.Where("gmt_delete is null").First(&acModel, id).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return acModel, err
	}
	return acModel, nil
}

func (as *ArticleServer) GetArticleChannel(query request.ArticleChannelQueryReq) ([]model.ArticleChannel, int, error) {
	var ac []model.ArticleChannel
	var count int64
	db := global.DB.Where("gmt_delete is null")
	if query.Name != nil {
		db = db.Where("name like ?", "%"+*query.Name+"%")
	}
	if query.PID != nil {
		db = db.Where("pid = ?", query.PID)
	}
	if query.IsVisible != nil {
		db = db.Where("is_visible = ?", query.IsVisible)
	}
	if query.Page > 0 && query.Num > 0 {
		db = db.Limit(query.Page).Offset(query.Page * (query.Num - 1))
	}
	err := db.Find(&ac).Limit(-1).Offset(-1).Count(&count).Error
	if err != nil {
		return nil, 0, err
	}
	return ac, int(count), nil
}

func (as *ArticleServer) GetArticleContentByID(id int) (model.ArticleContent, error) {
	var acModel model.ArticleContent
	err := global.DB.Where("gmt_delete is null").First(&acModel, id).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return acModel, err
	}
	return acModel, nil
}

func (as *ArticleServer) CreateArticleContent(ac model.ArticleContent) (int, error) {
	err := global.DB.Omit("gmt_delete").Create(&ac).Error
	if err != nil {
		return 0, err
	}
	return ac.ID, nil
}

func (as *ArticleServer) DeleteArticleContent(id int) error {
	var ac = model.ArticleContent{
		Basic: model.Basic{
			ID: id,
		},
	}
	err := global.DB.Model(&ac).Updates(map[string]interface{}{"gmt_delete": time.Now(), "state": false}).Error
	if err != nil {
		return err
	}
	return nil
}

func (as *ArticleServer) UpdateArticleContent(id uuid.UUID, data request.ArticleContentUpdateReq) error {
	db := global.DB.Model(model.ArticleContent{})
	err := db.Where("uuid = ?", id.String()).Updates(data).Error
	if err != nil {
		return err
	}
	return nil
}
