// Copyright 2014 mint.zhao.chiu@gmail.com. All rights reserved.
// Use of this source code is governed by a Apache License 2.0
// that can be found in the LICENSE file.
package convert

import (
	"strconv"
)

//字符串转长整型
func Str2Int64(s string) int64 {
	val, _ := strconv.ParseInt(s, 10, 64)
	return val
}

func Str2Uint64(s string) uint64 {
	val, _ := strconv.ParseUint(s, 10, 64)
	return val
}

//字符串转整形
func Str2Int(s string) int {
	return int(Str2Int64(s))
}

func Str2Byte(s string) byte {
	return byte(Str2Int64(s))
}

func Str2Uint16(s string) uint16 {
	return uint16(Str2Uint64(s))
}

//整形转字符串
func Int2str(i int) string {
	return strconv.Itoa(i)
}

//长整型转字符串
func Int642str(i int64) string{
	return strconv.FormatInt(i, 10)
}

func Float642Int(i float64) int {
	return int(Float642Int64(i))
}

func Float642Int64(i float64) int64 {
	return int64(i)
}
