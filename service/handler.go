/**
 * @Author: Nan
 * @Date: 2024/5/17 下午5:52
 */

package service

import (
	"encoding/json"
	"gorm.io/gorm"
	"net/http"
	"sweet-cms/enum"
	"sweet-cms/form/request"
	"time"
)

func parseValue(value interface{}, valueType enum.SysTableFieldType) interface{} {
	switch valueType {
	case enum.INT:
		return value.(int)
	case enum.FLOAT:
		return value.(float64)
	case enum.VARCHAR:
		return value.(string)
	case enum.BOOLEAN:
		return value.(bool)
	case enum.TEXT:
		return value.(string)
	case enum.DATE:
		t, _ := time.Parse(time.DateOnly, value.(string))
		return t
	case enum.DATETIME:
		t, _ := time.Parse(time.DateTime, value.(string))
		return t
	case enum.TIME:
		t, _ := time.Parse(time.TimeOnly, value.(string))
		return t
	default:
		return value
	}
}

func searchHandler(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req request.Basic
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		query := db
		for _, cond := range req.Query {
			value := parseValue(cond.Value, cond.Type)
			switch cond.Expression {
			case request.GT:
				query = query.Where(cond.Field+" > ?", value)
			case request.LT:
				query = query.Where(cond.Field+" < ?", value)
			case request.GTE:
				query = query.Where(cond.Field+" >= ?", value)
			case request.LTE:
				query = query.Where(cond.Field+" <= ?", value)
			case request.EQ:
				query = query.Where(cond.Field+" = ?", value)
			case request.NE:
				query = query.Where(cond.Field+" != ?", value)
			case request.LIKE:
				query = query.Where(cond.Field+" LIKE ?", value)
			case request.NOT_LIKE:
				query = query.Where(cond.Field+" NOT LIKE ?", value)
			case request.IN:
				query = query.Where(cond.Field+" IN (?)", value)
			case request.NOT_IN:
				query = query.Where(cond.Field+" NOT IN (?)", value)
			case request.IS_NULL:
				query = query.Where(cond.Field + " IS NULL")
			case request.IS_NOT_NULL:
				query = query.Where(cond.Field + " IS NOT NULL")
			default:
				http.Error(w, "unsupported operator", http.StatusBadRequest)
				return
			}
		}
		if req.Page > 0 && req.Num > 0 {
			query = query.Limit(req.Page).Offset(req.Page * (req.Num - 1))
		}
		return
	}
}