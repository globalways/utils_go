// Package: db
// File: xorm.go
// Created by mint
// Useage: xorm封装类
// DATE: 14-7-9 15:32
package db

import (
	"fmt"
	"github.com/go-xorm/core"
	"github.com/go-xorm/xorm"
	"utils/errors"
	"utils/page"
)

// 更改底层orm更简单
type XormEngine struct {
	xorm *xorm.Engine
}

//连接数据库
func (c *XormEngine) Connect(driverName, dataSourceName string) errors.GlobalWaysError {
	engine, err := xorm.NewEngine(driverName, dataSourceName)
	if err != nil {
		return errors.Newf(errors.CODE_DB_ERR_BADCONN, "DB Connection ret err: %v", err)
	}
	engine.SetMapper(core.SameMapper{})
	engine.ShowSQL = true

	c.xorm = engine

	return errors.ErrorOK()
}

//中断数据库连接
func (c *XormEngine) DisConnect() {
	c.xorm.Close()
}

//Xorm'Engine DBReader
func (r *XormEngine) GetById(id, bean interface{}) (bool, errors.GlobalWaysError) {

	if has, err := r.xorm.Id(id).Get(bean); err != nil {
		return false, errors.Newf(errors.CODE_DB_ERR_GETByID, "SelectById return error: %v", err)
	} else if !has {
		return false, errors.New(errors.CODE_DB_ERR_NODATA, errors.MSG_DB_ERR_NODATA)
	}

	return true, errors.ErrorOK()
}

func (r *XormEngine) GetByCol(colName, condi string, col, bean interface{}) (bool, errors.GlobalWaysError) {

	if has, err := r.xorm.Where(fmt.Sprintf("%s %s ?", colName, condi), col).Get(bean); err != nil {
		return false, errors.Newf(errors.CODE_DB_ERR_GETByCOL, "SelectByCol return error: %v", err)
	} else if !has {
		return false, errors.New(errors.CODE_DB_ERR_NODATA, errors.MSG_DB_ERR_NODATA)
	}

	return true, errors.ErrorOK()
}

func (r *XormEngine) FindByCol(bean, condiBean interface{}, pager *page.Page) (bool, errors.GlobalWaysError) {

	if err := r.xorm.Limit(int(pager.Perpage), int((pager.Current_page-1)*pager.Perpage)).Find(bean, condiBean); err != nil {
		return false, errors.Newf(errors.CODE_DB_ERR_FINDByCol, "FindByCol return error: %v", err)
	}

	return true, errors.ErrorOK()
}

//Xorm'Engine DBWriter
func (w *XormEngine) Insert(bean interface{}) (bool, errors.GlobalWaysError) {

	if _, err := w.xorm.Insert(bean); err != nil {
		return false, errors.Newf(errors.CODE_DB_ERR_INSERT, "Insert return error: %v", err)
	}

	return true, errors.ErrorOK()
}

func (w *XormEngine) Update(bean interface{}, condiBean ...interface{}) (bool, errors.GlobalWaysError) {

	if _, err := w.xorm.Update(bean, condiBean...); err != nil {
		return false, errors.Newf(errors.CODE_DB_ERR_UPDATE, "Update return error: %v", err)
	}

	return true, errors.ErrorOK()
}
