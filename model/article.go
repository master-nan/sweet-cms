package model

import (
	"database/sql/driver"
	"github.com/google/uuid"
	"reflect"
	"time"
)

type ArticleType string

const (
	DRAFT   ArticleType = "draft"
	REVIEW  ArticleType = "review"
	RELEASE ArticleType = "release"
	REJECT  ArticleType = "reject"
)

func (at ArticleType) Value() (driver.Value, error) {
	return string(at), nil
}

type ArticleBasic struct {
	BasicModel
	UUID         uuid.UUID   `gorm:"type:uuid;column:uuid" json:"uuid"`
	Title        string      `json:"title"`
	Cover        string      `json:"cover"`
	Introduction string      `gorm:"default:null" json:"introduction"`
	IsAd         bool        `gorm:"default:false" json:"isAd"`
	Status       ArticleType `gorm:"default:draft" json:"status"`
	GmtReview    time.Time   `gorm:"default:null" json:"gmtReview"`
	RejectReason string      `gorm:"default:null" json:"rejectReason"`
	IsComment    bool        `gorm:"default:false" json:"isComment"`
	ChannelId    int         `gorm:"type:int" json:"channelId"`
}

func (ab ArticleBasic) IsEmpty() bool {
	return reflect.DeepEqual(ab, ArticleBasic{})
}

type ArticleChannel struct {
	BasicModel
	PID       int    `gorm:"column:pid;type:int" json:"pid"`
	Name      string `json:"name"`
	Sequence  uint8  `json:"sequence"`
	IsVisible bool   `json:"isVisible"`
}

type ArticleComment struct {
	BasicModel
	ArticleUUID uuid.UUID `gorm:"column:article_uuid;type:uuid" json:"articleUUID"`
	PID         int       `json:"pid"`
	Comment     string    `json:"comment"`
}

type ArticleContent struct {
	BasicModel
	ArticleUUID string `json:"articleUUID"`
	Content     string `json:"content"`
}

type ArticleRead struct {
	BasicModel
	ArticleUUID string `json:"articleUUID"`
	IP          string `json:"ip"`
}
