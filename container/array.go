// Package: container
// File: array.go
// Created by mint
// Useage: 数组、切片相关
// DATE: 14-6-26 11:26
package container

import "reflect"

func Contains(element interface{}, target interface{}) bool {
	valElement := reflect.ValueOf(element)
	typSlice := reflect.TypeOf(target)
	valSlice := reflect.ValueOf(target)

	if !valElement.IsValid() || !valSlice.IsValid() {
		return false
	}

	switch typSlice.Kind() {
	case reflect.Slice, reflect.Array:
		for idx := 0; idx < valSlice.Len(); idx++ {
			val := valSlice.Index(idx)
			if !val.IsValid() {
				continue
			}

			if val.Interface() == valElement.Interface() {
				return true
			}
		}
	case reflect.Map:
		if valSlice.MapIndex(valElement).IsValid() {
			return true
		}
	}

	return false
}

func Delete(element interface{}, slice interface{}) bool {
	valElement := reflect.ValueOf(element)
	typSlice := reflect.TypeOf(slice)
	valSlice := reflect.ValueOf(slice)

	if !valElement.IsValid() || !valSlice.IsValid() {
		return false
	}

	switch typSlice.Kind() {
	case reflect.Slice:
		sliceLen := valSlice.Len()
		for idx := 0; idx < sliceLen; idx++ {
			val := valSlice.Index(idx)
			if !val.IsValid() {
				continue
			}

			if val.Interface() == valElement.Interface() {
				if idx == sliceLen-1 {
					valSlice = valSlice.Slice(0, idx)
				} else {
					valSlice = reflect.AppendSlice(valSlice.Slice(0, idx), valSlice.Slice(idx+1, sliceLen-1))
				}
			}
		}
	case reflect.Map:
	}

	return false
}

func Index(element interface{}, target interface{}) int {
	valElement := reflect.ValueOf(element)
	typSlice := reflect.TypeOf(target)
	valSlice := reflect.ValueOf(target)

	if !valElement.IsValid() || !valSlice.IsValid() {
		return -1
	}

	switch typSlice.Kind() {
	case reflect.Slice, reflect.Array:
		for idx := 0; idx < valSlice.Len(); idx++ {
			val := valSlice.Index(idx)
			if !val.IsValid() {
				continue
			}

			if val.Interface() == valElement.Interface() {
				return idx
			}
		}
	}

	return -1
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
