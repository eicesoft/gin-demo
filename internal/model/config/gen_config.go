///////////////////////////////////////////////////////////
// THIS FILE IS AUTO GENERATED by gormgen, DON'T EDIT IT //
//        ANY CHANGES DONE HERE WILL BE LOST             //
///////////////////////////////////////////////////////////

package sys_configs

import (
	"fmt"

	"eicesoft/web-demo/internal/model"
	"eicesoft/web-demo/pkg/core"

	"github.com/pkg/errors"
	"gorm.io/gorm"
)

func NewModel() *Config {
	return new(Config)
}

func NewQueryBuilder() *sysConfigsQueryBuilder {
	return new(sysConfigsQueryBuilder)
}

func (t *Config) Assign(src interface{}) {
	core.StructCopy(t, src)
}

func (t *Config) Create(db *gorm.DB) (id int32, err error) {
	if err = db.Create(t).Error; err != nil {
		return 0, errors.Wrap(err, "create err")
	}
	return t.Id, nil
}

func (t *Config) Delete(db *gorm.DB) (err error) {
	if err = db.Delete(t).Error; err != nil {
		return errors.Wrap(err, "delete err")
	}
	return nil
}

func (t *Config) Updates(db *gorm.DB, m map[string]interface{}) (err error) {
	if err = db.Model(&Config{}).Where("id = ?", t.Id).Updates(m).Error; err != nil {
		return errors.Wrap(err, "updates err")
	}
	return nil
}

type sysConfigsQueryBuilder struct {
	order []string
	where []struct {
		prefix string
		value  interface{}
	}
	limit  int
	offset int
}

func (qb *sysConfigsQueryBuilder) Updates(db *gorm.DB, m map[string]interface{}) (err error) {
	if err = qb.buildUpdateQuery(db).Updates(m).Error; err != nil {
		return errors.Wrap(err, "updates err")
	}
	return nil
}

func (qb *sysConfigsQueryBuilder) buildUpdateQuery(db *gorm.DB) *gorm.DB {
	ret := db.Model(&Config{})
	for _, where := range qb.where {
		ret = ret.Where(where.prefix, where.value)
	}
	return ret
}

func (qb *sysConfigsQueryBuilder) buildQuery(db *gorm.DB) *gorm.DB {
	ret := db
	for _, where := range qb.where {
		ret = ret.Where(where.prefix, where.value)
	}
	for _, order := range qb.order {
		ret = ret.Order(order)
	}
	ret = ret.Limit(qb.limit).Offset(qb.offset)
	return ret
}

func (qb *sysConfigsQueryBuilder) Count(db *gorm.DB) (int64, error) {
	var c int64
	res := qb.buildQuery(db).Model(&Config{}).Count(&c)
	if res.Error != nil && res.Error == gorm.ErrRecordNotFound {
		c = 0
	}
	return c, res.Error
}

func (qb *sysConfigsQueryBuilder) First(db *gorm.DB) (*Config, error) {
	ret := &Config{}
	res := qb.buildQuery(db).First(ret)
	if res.Error != nil && res.Error == gorm.ErrRecordNotFound {
		ret = nil
	}
	return ret, res.Error
}

func (qb *sysConfigsQueryBuilder) QueryOne(db *gorm.DB) (*Config, error) {
	qb.limit = 1
	ret, err := qb.QueryAll(db)
	if len(ret) > 0 {
		return ret[0], err
	}
	return nil, err
}

func (qb *sysConfigsQueryBuilder) QueryAll(db *gorm.DB) ([]*Config, error) {
	var ret []*Config
	err := qb.buildQuery(db).Find(&ret).Error
	return ret, err
}

func (qb *sysConfigsQueryBuilder) Limit(limit int) *sysConfigsQueryBuilder {
	qb.limit = limit
	return qb
}

func (qb *sysConfigsQueryBuilder) Offset(offset int) *sysConfigsQueryBuilder {
	qb.offset = offset
	return qb
}

func (qb *sysConfigsQueryBuilder) WhereId(p model.Predicate, value int32) *sysConfigsQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "id", p),
		value,
	})
	return qb
}

func (qb *sysConfigsQueryBuilder) WhereIdIn(value []int32) *sysConfigsQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "id", "IN"),
		value,
	})
	return qb
}

func (qb *sysConfigsQueryBuilder) WhereIdNotIn(value []int32) *sysConfigsQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "id", "NOT IN"),
		value,
	})
	return qb
}

func (qb *sysConfigsQueryBuilder) OrderById(asc bool) *sysConfigsQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "id "+order)
	return qb
}

