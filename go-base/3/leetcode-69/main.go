package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(mySqrt(math.MaxInt32))
}

func mySqrt(x int) int {
	a, b := 0, x
	for a <= b {
		mid := a + (b-a)/2
		if mid*mid == x {
			return mid
		}
		if mid*mid < x {
			a = mid + 1
		} else {
			b = mid - 1
		}
	}
	return b
}
