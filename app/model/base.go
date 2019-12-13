package model

import (
	"database/sql"
	"errors"
	"github.com/gogf/gf/database/gdb"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/os/glog"
)

var db gdb.DB

func init() {
	db = g.DB()
	db.SetDebug(true)
	db.SetMaxConnLifetime(30)
	db.SetMaxIdleConnCount(10)
	db.SetMaxOpenConnCount(100)
}

type IBase interface {
	GetTableName() string
}

type option struct {
	Field   string
	Where   interface{}
	Order   string
	Limit   []int
	GroupBy string
	Data    interface{}
}

type ModOption func(option *option)

func WithWhere(where interface{}) ModOption {
	return func(option *option) {
		option.Where = where
	}
}

func WithField(field string) ModOption {
	return func(option *option) {
		option.Field = field
	}
}

func WithOrder(order string) ModOption {
	return func(option *option) {
		option.Order = order
	}
}

func WithGroup(group string) ModOption {
	return func(option *option) {
		option.GroupBy = group
	}
}

func WithLimit(limit []int) ModOption {
	return func(option *option) {
		option.Limit = limit
	}
}

func WithData(data interface{}) ModOption {
	return func(option *option) {
		option.Data = data
	}
}

func getOptions(modOptions ...ModOption) *option {
	op := &option{}
	for _, fn := range modOptions {
		fn(op)
	}
	return op
}

func FindAll(obj IBase, modOptions ...ModOption) (g.List, error) {

	op := getOptions(modOptions...)
	safe := db.Table(obj.GetTableName()).Safe()
	if op.Field != "" {
		safe = safe.Fields(op.Field)
	}
	if op.Where != nil {
		safe = safe.Where(op.Where)
	}
	if op.Limit != nil {
		safe = safe.Limit(op.Limit...)
	}
	if op.Order != "" {
		safe = safe.OrderBy(op.Order)
	}
	if op.GroupBy != "" {
		safe = safe.GroupBy(op.GroupBy)
	}
	results, e := safe.Select()

	if e != nil {
		glog.Error(e)
		return nil, e
	}

	return results.List(), nil
}

func FindOne(obj IBase, modOptions ...ModOption) (g.Map, error) {
	op := getOptions(modOptions...)
	safe := db.Table(obj.GetTableName()).Safe()
	if op.Field != "" {
		safe = safe.Fields(op.Field)
	}
	if op.Where != nil {
		safe = safe.Where(op.Where)
	}
	records, e := safe.One()
	if e != nil && e != sql.ErrNoRows {
		glog.Error(e)
		return nil, e
	}
	return records.Map(), nil
}

func Insert(obj IBase, modOptions ...ModOption) (int64, error) {
	safe := db.Table(obj.GetTableName()).Safe()
	op := getOptions(modOptions...)
	if op.Data != nil {
		safe = safe.Data(op.Data)
	}
	insert, err := safe.Insert()
	if err != nil {
		return 0, err
	}
	insertId, err := insert.LastInsertId()
	if err != nil {
		return insertId, err
	}
	return insertId, nil
}

func Update(obj IBase, modOptions ...ModOption) (int64, error) {
	safe := db.Table(obj.GetTableName()).Safe()
	op := getOptions(modOptions...)
	if op.Data == nil {
		return 0, errors.New("修改信息为空")
	}
	safe = safe.Data(op.Data)
	if op.Where != nil {
		safe = safe.Where(op.Where)
	}

	if op.Limit != nil {
		safe = safe.Limit(op.Limit...)
	}
	result, err := safe.Update()
	if err != nil {
		return 0, err
	}
	updateId, err := result.RowsAffected()
	if err != nil {
		return 0, err
	}
	return updateId, nil
}
