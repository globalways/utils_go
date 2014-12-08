// Copyright 2014 mint.zhao.chiu@gmail.com. All rights reserved.
// Use of this source code is governed by a Apache License 2.0
// that can be found in the LICENSE file.
package errors

import (
	"encoding/json"
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
