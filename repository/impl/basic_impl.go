/**
 * @Author: Nan
 * @Date: 2024/7/11 上午11:25
 */

package impl

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type BasicImpl struct {
	db *gorm.DB
}

func NewBasicImpl(db *gorm.DB) *BasicImpl {
	return &BasicImpl{
		db,
	}
}

func (b *BasicImpl) ExecuteTx(ctx *gin.Context, fn func(tx *gorm.DB) error) error {
	return b.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		defer func() {
			if r := recover(); r != nil {
				tx.Rollback() // 回滚事务
				if e, ok := r.(error); ok {
					zap.L().Error("Recovered from panic:", zap.Error(e))
				} else {
					zap.L().Error("Recovered from panic:", zap.Any("recover:", r))
				}
			}
		}()
		return fn(tx) // 执行传入的函数
	})
}

func (b *BasicImpl) DBWithContext(ctx *gin.Context) *gorm.DB {
	return b.db.WithContext(ctx)
}
