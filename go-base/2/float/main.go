package main

import (
	"fmt"
	"math"
	"unsafe"
)

func main() {
	fmt.Printf("float32 length: %v range: %v ~ %v\n", unsafe.Sizeof(float32(1)), -math.MaxFloat32, math.MaxFloat32)
	fmt.Printf("float64 length: %v range: %v ~ %v\n", unsafe.Sizeof(float64(1)), -math.MaxFloat64, math.MaxFloat64)
	fmt.Printf("before transfer float32: %v\n", float32(1.328))
	fmt.Printf("after  transfer float64: %v\n", float64(float32(1.328)))
}
