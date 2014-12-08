// Package: algorith
// File: rand.go
// Created by mint
// Useage: 随机数工具
// DATE: 14-6-27 23:30
package algorith

import (
	"bytes"
	"crypto/rand"
	r "math/rand"
	"time"
	"strconv"
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
	r.Seed(time.Now().UTC().UnixNano())
	return min + r.Int63n(max-min)
}

//随机数字字符串
func RandInt2Str(min, max int64) string {
	i := RandomInt64(min, max)
	return strconv.Itoa(int(i))
}

// RandomCreateBytes generate random []byte by specify chars.
func RandomCreateBytes(n int, alphabets ...byte) []byte {
	const alphanum = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
	var bytes = make([]byte, n)
	var randby bool
	if num, err := rand.Read(bytes); num != n || err != nil {
		r.Seed(time.Now().UnixNano())
		randby = true
	}
	for i, b := range bytes {
		if len(alphabets) == 0 {
			if randby {
				bytes[i] = alphanum[r.Intn(len(alphanum))]
			} else {
				bytes[i] = alphanum[b%byte(len(alphanum))]
			}
		} else {
			if randby {
				bytes[i] = alphabets[r.Intn(len(alphabets))]
			} else {
				bytes[i] = alphabets[b%byte(len(alphabets))]
			}
		}
	}
	return bytes
}
