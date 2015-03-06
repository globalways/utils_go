// Copyright 2015 mint.zhao.chiu@gmail.com
//
// Licensed under the Apache License, Version 2.0 (the "License"): you may
// not use this file except in compliance with the License. You may obtain
// a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS, WITHOUT
// WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the
// License for the specific language governing permissions and limitations
// under the License.
package filter

import (
	"github.com/astaxie/beego/orm"
	"github.com/globalways/utils_go/container"
	"github.com/globalways/utils_go/convert"
	"github.com/globalways/utils_go/page"
	"net/http"
	"net/url"
	"strings"
)

var (
	_ url.Values
)

type SQLFilter struct {
	pager   *page.Paginator
	orderBy []string
	fields  []string
	params  map[string]interface{}
	request *http.Request
}

func (filter *SQLFilter) Pager() *page.Paginator {
	return filter.pager
}

func (filter *SQLFilter) OrderBy() []string {
	return filter.orderBy
}

func (filter *SQLFilter) SetOrderBys(orderBys ...string) {
	for _, orderBy := range orderBys {
		if container.InArray(orderBy, filter.orderBy) {
			continue
		}

		filter.orderBy = append(filter.orderBy, orderBy)
	}
}

func (filter *SQLFilter) Fields() []string {
	return filter.fields
}

func (filter *SQLFilter) SetFields(fields ...string) {
	for _, field := range fields {
		if container.InArray(field, filter.fields) {
			continue
		}

		filter.fields = append(filter.fields, field)
	}
}

func (filter *SQLFilter) DelFields(fields ...string) {

}

func (filter *SQLFilter) Params() orm.Params {
	if len(filter.params) != 0 {
		return orm.Params(filter.params)
	}

	params := make(map[string]interface{})
	for _, field := range filter.fields {
		params[field] = filter.request.Form.Get(field)
	}

	return orm.Params(params)
}

func (filter *SQLFilter) SetParam(key string, val interface{}) {
	filter.params[key] = val
}

func (filter *SQLFilter) SetParams(vals map[string]interface{}) {
	for key, val := range vals {
		filter.params[key] = val
	}
}

func (filter *SQLFilter) Query(querySeter orm.QuerySeter) orm.QuerySeter {
	if querySeter == nil {
		return nil
	}

	pager := filter.Pager()
	if pager != nil {
		querySeter = querySeter.Limit(pager.Size(), pager.Offset())
	}

	orderBy := filter.OrderBy()
	if len(orderBy) != 0 {
		querySeter = querySeter.OrderBy(orderBy...)
	}

	return querySeter
}

func NewSQLFilter(r *http.Request) *SQLFilter {
	if r.Form == nil {
		r.ParseForm()
	}

	pageNum := convert.Str2Int(r.Form.Get("page"))
	pageSize := convert.Str2Int(r.Form.Get("size"))
	orderby := r.Form.Get("orderby")
	fields := r.Form.Get("fields")

	sqlFilter := new(SQLFilter)
	sqlFilter.request = r
	if pageNum > 0 && pageSize > 0 {
		sqlFilter.pager = page.NewDBPaginator(pageNum, pageSize)
	}
	if orderby != "" {
		sqlFilter.orderBy = strings.Split(orderby, ",")
	}
	if fields != "" {
		sqlFilter.fields = strings.Split(fields, ",")
	}

	return sqlFilter
}

func NewDefaultSQLFilter() *SQLFilter {
	return &SQLFilter{
		params:  make(map[string]interface{}),
		orderBy: make([]string, 0),
		fields:  make([]string, 0),
	}
}
