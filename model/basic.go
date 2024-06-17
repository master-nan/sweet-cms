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
	ID            int            `gorm:"primaryKey;type:int" json:"id"`
	GmtCreate     CustomTime     `gorm:"type:datetime;autoCreateTime" json:"gmt_create"`
	GmtCreateUser *int           `json:"gmt_create_user"`
	GmtModify     CustomTime     `gorm:"type:datetime;autoCreateTime;autoUpdateTime" json:"gmt_modify"`
	GmtModifyUser *int           `json:"gmt_modify_user"`
	GmtDelete     gorm.DeletedAt `gorm:"type:datetime;comment:删除时间" json:"-"`
	GmtDeleteUser *int           `json:"-"`
	State         bool           `gorm:"default:true" json:"state"`
}

func (b *Basic) BeforeCreate(tx *gorm.DB) (err error) {
	return
}

func (b *Basic) AfterFind(tx *gorm.DB) (err error) {
	return
}
