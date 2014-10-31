// Package: container
// File: type.go
// Created by mint
// Useage: go类型相关
// DATE: 14-6-27 23:32
package container

import "reflect"

//是否Map类型
func IsMap(v interface{}) bool {
	return reflect.ValueOf(&v).Elem().Elem().Kind() == reflect.Map
}
