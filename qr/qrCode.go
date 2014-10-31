// Copyright 2014 mint.zhao.chiu@gmail.com. All rights reserved.
// Use of this source code is governed by a Apache License 2.0
// that can be found in the LICENSE file.
package qr

import "io/ioutil"

// 生成二维码
func GenQRCode(text string, level Level) []byte {

	c, err := Encode(text, level)
	if err != nil {
		return nil
	}

	return c.PNG()
}

// 存为png图片
func ToPNGFile(data []byte, file string) bool {
	if err := ioutil.WriteFile(file, data, 0666); err != nil {
		return false
	}

	return true
}
