// Package: db
// File: db.go
// Created by mint
// Useage: 数据库操作类
// DATE: 14-7-8 21:32
package db

import (
	"utils/errors"
	"utils/page"
)

//数据库操作接口 连接 增删改查
type DBConnector interface {
	Connect(driverName, dataSourceName string) errors.GlobalWaysError
	DisConnect()
}

type DBReader interface {
	DBConnector

	//根据主键获取一条记录，复合主键不能使用
	GetById(id, bean interface{}) (bool, errors.GlobalWaysError)

	//根据某一列获取一条记录，筛选条件：condi，其余条件不能使用。(可以在bean中添加其他＝条件)
	GetByCol(colName, condi string, col, bean interface{}) (bool, errors.GlobalWaysError)

	//根据多列获取多条记录，筛选条件：=，其余条件不能使用
	FindByCol(bean, condi interface{}, pager *page.Page) (bool, errors.GlobalWaysError)
}

type DBWriter interface {
	DBConnector

	Insert(bean interface{}) (bool, errors.GlobalWaysError)
	Update(bean interface{}, condiBean ...interface{}) (bool, errors.GlobalWaysError)
}
