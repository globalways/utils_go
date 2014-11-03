// Package: errors
// File: errorCode.go
// Created by mint
// Useage: 错误编号 & 错误信息
// DATE: 14-7-9 10:18
package errors

const (
	CODE_SUCCESS = 0
)

// DB ERROR
const (
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

// http error code
const (
	CODE_HTTP_ERR_BASE           = -200
	CODE_HTTP_ERR_NOT_HTTPS      = CODE_HTTP_ERR_BASE - 1
	CODE_HTTP_ERR_INVALID_PARAMS = CODE_HTTP_ERR_BASE - 2
)

var _globalWaysErrors map[int]string
func init() {
	// init error
	_globalWaysErrors = make(map[int]string)

	_globalWaysErrors[CODE_DB_DATA_EXIST] = "data exist"
	_globalWaysErrors[CODE_DB_ERR_NODATA] = "no data exist"
	_globalWaysErrors[CODE_DB_ERR_GET] = "get data error"
	_globalWaysErrors[CODE_DB_ERR_INSERT] = "insert data error"
	_globalWaysErrors[CODE_DB_ERR_UPDATE] = "update data error"
	_globalWaysErrors[CODE_DB_ERR_FIND] = "find data error"

	_globalWaysErrors[CODE_HTTP_ERR_NOT_HTTPS] = "api just allow https connection."
	_globalWaysErrors[CODE_HTTP_ERR_INVALID_PARAMS] = "invalid parameters."
}

func GetCodeMessage(code int) string {
	message := "something bad happend."
	if msg, ok := _globalWaysErrors[code]; ok {
		message = msg
	}

	return message
}
