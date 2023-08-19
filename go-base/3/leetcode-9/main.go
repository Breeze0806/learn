package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(isPalindrome(math.MaxInt32))
}

func isPalindrome(x int) bool {
	if x < 0 { //分支结构
		return false
	}
	x64 := x //整形 顺序结构
	px64 := 0
	for x64 != 0 { //循环结构
		px64 = px64*10 + x64%10 // 顺序结构
		x64 /= 10
	}
	return px64 == x //布尔型
}
