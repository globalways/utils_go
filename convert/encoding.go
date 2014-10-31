// Copyright 2014 mint.zhao.chiu@gmail.com. All rights reserved.
// Use of this source code is governed by a Apache License 2.0
// that can be found in the LICENSE file.
package convert

import (
	"net/url"
	"github.com/axgle/mahonia"
)


//utf-8转gbk
func Utf8ToGBK(str string) string {
	//字符集转换
	enc := mahonia.NewEncoder("gbk")
	return enc.ConvertString(str)
}

//gbk转utf-8
func GBKToUtf8(str string) string {
	//字符集转换
	enc := mahonia.NewDecoder("gbk")
	return enc.ConvertString(str)
}

//url编码
func UrlEncode(s string) string {
	return url.QueryEscape(s)
}
