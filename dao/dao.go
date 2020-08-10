package dao

import (
	"new_gin_project/db"
	"new_gin_project/utils"
)

type Dao struct {
	gormClient *utils.GormDB
}

func New() *Dao {
	return &Dao{
		gormClient: db.GormClient,
	}
}

// 分页查询，resultPointer是一个数组指针
// 同时也会返回一个封装好了的page对象

/*
func (p *Dao) PageQuery(rawSql string, params []interface{}, resultPointer interface{}, pagination *utils.Pagination, orderBys ...string) (page *utils.PageDTO) {
	count := 0
	err := p.gormClient.Client.Raw("select count(*) from ("+rawSql+") tmp", params...).Count(&count).Error
	if err != nil {
		panic(err)
		return
	}
	if count < pagination.Offset() {
		return
	}
	if len(orderBys) > 0 {
		orderBySql := " order by "
		for _, v := range orderBys {
			orderBySql += " " + v
		}
		rawSql += orderBySql
	}
	err = p.gormClient.Client.Raw(rawSql, params...).Offset(pagination.Offset()).Limit(pagination.GetPageSize()).Scan(resultPointer).Error
	if err != nil {
		panic(err)
	}
	pagination.Total = count
	page = &base_proto.PageDTO{
		Total: count,
		List:  resultPointer,
	}
	return
}*/
