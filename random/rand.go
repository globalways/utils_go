// Package: algorith
// File: rand.go
// Created by mint
// Useage: 随机数工具
// DATE: 14-6-27 23:30
package random

import (
	"bytes"
	"time"
	"github.com/globalways/utils_go/convert"
	"math/rand"
)

//生成随机字符串
func RandomString(num int) string {
	var result bytes.Buffer
	var temp string
	for i := 0; i < num; {
		if string(RandomInt64(65, 90)) != temp {
			temp = string(RandomInt64(65, 90))
			result.WriteString(temp)
			i++
		}
	}
	return result.String()
}

//生成随机数字
func RandomInt64(min, max int64) int64 {
	rand.Seed(time.Now().UTC().UnixNano())
	return min + rand.Int63n(max-min)
}

func RandomInt(min, max int) int {
	return int(RandomInt64(int64(min), int64(max)))
}

func RandomUint(min, max uint) uint {
	return uint(RandomInt64(int64(min), int64(max)))
}

//随机数字字符串
func RandInt64Str(min, max int64) string {
	i := RandomInt64(min, max)
	return convert.Int642str(i)
}

func RandIntStr(min, max int) string {
	i := RandomInt(min, max)
	return convert.Int2str(i)
}
