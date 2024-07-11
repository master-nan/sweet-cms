/**
 * @Author: Nan
 * @Date: 2024/7/11 上午11:24
 */

package repository

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type BasicRepository interface {
	ExecuteTx(ctx *gin.Context, fn func(tx *gorm.DB) error) error
	DBWithContext(*gin.Context) *gorm.DB
}
