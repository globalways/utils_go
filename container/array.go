// Package: container
// File: array.go
// Created by mint
// Useage: 数组、切片相关
// DATE: 14-6-26 11:26
package container

import "reflect"

func Contains(element interface{}, slice interface{}) bool {
	valElement := reflect.ValueOf(element)
	typSlice := reflect.TypeOf(slice)
	valSlice := reflect.ValueOf(slice)

	if !valElement.IsValid() || !valSlice.IsValid() {
		return false
	}

	if typSlice.Kind() != reflect.Slice && typSlice.Kind() != reflect.Array {
		return false
	}

	for idx := 0; idx < valSlice.Len(); idx++ {
		val := valSlice.Index(idx)
		if !val.IsValid() {
			continue
		}

		if val.Interface() == valElement.Interface() {
			return true
		}
	}

	return false
}

//元素是否包含在数组中
func InArray(s string, arr []string) bool {
	for _, v := range arr {
		if v == s {
			return true
		}
	}

	return false
}
