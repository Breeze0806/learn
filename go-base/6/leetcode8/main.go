package main

import (
	"fmt"
	"math"
)

func myAtoi(s string) int {
	ans := 0
	sign := 0
	for i := range s {
		switch {
		case s[i] == ' ' && sign == 0: //当前空格，并且之前还未出现符号
		case s[i] == '+' && sign == 0: //当前+，并且之前未出现符号
			sign = 1
		case s[i] == '-' && sign == 0: //当前-，并且之前未出现符号
			sign = -1
		case s[i] <= '9' && s[i] >= '0' && sign == 0: //当前数字字符，并且之前未出现符号
			sign = 1
			ans = ans*10 + int(s[i]-'0')
		case s[i] <= '9' && s[i] >= '0' && sign != 0: //当前数字字符，并且之前出现符号
			ans = ans*10 + int(s[i]-'0')
			if sign*ans > math.MaxInt32 { //当前数字字符超过最大值截断
				return math.MaxInt32
			}
			if sign*ans < math.MinInt32 { //当前数字字符小于最大值截断
				return math.MinInt32
			}
		default: //其他情况
			return sign * ans
		}
	}
	return sign * ans
}

func main() {
	fmt.Println(myAtoi("42"))
}