func (qb *sysConfigsQueryBuilder) WhereTitle(p model.Predicate, value string) *sysConfigsQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "title", p),
		value,
	})
	return qb
}

func (qb *sysConfigsQueryBuilder) WhereTitleIn(value []string) *sysConfigsQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "title", "IN"),
		value,
	})
	return qb
}

func (qb *sysConfigsQueryBuilder) WhereTitleNotIn(value []string) *sysConfigsQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "title", "NOT IN"),
		value,
	})
	return qb
}

func (qb *sysConfigsQueryBuilder) OrderByTitle(asc bool) *sysConfigsQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "title "+order)
	return qb
}

func (qb *sysConfigsQueryBuilder) WhereKey(p model.Predicate, value string) *sysConfigsQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "key", p),
		value,
	})
	return qb
}

func (qb *sysConfigsQueryBuilder) WhereKeyIn(value []string) *sysConfigsQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "key", "IN"),
		value,
	})
	return qb
}

func (qb *sysConfigsQueryBuilder) WhereKeyNotIn(value []string) *sysConfigsQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "key", "NOT IN"),
		value,
	})
	return qb
}

func (qb *sysConfigsQueryBuilder) OrderByKey(asc bool) *sysConfigsQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "key "+order)
	return qb
}

func (qb *sysConfigsQueryBuilder) WhereValue(p model.Predicate, value string) *sysConfigsQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "value", p),
		value,
	})
	return qb
}

func (qb *sysConfigsQueryBuilder) WhereValueIn(value []string) *sysConfigsQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "value", "IN"),
		value,
	})
	return qb
}

func (qb *sysConfigsQueryBuilder) WhereValueNotIn(value []string) *sysConfigsQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "value", "NOT IN"),
		value,
	})
	return qb
}

func (qb *sysConfigsQueryBuilder) OrderByValue(asc bool) *sysConfigsQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "value "+order)
	return qb
}

func (qb *sysConfigsQueryBuilder) WhereType(p model.Predicate, value int32) *sysConfigsQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "type", p),
		value,
	})
	return qb
}

func (qb *sysConfigsQueryBuilder) WhereTypeIn(value []int32) *sysConfigsQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "type", "IN"),
		value,
	})
	return qb
}

func (qb *sysConfigsQueryBuilder) WhereTypeNotIn(value []int32) *sysConfigsQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "type", "NOT IN"),
		value,
	})
	return qb
}

func (qb *sysConfigsQueryBuilder) OrderByType(asc bool) *sysConfigsQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "type "+order)
	return qb
}

func (qb *sysConfigsQueryBuilder) WhereCreatedAt(p model.Predicate, value int32) *sysConfigsQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "created_at", p),
		value,
	})
	return qb
}

func (qb *sysConfigsQueryBuilder) WhereCreatedAtIn(value []int32) *sysConfigsQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "created_at", "IN"),
		value,
	})
	return qb
}

func (qb *sysConfigsQueryBuilder) WhereCreatedAtNotIn(value []int32) *sysConfigsQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "created_at", "NOT IN"),
		value,
	})
	return qb
}

func (qb *sysConfigsQueryBuilder) OrderByCreatedAt(asc bool) *sysConfigsQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "created_at "+order)
	return qb
}

func (qb *sysConfigsQueryBuilder) WhereUpdatedAt(p model.Predicate, value int32) *sysConfigsQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "updated_at", p),
		value,
	})
	return qb
}

func (qb *sysConfigsQueryBuilder) WhereUpdatedAtIn(value []int32) *sysConfigsQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "updated_at", "IN"),
		value,
	})
	return qb
}

func (qb *sysConfigsQueryBuilder) WhereUpdatedAtNotIn(value []int32) *sysConfigsQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "updated_at", "NOT IN"),
		value,
	})
	return qb
}

func (qb *sysConfigsQueryBuilder) OrderByUpdatedAt(asc bool) *sysConfigsQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "updated_at "+order)
	return qb
}

func (qb *sysConfigsQueryBuilder) WhereIsDelete(p model.Predicate, value int32) *sysConfigsQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "is_delete", p),
		value,
	})
	return qb
}

func (qb *sysConfigsQueryBuilder) WhereIsDeleteIn(value []int32) *sysConfigsQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "is_delete", "IN"),
		value,
	})
	return qb
}

func (qb *sysConfigsQueryBuilder) WhereIsDeleteNotIn(value []int32) *sysConfigsQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "is_delete", "NOT IN"),
		value,
	})
	return qb
}

func (qb *sysConfigsQueryBuilder) OrderByIsDelete(asc bool) *sysConfigsQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "is_delete "+order)
	return qb
}
