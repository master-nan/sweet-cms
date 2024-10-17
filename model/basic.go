package model

import (
	"database/sql/driver"
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
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
	Id         int            `gorm:"primaryKey;comment:ID" json:"id"`
	GmtCreate  CustomTime     `gorm:"type:datetime;autoCreateTime;comment:创建时间" json:"gmtCreate"`
	CreateUser *int           `gorm:"comment:创建人ID" json:"createUser"`
	CreateName *string        `gorm:"size:128;comment:创建人" json:"createName"`
	GmtModify  CustomTime     `gorm:"type:datetime;autoCreateTime;autoUpdateTime;comment:修改时间" json:"gmtModify"`
	ModifyUser *int           `gorm:"column:modify_user;comment:修改人ID" json:"modifyUser"`
	ModifyName *string        `gorm:"size:128;comment:修改人" json:"modifyName"`
	GmtDelete  gorm.DeletedAt `gorm:"type:datetime;comment:删除时间" json:"-"`
	DeleteUser *int           `gorm:"column:delete_user;comment:删除人ID" json:"deleteUser"`
	DeleteName *string        `gorm:"size:128;comment:删除人" json:"-"`
	State      bool           `gorm:"default:true;comment:状态" json:"state"`
}

func (b *Basic) BeforeCreate(tx *gorm.DB) (err error) {
	var user SysUser
	ctx, ok := tx.Statement.Context.(*gin.Context)
	if ok {
		obj, exists := ctx.Get("user")
		if exists {
			user, _ = obj.(SysUser)
			tx.Statement.SetColumn("create_user", user.EmployeeId)
		}
	}
	return
}

func (b *Basic) BeforeUpdate(tx *gorm.DB) error {
	var user SysUser
	ctx, ok := tx.Statement.Context.(*gin.Context)
	if ok {
		obj, exists := ctx.Get("user")
		if exists {
			user, _ = obj.(SysUser)
			tx.Statement.SetColumn("modify_user", user.EmployeeId)
		}
	}
	return nil
}

func (b *Basic) BeforeDelete(tx *gorm.DB) error {
	var user SysUser
	ctx, ok := tx.Statement.Context.(*gin.Context)
	if ok {
		obj, exists := ctx.Get("user")
		if exists {
			user, _ = obj.(SysUser)

			tx.Statement.AddClause(clause.Update{})
			tx.Statement.AddClause(clause.Set{
				{Column: clause.Column{Name: "delete_user"}, Value: user.EmployeeId},
				{Column: clause.Column{Name: "gmt_delete"}, Value: time.Now()},
			})
			tx.Statement.Build(
				clause.Update{}.Name(),
				clause.Set{}.Name(),
				clause.Where{}.Name(),
			)
		}
	}
	return nil
}
