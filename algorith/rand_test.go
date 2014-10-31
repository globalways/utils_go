// Package: algorith
// File: rand_test.go
// Created by mint
// Useage: 随机数
// DATE: 14-7-8 17:41
package algorith

import "testing"

func TestRandInt64(t *testing.T) {
	min := int64(0)
	max := int64(999999999999)
	rand := RandomInt64(min, max)
	if rand < min || rand > max {
		t.Errorf("Error:%v", rand)
	}
}
