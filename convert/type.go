// Copyright 2014 mint.zhao.chiu@gmail.com. All rights reserved.
// Use of this source code is governed by a Apache License 2.0
// that can be found in the LICENSE file.
package convert

import "strconv"

//字符串转长整型
func Str2int64(s string) (int64, error) {
	return strconv.ParseInt(s, 10, 64)
}

//字符串转整形
func Str2int(s string) (int, error) {
	return strconv.Atoi(s)
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
	return int(i)
}
