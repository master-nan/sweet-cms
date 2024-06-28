/**
 * @Author: Nan
 * @Date: 2023/8/24 21:42
 */

package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"log"
	"os"
	"sweet-cms/model"
	"time"
)

type Product struct {
	ID        uint           `gorm:"primarykey"`
	CreatedAt time.Time      `gorm:"column:gmt_create"`
	UpdatedAt time.Time      `gorm:"column:gmt_update"`
	DeletedAt gorm.DeletedAt `gorm:"index;column:gmt_delete"`
	Code      string
	Price     uint
}

func main() {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", "root", "123456", "127.0.0.1", 3306, "sweet-cms")
	dbLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold:             time.Second,
			Colorful:                  false,
			IgnoreRecordNotFoundError: true,
			LogLevel:                  logger.Info,
		},
	)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   "cms_",
			SingularTable: true,
		},
		Logger: dbLogger,
	})
	if err != nil {
		panic(err)
	}
	sqlDB, err := db.DB()
	if err != nil {
		panic(err)
	}
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour)

	//db.Migrator().DropTable(&model.SysTable{}, &model.SysTableField{}, &model.SysDict{}, &model.SysDictItem{}, &model.SysConfigure{},
	//	&model.AccessLog{}, &model.LoginLog{})
	//// 迁移 schema
	//db.AutoMigrate(&model.SysTable{}, &model.SysTableField{}, &model.SysDict{}, &model.SysDictItem{}, &model.SysConfigure{}, &model.AccessLog{}, &model.LoginLog{})
	db.Migrator().DropTable(&model.SysUser{})
	db.AutoMigrate(&model.SysUser{})

	// Create
	//m := &model.SysConfigure{EnableCaptcha: false}
	//sf, err := utils.NewSnowflake(1)
	//if err != nil {
	//	panic(err)
	//}
	//uniqueID, err := sf.GenerateUniqueID()
	//m.ID = int(uniqueID)
	//db.Create(m)

	// Read
	//var product Product
	//db.First(&product, 1) // 根据整型主键查找
	//db.First(&product, "code = ?", "D42") // 查找 code 字段值为 D42 的记录
	//
	//// UpdateSysDict - 将 product 的 price 更新为 200
	//db.Model(&product).UpdateSysDict("Price", 200)
	//// UpdateSysDict - 更新多个字段
	//db.Model(&product).Updates(Product{Price: 200, Code: "F42"}) // 仅更新非零值字段
	//db.Model(&product).Updates(map[string]inter{}{"Price": 200, "Code": "F42"})

	// DeleteSysDictById - 删除 product
	//db.DeleteSysDictById(&product, 1)
}
