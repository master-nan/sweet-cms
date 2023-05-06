package model

import (
	"time"
)

type BasicModel struct {
	ID        int       `gorm:"primaryKey;type:int" json:"id"`
	GmtCreate time.Time `gorm:"autoCreateTime" json:"gmtCreate"`
	GmtModify time.Time `gorm:"autoCreateTime;autoUpdateTime" json:"gmtModify"`
	GmtDelete time.Time `json:"-"`
	State     bool      `gorm:"default:true" json:"-"`
}
