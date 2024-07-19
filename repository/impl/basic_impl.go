/**
 * @Author: Nan
 * @Date: 2024/7/11 上午11:25
 */

package impl

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"sweet-cms/form/request"
	"sweet-cms/repository/util"
	"sync"
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

func (b *BasicImpl) CountAsync(query *gorm.DB, totalChan chan int64, errorChan chan error) {
	go func() {
		var total int64
		err := query.Limit(-1).Offset(-1).Count(&total).Error
		if err != nil {
			errorChan <- err
		} else {
			totalChan <- total
		}
	}()
}

func (b *BasicImpl) PaginateAndCountAsync(basic request.Basic, result interface{}) (int64, error) {
	var (
		total     int64
		wg        sync.WaitGroup
		err       error
		totalChan = make(chan int64)
		errorChan = make(chan error)
		query     = util.ExecuteQuery(b.db, basic)
	)

	// 异步查询总数
	wg.Add(1)
	go func() {
		defer wg.Done()
		b.CountAsync(query, totalChan, errorChan)
	}()

	// 分页查询
	wg.Add(1)
	go func() {
		defer wg.Done()
		if e := query.Find(result).Error; e != nil {
			errorChan <- e
		}
	}()

	// 等待两个查询完成并关闭通道
	go func() {
		wg.Wait()
		close(totalChan)
		close(errorChan)
	}()

	select {
	case err = <-errorChan:
		return 0, err
	case total = <-totalChan:
		// 如果errorChan中有错误，即使我们已经接收到总数，也应返回错误
		select {
		case err = <-errorChan:
			return 0, err
		default:
			return total, nil
		}
	}
}
