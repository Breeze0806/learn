package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(mySqrt(math.MaxInt32))
}

func mySqrt(x int) int {
	a, b := 0, x //多返回值赋值
	for a <= b { // 无；结构的for
		mid := a + (b-a)/2 //二分区间
		if mid*mid == x {  //if语句
			return mid
		}
		if mid*mid < x { //if-else语句
			a = mid + 1 //选择较大的区间
		} else {
			b = mid - 1 //选择较小的区间
		}
	}
	return b
}
