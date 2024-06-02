package model

import (
	"github.com/google/uuid"
	"reflect"
	"sweet-cms/enum"
	"time"
)

type ArticleBasic struct {
	Basic
	UUID         uuid.UUID        `gorm:"type:uuid;column:uuid" json:"uuid"`
	Title        string           `json:"title"`
	Cover        string           `json:"cover"`
	Introduction string           `gorm:"default:null" json:"introduction"`
	IsAd         bool             `gorm:"default:false" json:"isAd"`
	Status       enum.ArticleType `gorm:"default:draft" json:"status"`
	GmtReview    time.Time        `gorm:"default:null" json:"gmtReview"`
	RejectReason string           `gorm:"default:null" json:"rejectReason"`
	IsComment    bool             `gorm:"default:false" json:"isComment"`
	ChannelId    int              `gorm:"type:int" json:"channelId"`
}

func (ab ArticleBasic) IsEmpty() bool {
	return reflect.DeepEqual(ab, ArticleBasic{})
}

type ArticleChannel struct {
	Basic
	PID       int    `gorm:"column:pid;type:int" json:"pid"`
	Name      string `json:"name"`
	Sequence  uint8  `json:"sequence"`
	IsVisible bool   `json:"isVisible"`
}

type ArticleComment struct {
	Basic
	ArticleUUID uuid.UUID `gorm:"column:article_uuid;type:uuid" json:"articleUUID"`
	PID         int       `json:"pid"`
	Comment     string    `json:"comment"`
}

type ArticleContent struct {
	Basic
	ArticleUUID string `json:"articleUUID"`
	Content     string `json:"content"`
}

type ArticleRead struct {
	Basic
	ArticleUUID string `json:"articleUUID"`
	IP          string `json:"ip"`
}
