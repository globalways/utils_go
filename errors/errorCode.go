// Package: errors
// File: errorCode.go
// Created by mint
// Useage: 错误编号 & 错误信息
// DATE: 14-7-9 10:18
package errors

//错误编号
const (
	CODE_SUCCESS = 0

	//DB ERROR
	CODE_DB_ERR_BASE    = -100
	CODE_DB_ERR_BADCONN = CODE_DB_ERR_BASE - 1
	CODE_DB_ERR_NODATA  = CODE_DB_ERR_BASE - 2
	CODE_DB_ERR_GET     = CODE_DB_ERR_BASE - 3
	CODE_DB_ERR_FIND    = CODE_DB_ERR_BASE - 4
	CODE_DB_ERR_INSERT  = CODE_DB_ERR_BASE - 5
	CODE_DB_ERR_UPDATE  = CODE_DB_ERR_BASE - 6
	CODE_DB_ERR_COMMIT  = CODE_DB_ERR_BASE - 7
	CODE_DB_DATA_EXIST  = CODE_DB_ERR_BASE - 8
)

//错误信息
const (
	MSG_SUCCESS = "Get data successful"

	MSG_DB_ERR_COMMIT = "Commit error"
)

// http error code
const (
	CODE_INVALID_PARAMS     = -201
	CODE_HTTP_ERR_GET       = -202
	CODE_HTTP_READ_REQ_BODY = -203
)

// json error code
const (
	CODE_JSON_ERR_BASE     = -300
	CODE_JSON_ERR_MASHAL   = CODE_JSON_ERR_BASE - 1
	CODE_JSON_ERR_UNMASHAL = CODE_JSON_ERR_BASE - 2
)

func init() {
	// init error
	GlobalWaysErrors = make(map[int]string)

	GlobalWaysErrors[CODE_DB_DATA_EXIST] = "data exist"
	GlobalWaysErrors[CODE_DB_ERR_NODATA] = "no data exist"
	GlobalWaysErrors[CODE_DB_ERR_GET] = "get data error: %v"
	GlobalWaysErrors[CODE_DB_ERR_INSERT] = "insert data error: %v"
	GlobalWaysErrors[CODE_DB_ERR_UPDATE] = "update data error: %v"
	GlobalWaysErrors[CODE_DB_ERR_FIND] = "find data error: %v"

	GlobalWaysErrors[CODE_INVALID_PARAMS] = "invalid parameters"
	GlobalWaysErrors[CODE_HTTP_ERR_GET] = "http get error: %v"
	GlobalWaysErrors[CODE_HTTP_READ_REQ_BODY] = "read http request body error: %v"

	//json
	GlobalWaysErrors[CODE_JSON_ERR_MASHAL] = "json marshal error: %v"
	GlobalWaysErrors[CODE_JSON_ERR_UNMASHAL] = "json unmarshal error: %v"
}
