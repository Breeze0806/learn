package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(isPalindrome(math.MaxInt32))
}

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	x64 := x //整形
	px64 := 0
	for x64 != 0 {
		px64 = px64*10 + x64%10
		x64 /= 10
	}
	fmt.Println(px64)
	return px64 == x //布尔型
}

func isPalindromeInt32(x int) bool {
	if x < 0 {
		return false
	}

	x32 := int32(x)
	px32 := int32(x)
	for x32 != 0 {
		px32 = px32*10 + x32%10
		x32 /= 10
	}
	fmt.Println(px32)
	return px32 == int32(x)
}

func isPalindromeInt64(x int) bool {
	if x < 0 {
		return false
	}
	x64 := int64(x)
	px64 := int64(0)
	for x64 != 0 {
		px64 = px64*10 + x64%10
		x64 /= 10
	}
	fmt.Println(px64)
	return px64 == int64(x)
}
