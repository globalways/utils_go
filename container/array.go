// Package: container
// File: array.go
// Created by mint
// Useage: 数组、切片相关
// DATE: 14-6-26 11:26
package container

//元素是否包含在数组中
func InArray(s string, arr []string) bool {
	for _, v := range arr {
		if v == s {
			return true
		}
	}

	return false
}
