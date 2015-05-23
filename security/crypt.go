// Package: security
// File: crypt.go
// Created by mint
// Useage: 密码相关安全工具
// DATE: 14-6-27 23:04
package security

import (
	"crypto/md5"
	"crypto/sha1"
	"encoding/base64"
	"encoding/hex"
	"golang.org/x/crypto/bcrypt"
	"strings"
)

//密码加密
func GenerateFromPassword(password string) string {
	hashPwd, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		//TODO 如果bcrypt失败，怎么办？
		return ""
	}

	return string(hashPwd)
}

//密码判断，匹配返回true
// 否则false
func CompareHashAndPassword(hashPwd, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashPwd), []byte(password))
	if err != nil {
		return false
	}

	return true
}

//base64加密
//例如:str := utils.Base64Encode([]byte("Hello, playground"))
func Base64Encode(src []byte) string {
	return base64.StdEncoding.EncodeToString(src)
}

//base64解密
func Base64Decode(src string) string {
	code, err := base64.StdEncoding.DecodeString(src)
	if err != nil {
		return ""
	}

	return string(code)
}

//sha1加密
func SHA1(s string) string {
	return hex.EncodeToString(SHA1Byte(s))
}

func SHA1Byte(s string) []byte {
	h := sha1.New()
	h.Write([]byte(s))
	return h.Sum(nil)
}

// md5 16位加密 大写
func Md5Upper16(param, salt string) string {
	h := md5.New()
	h.Write([]byte(param))

	return strings.ToUpper(hex.EncodeToString(h.Sum(nil))[8:24])
}
