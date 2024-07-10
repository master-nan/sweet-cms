package model

import (
	"database/sql/driver"
	"fmt"
	"gorm.io/gorm"
	"time"
)

type CustomTime time.Time

func (t *CustomTime) UnmarshalJSON(data []byte) (err error) {
	now, err := time.ParseInLocation(`"`+time.DateTime+`"`, string(data), time.Local)
	*t = CustomTime(now)
	return
}
func (t CustomTime) MarshalJSON() ([]byte, error) {
	b := make([]byte, 0, len(time.DateTime)+2)
	b = append(b, '"')
	b = time.Time(t).AppendFormat(b, time.DateTime)
	b = append(b, '"')
	return b, nil
}
func (t CustomTime) String() string {
	return time.Time(t).Format(time.DateTime)
}

func (t CustomTime) Value() (driver.Value, error) {
	return time.Time(t), nil
}

func (t *CustomTime) Scan(value interface{}) error {
	if value == nil {
		*t = CustomTime(time.Time{})
		return nil
	}
	switch v := value.(type) {
	case time.Time:
		*t = CustomTime(v)
		return nil
	case []byte:
		parsedTime, err := time.Parse(time.DateTime, string(v))
		if err != nil {
			return err
		}
		*t = CustomTime(parsedTime)
		return nil
	case string:
		parsedTime, err := time.Parse(time.DateTime, v)
		if err != nil {
			return err
		}
		*t = CustomTime(parsedTime)
		return nil
	default:
		return fmt.Errorf("unsupported scan type for CustomTime: %T", value)
	}
}

type Basic struct {
	Id             int            `gorm:"primaryKey;comment:ID" json:"id"`
	GmtCreate      CustomTime     `gorm:"type:datetime;autoCreateTime;comment:创建时间" json:"gmtCreate"`
	GmtCreateUser  *int           `gorm:"comment:创建人ID" json:"gmtCreateUser"`
	CreateUserName *string        `gorm:"size:128;comment:创建人" json:"CreateUserName"`
	GmtModify      CustomTime     `gorm:"type:datetime;autoCreateTime;autoUpdateTime;comment:修改时间" json:"gmtModify"`
	GmtModifyUser  *int           `gorm:"comment:修改人ID" json:"gmtModifyUser"`
	ModifyUserName *string        `gorm:"size:128;comment:修改人" json:"modifyUserName"`
	GmtDelete      gorm.DeletedAt `gorm:"type:datetime;comment:删除时间" json:"-"`
	GmtDeleteUser  *int           `gorm:"comment:删除人ID" json:"-"`
	DeleteUserName *string        `gorm:"size:128;comment:删除人" json:"deleteUserName"`
	State          bool           `gorm:"default:true;comment:状态" json:"state"`
}

func (b *Basic) BeforeCreate(tx *gorm.DB) (err error) {
	return
}

func (b *Basic) AfterFind(tx *gorm.DB) (err error) {
	return
}
