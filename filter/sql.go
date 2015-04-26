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
	"fmt"
	"github.com/aiwuTech/devKit/container"
	"github.com/aiwuTech/devKit/convert"
	"github.com/aiwuTech/devKit/page"
	"github.com/astaxie/beego/orm"
	"log"
	"net/http"
	"net/url"
	"strings"
)

var (
	_ url.Values
)

const (
	Invalid byte = iota
	Exact
	Exact_i
	Contains
	Contains_i
	In
	GT
	GTE
	LT
	LTE
	Start
	Start_i
	End
	End_i
	Null
	Between
	OrderBy
)

const (
	Type_invalid byte = iota
	Type_string
	Type_int
)

const (
	itemSep   = "|"
	insideSep = "-"
	valueSep  = ","
)

// curl -i -d "fields=id,status,created&page=2&size=10&search=status-5-1,4-2|hong_id-6-50000-2" -G 127.0.0.1:8081/v1/admins/members

type SQLFilterItem struct {
	key       string
	operation byte
	value     string
	valType   byte
}

func parseSQLFilterItems(s string) (filterItems []*SQLFilterItem) {
	defer func() {
		for _, item := range filterItems {
			log.Printf("item: %+v\n", item)
		}
	}()

	filterItems = make([]*SQLFilterItem, 0)
	if s == "" {
		return
	}

	items := strings.Split(s, itemSep)
	for _, item := range items {
		insides := strings.Split(item, insideSep)
		if len(insides) != 4 {
			continue
		}

		key := insides[0]
		operation := convert.Str2Byte(insides[1])
		strValue := insides[2]
		valType := convert.Str2Byte(insides[3])

		filterItem := &SQLFilterItem{
			key:       key,
			operation: operation,
			value:     strValue,
			valType:   valType,
		}

		filterItems = append(filterItems, filterItem)
	}

	return
}

func (item *SQLFilterItem) Value() interface{} {
	operation := item.operation
	value := item.value
	valType := item.valType

	switch operation {
	case Exact, Exact_i:
		switch valType {
		case Type_string:
			return value
		case Type_int:
			return convert.Str2Int64(value)
		default:
			return value
		}
	case Contains, Contains_i, Start, Start_i, End, End_i:
		return value
	case GT, GTE, LT, LTE:
		return convert.Str2Int64(value)
	case In, Between:
		stringValues := make([]string, 0)
		if value != "" {
			stringValues = strings.Split(value, valueSep)
		}
		switch valType {
		case Type_string:
			return stringValues
		case Type_int:
			intValues := make([]int64, 0)
			for _, v := range stringValues {
				intValues = append(intValues, convert.Str2Int64(v))
			}
			return intValues
		}
	case Null:
		upperValue := strings.ToUpper(value)
		if upperValue == "TRUE" || upperValue == "1" || upperValue == "YES" {
			return true
		} else {
			return false
		}
	default:
		return value
	}

	return nil
}

type SQLFilter struct {
	pager       *page.Paginator
	fields      []string
	params      map[string]interface{}
	filterItems []*SQLFilterItem
	condition   *orm.Condition
	orderBys    []string

	request *http.Request
}

func (filter *SQLFilter) Pager() *page.Paginator {
	return filter.pager
}

func (filter *SQLFilter) SetCondition(cond *orm.Condition) {
	filter.condition = filter.condition.AndCond(cond)
}

func (filter *SQLFilter) SetOrderBys(orders ...string) {
	for _, order := range orders {
		if order == "" || container.Contains(order, filter.orderBys) {
			continue
		}

		filter.orderBys = append(filter.orderBys, order)
	}
}

func (filter *SQLFilter) SetPager(pageNum, pageSize int) {
	if pageNum > 0 && pageSize > 0 {
		filter.pager = page.NewDBPaginator(pageNum, pageSize)
	}
}

