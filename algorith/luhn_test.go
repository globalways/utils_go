// Package: algorith
// File: luhn_test.go
// Created by mint
// Useage: luhn验证算法测试
// DATE: 14-7-3 18:53
package algorith

import "testing"

func TestGenerateChkDigit(t *testing.T) {
	card := "632086000000000007"
	if b := GenLuhnCheckDigit([]byte(card)); b != 0 {
		t.Errorf("Error: %v", b)
	}
}

func TestValidateLuhn1(t *testing.T) {
	card := "6320860000000000070"
	if !ValidateLuhn(card) {
		t.Errorf("Error")
	}
}

func TestValidateLuhn2(t *testing.T) {
	card := "6320860000000000061"
	if ValidateLuhn(card) {
		t.Errorf("Error")
	}
}
