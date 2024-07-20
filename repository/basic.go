/**
 * @Author: Nan
 * @Date: 2024/7/11 上午11:24
 */

package repository

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"sweet-cms/form/request"
)

type BasicRepository interface {
	ExecuteTx(ctx *gin.Context, fn func(tx *gorm.DB) error) error
	DBWithContext(*gin.Context) *gorm.DB
	CountAsync(*gorm.DB, chan int64, chan error)
	PaginateAndCountAsync(request.Basic, interface{}) (int64, error)
	Create(*gorm.DB, interface{}) error
	Update(*gorm.DB, interface{}) error
	DeleteById(*gorm.DB, int) error
	DeleteByIds(*gorm.DB, []int) error
	FindById(id int) (interface{}, error)
	WithPreload(...string) BasicRepository
	WithContext(*gin.Context) BasicRepository
}
