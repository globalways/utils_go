// Copyright 2014 mint.zhao.chiu@gmail.com. All rights reserved.
// Use of this source code is governed by a Apache License 2.0
// that can be found in the LICENSE file.
package errors

import (
	"encoding/json"
	"fmt"
)

type CommonResponse struct {
	Code        int    `json:code`
	Message     string `json:message`
	Description string `jsong:description`
}

type FieldError struct {
	Field   string `json:field`
	Message string `json:message`
}

type FieldErrors struct {
	Code    int           `json:code`
	Message string        `json:message`
	Errors  []*FieldError `json:errors`
}

// new common response
func NewCommonOutRsp(gErr GlobalWaysError) *CommonResponse {
	code := gErr.GetCode()
	msg := gErr.GetMessage()
	desc := gErr.GetInner().Error()

	return &CommonResponse{
		Code:        code,
		Message:     msg,
		Description: desc,
	}
}

// new fielderror
func NewFieldError(field string, msg string) *FieldError {
	return &FieldError{
		Field:   field,
		Message: msg,
	}
}

// new fieldErrors
func NewFiledErrors(code int, errs []*FieldError) *FieldErrors {
	return &FieldErrors{
		Code:    code,
		Message: GetCodeMessage(code),
		Errors:  errs,
	}
}

func UnmarshalFiledErrors(bytes []byte) *FieldErrors {
	fieldErrors := new(FieldErrors)
	json.Unmarshal(bytes, fieldErrors)

	return fieldErrors
}

func UnmarshalCommonResponse(bytes []byte) *CommonResponse {
	commonRsp := new(CommonResponse)
	json.Unmarshal(bytes, commonRsp)

	return commonRsp
}

// 状态码json
type Status struct {
	Code    int    `json:"code"`
	Message string `json:"msg"`
}

// 客户端返回json
type ClientRsp struct {
	Status *Status     `json:"status"`
	Body   interface{} `json:"body"`
}

// 新建格式化状态码
func newStatusf(code int, format string, args ...interface{}) *Status {
	return &Status{
		Code:    code,
		Message: fmt.Sprintf(format, args...),
	}
}

// 新建状态码
func NewStatus(code int) *Status {
	return newStatusf(code, GetCodeMessage(code))
}

// 新建格式化客户端返回值
func NewClientRspf(code int, format string, args ...interface{}) *ClientRsp {
	return &ClientRsp{
		Status: newStatusf(code, format, args...),
	}
}

// 新建固定客户端返回值
func NewClientRsp(code int) *ClientRsp {
	return &ClientRsp{
		Status: newStatusf(code, GetCodeMessage(code)),
	}
}

// 新建globalwaysError错误
func NewGlobalwaysErrorRsp(gErr GlobalWaysError) *ClientRsp {
	return &ClientRsp{
		Status: newStatusf(gErr.GetCode(), gErr.GetMessage()),
	}
}
