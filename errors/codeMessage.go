// Package: errors
// File: errorCode.go
// Created by mint
// Useage: 错误编号 & 错误信息
// DATE: 14-7-9 10:18
package errors

const (
	CODE_OPT_BASE = 100
	CODE_OPT_NO_MORE_DATA = CODE_OPT_BASE - 1
)

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
	CODE_HTTP_ERR_NOT_ALLOW_GET  = CODE_HTTP_ERR_BASE - 3
)

// business error code
const (
	CODE_BISS_ERR_BASE           = -300
	CODE_BISS_ERR_TEL_ALREADY_IN = CODE_BISS_ERR_BASE - 1
	CODE_BISS_ERR_SMS_GATE_FAIL  = CODE_BISS_ERR_BASE - 2
	CODE_BISS_ERR_SMS_CODE       = CODE_BISS_ERR_BASE - 3
	CODE_BISS_ERR_USER_NAME      = CODE_BISS_ERR_BASE - 4
	CODE_BISS_ERR_PASSWORD       = CODE_BISS_ERR_BASE - 5
	CODE_BISS_ERR_VARIFY_CARD    = CODE_BISS_ERR_BASE - 6
	CODE_BISS_ERR_HAS_OWNER      = CODE_BISS_ERR_BASE - 7
	CODE_BISS_ERR_REG            = CODE_BISS_ERR_BASE - 8
	CODE_BISS_ERR_USER_ID        = CODE_BISS_ERR_BASE - 9
	CODE_BISS_ERR_NO_STORE       = CODE_BISS_ERR_BASE - 10
)

// system internal error code
const (
	CODE_SYS_ERR_BASE = -400
)

var _globalWaysErrors map[int]string

func init() {
	// init error
	_globalWaysErrors = make(map[int]string)

	_globalWaysErrors[CODE_SUCCESS] = "everything is ok."

	_globalWaysErrors[CODE_OPT_NO_MORE_DATA] = "没有更多数据啦."

	_globalWaysErrors[CODE_DB_DATA_EXIST] = "data exist"
	_globalWaysErrors[CODE_DB_ERR_NODATA] = "no data exist"
	_globalWaysErrors[CODE_DB_ERR_GET] = "get data error"
	_globalWaysErrors[CODE_DB_ERR_INSERT] = "insert data error"
	_globalWaysErrors[CODE_DB_ERR_UPDATE] = "update data error"
	_globalWaysErrors[CODE_DB_ERR_FIND] = "find data error"

	_globalWaysErrors[CODE_HTTP_ERR_NOT_HTTPS] = "api just allow https connection."
	_globalWaysErrors[CODE_HTTP_ERR_INVALID_PARAMS] = "参数错误."
	_globalWaysErrors[CODE_HTTP_ERR_NOT_ALLOW_GET] = "不允许GET请求."

	_globalWaysErrors[CODE_BISS_ERR_TEL_ALREADY_IN] = "该手机号已被注册."
	_globalWaysErrors[CODE_BISS_ERR_SMS_GATE_FAIL] = "请求短信网关错误."
	_globalWaysErrors[CODE_BISS_ERR_SMS_CODE] = "sms auth code wrong."
	_globalWaysErrors[CODE_BISS_ERR_USER_NAME] = "用户名错误."
	_globalWaysErrors[CODE_BISS_ERR_PASSWORD] = "密码错误."
	_globalWaysErrors[CODE_BISS_ERR_VARIFY_CARD] = "会员卡不正确."
	_globalWaysErrors[CODE_BISS_ERR_HAS_OWNER] = "当前会员卡已被其他用户绑定."
	_globalWaysErrors[CODE_BISS_ERR_REG] = "注册失败."
	_globalWaysErrors[CODE_BISS_ERR_USER_ID] = "会员ID错误."
	_globalWaysErrors[CODE_BISS_ERR_NO_STORE] = "商铺打烊了，请明天再来吧."

	_globalWaysErrors[CODE_SYS_ERR_BASE] = "服务器去月球旅行啦."
}

func GetCodeMessage(code int) string {
	message := "OK."
	if msg, ok := _globalWaysErrors[code]; ok {
		message = msg
	}

	return message
}
