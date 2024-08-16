/**
 * @Author: Nan
 * @Date: 2024/7/11 上午11:25
 */

package impl

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"reflect"
	"sweet-cms/form/request"
	"sweet-cms/repository"
	"sweet-cms/repository/util"
	"sync"
)

type BasicImpl struct {
	db       *gorm.DB
	preloads []string
	model    interface{}
	ctx      *gin.Context
}

func NewBasicImpl(db *gorm.DB, model interface{}) *BasicImpl {
	return &BasicImpl{
		db:    db,
		model: model,
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

func (b *BasicImpl) Create(tx *gorm.DB, entity interface{}) error {
	if b.model == nil {
		return errors.New("model not set")
	}
	modelInstance := reflect.New(reflect.TypeOf(b.model).Elem()).Interface()
	return tx.Model(modelInstance).Create(entity).Error
}

func (b *BasicImpl) Update(tx *gorm.DB, entity interface{}, id int) error {
	if b.model == nil {
		return errors.New("model not set")
	}
	modelInstance := reflect.New(reflect.TypeOf(b.model).Elem()).Interface()
	return tx.Model(modelInstance).Where("id = ?", id).Omit("id").Updates(entity).Error
}

func (b *BasicImpl) DeleteById(tx *gorm.DB, id int) error {
	if b.model == nil {
		return errors.New("model not set")
	}
	modelInstance := reflect.New(reflect.TypeOf(b.model).Elem()).Interface()
	return tx.Delete(modelInstance, id).Error
}

func (b *BasicImpl) DeleteByField(tx *gorm.DB, field string, value interface{}) error {
	if b.model == nil {
		return errors.New("model not set")
	}
	modelInstance := reflect.New(reflect.TypeOf(b.model).Elem()).Interface()
	return tx.Where(fmt.Sprintf("%s = ?", field), value).Delete(modelInstance).Error
}

func (b *BasicImpl) DeleteByIds(tx *gorm.DB, ids []int) error {
	if b.model == nil {
		return errors.New("model not set")
	}
	modelInstance := reflect.New(reflect.TypeOf(b.model).Elem()).Interface()
	return tx.Where("id in ?", ids).Delete(modelInstance).Error
}

func (b *BasicImpl) DeleteByFieldIn(tx *gorm.DB, field string, values []interface{}) error {
	if b.model == nil {
		return errors.New("model not set")
	}
	modelInstance := reflect.New(reflect.TypeOf(b.model).Elem()).Interface()
	return tx.Where(fmt.Sprintf("%s in ?", field), values).Delete(modelInstance).Error
}

func (b *BasicImpl) FindById(id int) (interface{}, error) {
	if b.model == nil {
		return nil, errors.New("model not set")
	}
	entity := reflect.New(reflect.TypeOf(b.model).Elem()).Interface()
	query := b.db
	if b.ctx != nil {
		query = query.WithContext(b.ctx)
	}
	for _, preload := range b.preloads {
		query = query.Preload(preload)
	}
	err := query.First(entity, id).Error
	return reflect.ValueOf(entity).Elem().Interface(), err
}

func (b *BasicImpl) FindByField(field string, value interface{}) (interface{}, error) {
	if b.model == nil {
		return nil, errors.New("model not set")
	}
	entity := reflect.New(reflect.TypeOf(b.model).Elem()).Interface()
	query := b.db
	if b.ctx != nil {
		query = query.WithContext(b.ctx)
	}
	for _, preload := range b.preloads {
		query = query.Preload(preload)
	}
	err := query.Where(fmt.Sprintf("%s = ?", field), value).First(&entity).Error
	return entity, err
}

func (b *BasicImpl) WithPreload(preloads ...string) repository.BasicRepository {
	newImpl := *b
	newImpl.preloads = append(newImpl.preloads, preloads...)
	return &newImpl
}

func (b *BasicImpl) WithContext(ctx *gin.Context) repository.BasicRepository {
	newImpl := *b
	newImpl.ctx = ctx
	return &newImpl
}
