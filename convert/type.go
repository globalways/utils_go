// Copyright 2014 mint.zhao.chiu@gmail.com. All rights reserved.
// Use of this source code is governed by a Apache License 2.0
// that can be found in the LICENSE file.
package convert

import (
	"errors"
	"reflect"
	"strconv"
	"strings"
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

func Str2Float64(s string) float64 {
	v, _ := strconv.ParseFloat(s, 64)
	return v
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

func Str2Uint(s string) uint {
	return uint(Str2Uint64(s))
}

//整形转字符串
func Int2str(i int) string {
	return strconv.Itoa(i)
}

//长整型转字符串
func Int642str(i int64) string {
	return strconv.FormatInt(i, 10)
}

// parse struct fields value to map[string]interface{}
// the string's name not struct field name,just struct
// tag name
// e.g:
// type eg struct {
// 	Name string `tag:"parse(name)"`
// }
// return map[string]interface{} {"name": "xxx"}
func ParseStruct(s interface{}, tag, param string) (map[string]interface{}, error) {
	structVal := make(map[string]interface{})

	if s == nil {
		return structVal, nil
	}

	val := reflect.ValueOf(s)
	typ := reflect.TypeOf(s)

	if val.Kind() != reflect.Ptr {
		return nil, errors.New("can not use non-ptr struct.")
	}

	for i := 0; i < typ.Elem().NumField(); i++ {
		if !val.Elem().IsValid() || !val.Elem().Field(i).IsValid() {
			continue
		}

		ftag := typ.Elem().Field(i).Tag.Get(tag)
		paramVal := ""
		if ftag == "" {
			paramVal = snakeString(typ.Elem().Field(i).Name)
		} else {
			for _, v := range strings.Split(ftag, ";") {
				v = strings.TrimSpace(v)
				if n := strings.Index(v, "("); n > 0 && strings.Index(v, ")") == len(v)-1 {
					if v[:n] == param {
						paramVal = v[n+1 : len(v)-1]
						break
					}
				}
			}
		}

		if paramVal == "" {
			continue
		}

		structVal[paramVal] = val.Elem().Field(i).Interface()
	}

	return structVal, nil
}

// snake string, XxYy to xx_yy
func snakeString(s string) string {
	data := make([]byte, 0, len(s)*2)
	j := false
	num := len(s)
	for i := 0; i < num; i++ {
		d := s[i]
		if i > 0 && d >= 'A' && d <= 'Z' && j {
			data = append(data, '_')
		}
		if d != '_' {
			j = true
		}
		data = append(data, d)
	}
	return strings.ToLower(string(data[:len(data)]))
}
