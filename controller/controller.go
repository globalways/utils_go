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
	"github.com/globalways/utils_go/errors"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/validation"
)

type BasicController struct {
	beego.Controller
	valid *validation.Validation
	fieldErrors []*errors.FieldError
}

func (c *BasicController) Prepare() {
	//prepare for enable gzip
	c.Ctx.Output.EnableGzip = true
}

func (c *BasicController) Finish() {
	c.fieldErrors = c.fieldErrors[:0]
	c.valid.Clear()
}

// forbiden http get method
func (c *BasicController) ForbidenGet() {
	if c.Ctx.Input.IsGet() {
		c.RenderJson(errors.NewClientRsp(errors.CODE_HTTP_ERR_NOT_ALLOW_GET))
	}
}

// forbiden http
func (c *BasicController) ForbidenHttp() {
	if !c.Ctx.Input.IsSecure() {
		c.RenderJson(errors.NewClientRsp(errors.CODE_HTTP_ERR_NOT_ALLOW_HTTP))
	}
}

// handle http request param error
func (c *BasicController) HandleParamError() bool {
	if c.isParamsWrong() {
		c.RenderJson(errors.NewClientRspf(errors.CODE_HTTP_ERR_INVALID_PARAMS, c.fieldErrors[0].Message))

		for _, err := range c.fieldErrors {
			beego.BeeLogger.Debug("filedError: %v", err)
		}

		return true
	}

	return false
}

// parse params is wrong, if wrong, fill response with errors
func (c *BasicController) isParamsWrong() bool {
	return len(c.fieldErrors) != 0
}

// append a new parameter wrong info
func (c *BasicController) AppenWrongParams(err *errors.FieldError) {
	c.fieldErrors = append(c.fieldErrors, err)
}

// valid paramemter
func (c *BasicController) Validation(obj interface{}) {
	b, err := c.valid.Valid(obj)
	if err != nil {
		c.AppenWrongParams(errors.NewFieldError("valid", err.Error()))
	}

	if !b {
		for _, err := range c.valid.Errors {
			c.AppenWrongParams(errors.NewFieldError(err.Key, err.Message))
		}
	}
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

// http internal error
func (c *BasicController) RenderInternalError() {
	c.RenderJson(errors.NewClientRsp(errors.CODE_SYS_ERR_BASE))
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

// set http response contenttype
func (c *BasicController) SetHttpContentType(ext string) {
	c.Ctx.Output.ContentType(ext)
}

// combine url
func (c *BasicController) CombineUrl(router string) string {
	return c.Ctx.Input.Site() + router
}
