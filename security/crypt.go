// Package: security
// File: crypt.go
// Created by mint
// Useage: 密码相关安全工具
// DATE: 14-6-27 23:04
package security

import (
	"crypto/md5"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"crypto/sha1"
	"crypto/cipher"
	"crypto/aes"
	"golang.org/x/crypto/bcrypt"
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

var key = MD5byte("gwsadmin")
var iv = []byte{0x00, 0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09, 0x0a, 0x0b, 0x0c, 0x0d, 0x0e, 0x0f}

//md5加密
func MD5(s string) string {
	return MD5Ex(s)
}

func MD5byte(s string) []byte {
	h := md5.New()
	h.Write([]byte(s))
	return h.Sum(nil)
}

//加盐强密码
func MD5Ex(s string) string {
	h := md5.New()
	h.Write(key)
	h.Write([]byte(s))
	h.Write(iv)
	return fmt.Sprintf("%x", h.Sum(nil))
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

//AES编码
func AesEncode(src []byte) ([]byte, error) {
	var s []byte
	c, err := aes.NewCipher(key)
	if err == nil {
		cfb := cipher.NewCFBEncrypter(c, iv)
		s = make([]byte, len(src))
		cfb.XORKeyStream(s, src)
	}
	return s, err
}

//AES解码
func AesDecode(src []byte) ([]byte, error) {
	var s []byte
	c, err := aes.NewCipher(key)
	if err == nil {
		cfb := cipher.NewCFBDecrypter(c, iv)
		s = make([]byte, len(src))
		cfb.XORKeyStream(s, src)
	}
	return s, err
}
