// Package: ipv4
// File: ipv4.go
// Created by mint
// Useage: ip相关工具类
// DATE: 14-6-26 9:51
package ipv4

import (
	"net"
	"strings"
)

//获取客户端IP
func GetClientIP() (string, error) {
	conn, err := net.Dial("udp", "google.com:80")
	if err != nil {
		return "", err
	}
	defer conn.Close()

	return strings.Split(conn.LocalAddr().String(), ":")[0], nil
}
