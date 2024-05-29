package model

import (
	"encoding/json"
	"gorm.io/gorm"
	"time"
)

type Basic struct {
	ID            int            `gorm:"primaryKey;type:int" json:"id"`
	GmtCreate     time.Time      `gorm:"type:datetime;autoCreateTime" json:"gmt_create"`
	GmtCreateUser int            `json:"gmt_create_user"`
	GmtModify     time.Time      `gorm:"type:datetime;autoCreateTime;autoUpdateTime" json:"gmt_modify"`
	GmtModifyUser int            `json:"gmt_modify_user"`
	GmtDelete     gorm.DeletedAt `gorm:"type:datetime;comment:删除时间" json:"-"`
	GmtDeleteUser int            `json:"gmt_delete_user"`
	State         bool           `gorm:"default:true" json:"state"`
}

func (b *Basic) BeforeCreate(tx *gorm.DB) (err error) {
	//uniqueID, err := global.SF.GenerateUniqueID()
	//if err != nil {
	//	zap.S().Errorf("获取id失败：%s", err)
	//	return err
	//}
	//b.ID = int(uniqueID)
	return
}

func (b Basic) MarshalJSON() ([]byte, error) {
	type Alias Basic
	return json.Marshal(&struct {
		GmtCreate string `json:"gmt_create"`
		GmtModify string `json:"gmt_modify"`
		Alias
	}{
		GmtCreate: b.GmtCreate.Format(time.DateTime),
		GmtModify: b.GmtModify.Format(time.DateTime),
		Alias:     (Alias)(b),
	})
}

func (b *Basic) UnmarshalJSON(data []byte) error {
	type Alias Basic
	aux := &struct {
		GmtCreate string `json:"gmt_create"`
		GmtModify string `json:"gmt_modify"`
		*Alias
	}{
		Alias: (*Alias)(b),
	}
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}
	var err error
	b.GmtCreate, err = time.Parse(time.DateTime, aux.GmtCreate)
	if err != nil {
		return err
	}
	b.GmtModify, err = time.Parse(time.DateTime, aux.GmtModify)
	if err != nil {
		return err
	}
	return nil
}

func (b *Basic) AfterFind(tx *gorm.DB) (err error) {
	return
}
