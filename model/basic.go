package model

import (
	"go.uber.org/zap"
	"gorm.io/gorm"
	"sweet-cms/global"
	"time"
)

type Basic struct {
	ID            int            `gorm:"primaryKey;type:int" json:"id"`
	GmtCreate     time.Time      `gorm:"autoCreateTime" json:"gmt_create"`
	GmtCreateUser int            `json:"gmt_create_user"`
	GmtModify     time.Time      `gorm:"autoCreateTime;autoUpdateTime" json:"gmt_modify"`
	GmtModifyUser int            `json:"gmt_modify_user"`
	GmtDelete     gorm.DeletedAt `gorm:"type:time;comment:删除时间" json:"-"`
	GmtDeleteUser int            `json:"gmt_delete_user"`
	State         bool           `gorm:"default:true" json:"state"`
}

func (b *Basic) BeforeCreate(tx *gorm.DB) (err error) {
	uniqueID, err := global.SF.GenerateUniqueID()
	if err != nil {
		zap.S().Errorf("获取id失败：%s", err)
		return err
	}
	b.ID = int(uniqueID)
	return
}
