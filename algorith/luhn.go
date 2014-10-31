// Package: algorith
// File: luhn.go
// Created by mint
// Useage: luhn验证算法
/*
Luhn算法会通过校验码对一串数字进行验证，校验码通常会被加到这串数字的末尾处，从而得到一个完整的身份识别码。

我们以数字“7992739871”为例，计算其校验位：

从校验位开始，从右往左，偶数位乘2（例如，7*2=14），然后将两位数字的个位与十位相加（例如，10：1+0=1，14：1+4=5）；
把得到的数字加在一起（本例中得到67）；
将数字的和取模10（本例中得到7），再用10去减（本例中得到3），得到校验位。

原始数字	    7	9	9	2	7	3	9	8	7	1	x
偶数位乘2	    7	18	9	4	7	6	9	16	7	2	x
将数字相加	7	9	9	4	7	6	9	7	7	2	=67

另一种方法是：
从校验位开始，从右往左，偶数位乘2，然后将两位数字的个位与十位相加；
计算所有数字的和（67）；
乘以9（603）；
取其个位数字（3），得到校验位。
*/
// DATE: 14-7-3 17:41
package algorith

import "strconv"

// 验证是否正确有效
func ValidateLuhn(s string) bool {

	bytes := []byte(s)
	digit, _ := strconv.ParseUint(string(bytes[len(bytes)-1]), 10, 8)

	if GenLuhnCheckDigit(bytes[0:len(bytes)-1]) == byte(digit) {
		return true
	}

	return false
}

func sumDigit(byteDigit []byte) uint8 {

	chkSum := uint8(0)
	bOdd := true
	for i := len(byteDigit) - 1; i >= 0; i-- {
		bit, _ := strconv.ParseUint(string(byteDigit[i]), 10, 8)
		if bOdd {
			chkSum += sumBitsAndTen(uint8(bit) * 2)
		} else {
			chkSum += uint8(bit)
		}

		bOdd = !bOdd
	}

	return chkSum
}

// 参数byteDigit不包含校验位，那么验证时最后那位（循环首位）即为描述中的偶数位
func GenLuhnCheckDigit(byteDigit []byte) byte {

	digit := 10 - (sumDigit(byteDigit) % 10)
	if digit == 10 {
		digit = 0
	}
	return digit
}

func sumBitsAndTen(b uint8) uint8 {
	return (b / 10) + (b % 10)
}
