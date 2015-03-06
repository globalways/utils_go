// Copyright 2014 mit.zhao.chiu@gmail.com
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
package controller

import (
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/validation"
	"github.com/globalways/response"
	"github.com/globalways/utils_go/filter"
	"net/http"
	"net/url"
	"strings"
)

var (
	_ logs.BeeLogger
	_ url.Values
	_ context.BeegoInput
)

type BasicController struct {
	beego.Controller
	valid       *validation.Validation
	fieldErrors []*response.FieldError
}

func (c *BasicController) Prepare() {
	c.valid = new(validation.Validation)
	c.fieldErrors = make([]*response.FieldError, 0)

	//prepare for enable gzip
	c.Ctx.Output.EnableGzip = true
}

func (c *BasicController) Finish() {
	c.valid.Clear()
	c.fieldErrors = c.fieldErrors[:0]
}

func (c *BasicController) IsGet() bool {
	return c.Ctx.Input.IsGet()
}

func (c *BasicController) IsPost() bool {
	return c.Ctx.Input.IsPost()
}

func (c *BasicController) Uri() string {
	return c.Ctx.Input.Uri()
}

func (c *BasicController) Url() string {
	return c.Ctx.Input.Url()
}

func (c *BasicController) IP() string {
	return c.Ctx.Input.IP()
}

// forbiden http get method
func (c *BasicController) ForbidenGet() {
	if c.Ctx.Input.IsGet() {
		c.RenderJson(response.NewResponseMsg(response.Err_Code_Public_Not_Allow_Get))
	}
}

// forbiden http
func (c *BasicController) ForbidenHttp() {
	if !c.Ctx.Input.IsSecure() {
		c.RenderJson(response.NewResponseMsg(response.Err_Code_Public_Not_Allow_Http))
	}
}

// handle http request param error
func (c *BasicController) HandleParamError() bool {
	if c.IsParamsWrong() {
		c.RenderJson(response.NewResponseMsgInvalidParam(fmt.Sprintf("%s%s.", c.fieldErrors[0].Field, c.fieldErrors[0].Message)))
		return true
	}

	return false
}

// parse params is wrong, if wrong, fill response with errors
func (c *BasicController) IsParamsWrong() bool {
	return len(c.fieldErrors) != 0
}

// append a new parameter wrong info
func (c *BasicController) AppenWrongParams(err *response.FieldError) {
	c.fieldErrors = append(c.fieldErrors, err)
}

// valid paramemter
func (c *BasicController) ValidationError(obj interface{}) bool {
	b, err := c.valid.Valid(obj)
	if err != nil {
		c.RenderJson(response.NewResponseMsgInvalidValid(err.Error()))
		return true
	}

	if !b {
		err := c.valid.Errors[0]
		c.RenderJson(response.NewResponseMsgInvalidValid(fmt.Sprintf("%s:%s", err.Key, err.Message)))
		return true
	}

	return false
}

// http json response
func (c *BasicController) RenderJson(data interface{}) {
	c.Data["json"] = data
	c.ServeJson()
}

// http png response
func (c *BasicController) RenderPng(data []byte) {
	c.Ctx.Output.EnableGzip = false
	c.SetHttpContentType("image/png")
	c.SetHttpBody(data)
}

// http plain response
func (c *BasicController) RenderText(data string) {
	c.Ctx.Output.EnableGzip = false
	c.SetHttpContentType("text/plain; charset=utf-8")
	c.SetHttpBody([]byte(data))
}

// set http status
func (c *BasicController) SetHttpStatus(status int) {
	c.Ctx.Output.SetStatus(status)
}

// set http response header
func (c *BasicController) SetHttpHeader(key, val string) {
	c.Ctx.Output.Header(key, val)
}

// set http response body
func (c *BasicController) SetHttpBody(body []byte) {
	c.Ctx.Output.Body(body)
}

// get http request body
func (c *BasicController) GetHttpBody() []byte {
	return c.Ctx.Input.RequestBody
}

// get http request
func (c *BasicController) GetHttpRequest() *http.Request {
	return c.Ctx.Request
}

// get request sql filter
func (c *BasicController) RequestSQLFilter() *filter.SQLFilter {
	return filter.NewSQLFilter(c.GetHttpRequest())
}

// unmarshal json http body to object
func (c *BasicController) HttpBodyUnmarshal(obj interface{}) error {
	err := json.Unmarshal(c.GetHttpBody(), obj)
	if err != nil {
		c.Debug("http body unmarshal error: %v", err)
	}

	return err
}

// set http response contenttype
func (c *BasicController) SetHttpContentType(ext string) {
	c.Ctx.Output.ContentType(ext)
}

// combine url
func (c *BasicController) CombineUrl(router string) string {
	return c.Ctx.Input.Site() + router
}

func (c *BasicController) Debug(format string, v ...interface{}) {
	beego.BeeLogger.Debug(format, v...)
}

func (c *BasicController) Info(format string, v ...interface{}) {
	beego.BeeLogger.Info(format, v...)
}

func (c *BasicController) Error(format string, v ...interface{}) {
	beego.BeeLogger.Error(format, v...)
}

func (c *BasicController) UnmarshalForm() (args map[string]interface{}) {
	args = make(map[string]interface{})
	values := c.Input()
	for k, v := range values {
		args[k] = v[0]
	}

	return
}

// fields
func (c *BasicController) FieldArgs(original map[string]interface{}) (args map[string]interface{}) {
	args = make(map[string]interface{})
	fields := c.Fields()
	if len(fields) != 0 {
		for _, field := range fields {
			if val, ok := original[field]; !ok {
				c.AppenWrongParams(response.NewFieldError(field, fmt.Sprintf("更新参数值%v未传递.", field)))
			} else {
				args[field] = val
			}
		}
	} else {
		args = original
	}

	return
}

func (c *BasicController) Fields() (fields []string) {
	// 拆分fields
	fields = make([]string, 0)
	f := c.GetString("fields")
	if f != "" {
		fields = strings.Split(f, ",")
	}

	return
}

func (c *BasicController) DefaultPageSize() {
	ps := c.GetString("page")
	if ps == "" {
		c.Ctx.Request.Form.Add("page", "1")
	}

	ss := c.GetString("size")
	if ss == "" {
		c.Ctx.Request.Form.Add("size", "10")
	}

	return
}

func (c *BasicController) DeletePageSize() {
	c.Ctx.Request.Form.Del("page")
	c.Ctx.Request.Form.Del("size")
}

func (c *BasicController) Default(obj DefaultValuer) {
	obj.Default()
}