func (filter *SQLFilter) SetFilterItems(items ...*SQLFilterItem) {
	for _, item := range items {
		if item == nil || container.Contains(item, filter.filterItems) {
			continue
		}

		filter.filterItems = append(filter.filterItems, item)
	}
}

func (filter *SQLFilter) Fields() []string {
	return filter.fields
}

func (filter *SQLFilter) SetFields(fields ...string) {
	for _, field := range fields {
		if field == "" || container.Contains(field, filter.fields) {
			continue
		}

		filter.fields = append(filter.fields, field)
	}
}

func (filter *SQLFilter) DelFields(fields ...string) {
	for _, field := range fields {
		container.Delete(field, filter.fields)
	}
}

func (filter *SQLFilter) Params() orm.Params {
	if len(filter.params) != 0 {
		return orm.Params(filter.params)
	}

	params := make(map[string]interface{})
	for _, field := range filter.fields {
		params[field] = filter.request.Form.Get(field)
	}

	log.Printf("params: %+v\n", params)
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

	for _, item := range filter.filterItems {
		key := item.key
		operation := item.operation
		value := item.Value()

		switch operation {
		case Invalid:
			continue
		case Exact:
			querySeter = querySeter.Filter(key, value)
		case Exact_i:
			querySeter = querySeter.Filter(fmt.Sprintf("%s__iexact", key), value)
		case Contains:
			querySeter = querySeter.Filter(fmt.Sprintf("%s__contains", key), value)
		case Contains_i:
			querySeter = querySeter.Filter(fmt.Sprintf("%s__icontains", key), value)
		case In:
			querySeter = querySeter.Filter(fmt.Sprintf("%s__in", key), value)
		case GT:
			querySeter = querySeter.Filter(fmt.Sprintf("%s__gt", key), value)
		case GTE:
			querySeter = querySeter.Filter(fmt.Sprintf("%s__gte", key), value)
		case LT:
			querySeter = querySeter.Filter(fmt.Sprintf("%s__lt", key), value)
		case LTE:
			querySeter = querySeter.Filter(fmt.Sprintf("%s__lte", key), value)
		case Start:
			querySeter = querySeter.Filter(fmt.Sprintf("%s__startswith", key), value)
		case Start_i:
			querySeter = querySeter.Filter(fmt.Sprintf("%s__istartswith", key), value)
		case End:
			querySeter = querySeter.Filter(fmt.Sprintf("%s__endswith", key), value)
		case End_i:
			querySeter = querySeter.Filter(fmt.Sprintf("%s__iendswith", key), value)
		case Null:
			querySeter = querySeter.Filter(fmt.Sprintf("%s__isnull", key), value)
		}
	}

	// condition
	if !filter.condition.IsEmpty() {
		querySeter = querySeter.SetCond(filter.condition)
	}

	// order by
	querySeter = querySeter.OrderBy(filter.orderBys...)

	return querySeter
}

func NewSQLFilter(r *http.Request) *SQLFilter {
	if r.Form == nil {
		r.ParseForm()
	}

	log.Println(r.URL.RequestURI())

	pageNum := convert.Str2Int(r.Form.Get("page"))
	pageSize := convert.Str2Int(r.Form.Get("size"))
	fields := r.Form.Get("fields")
	search := r.Form.Get("search")

	sqlFilter := new(SQLFilter)
	sqlFilter.request = r
	sqlFilter.condition = orm.NewCondition()
	sqlFilter.orderBys = make([]string, 0)

	sqlFilter.SetPager(pageNum, pageSize)
	sqlFilter.SetFields(strings.Split(fields, ",")...)
	sqlFilter.SetFilterItems(parseSQLFilterItems(search)...)

	return sqlFilter
}

func NewDefaultSQLFilter() *SQLFilter {
	return &SQLFilter{
		params:      make(map[string]interface{}),
		fields:      make([]string, 0),
		filterItems: make([]*SQLFilterItem, 0),
		condition:   orm.NewCondition(),
		orderBys:    make([]string, 0),
	}
}
