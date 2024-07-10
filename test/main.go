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
	"sweet-cms/utils"
	"time"
)

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
		DisableForeignKeyConstraintWhenMigrating: true,
		Logger:                                   dbLogger,
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

	db.Migrator().DropTable(&model.SysConfigure{})
	db.AutoMigrate(&model.SysConfigure{})

	// Create
	m := &model.SysConfigure{EnableCaptcha: false}
	sf, err := utils.NewSnowflake(1)
	if err != nil {
		panic(err)
	}
	uniqueID, err := sf.GenerateUniqueID()
	m.Id = int(uniqueID)
	db.Create(m)

	db.Migrator().DropTable(&model.SysTable{}, &model.SysTableField{}, &model.SysTableRelation{}, &model.SysTableIndex{}, &model.SysTableIndexField{}, &model.SysDict{}, &model.SysDictItem{}, &model.AccessLog{}, &model.LoginLog{})
	// 迁移 schema
	db.AutoMigrate(&model.SysTable{}, &model.SysTableField{}, &model.SysTableRelation{}, &model.SysTableIndex{}, &model.SysTableIndexField{}, &model.SysDict{}, &model.SysDictItem{}, &model.AccessLog{}, &model.LoginLog{})

}
