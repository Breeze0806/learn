package main

import (
	"fmt"
	"math"
	"unsafe"
)

func main() {
	fmt.Printf("int8   length: %v range: %v ~ %v\n", unsafe.Sizeof(int8(1)), math.MinInt8, math.MaxInt8)
	fmt.Printf("int16  length: %v range: %v ~ %v\n", unsafe.Sizeof(int16(1)), math.MinInt16, math.MaxInt16)
	fmt.Printf("int32  length: %v range: %v ~ %v\n", unsafe.Sizeof(int32(1)), math.MinInt32, math.MaxInt32)
	fmt.Printf("int64  length: %v range: %v ~ %v\n", unsafe.Sizeof(int64(1)), math.MinInt64, math.MaxInt64)
	fmt.Printf("int    length: %v\n", unsafe.Sizeof(int(1)))

	fmt.Printf("uint8  length: %v range: 0 ~ %v\n", unsafe.Sizeof(uint8(1)), math.MaxUint8)
	fmt.Printf("uint16 length: %v range: 0 ~ %v\n", unsafe.Sizeof(uint16(1)), math.MaxUint16)
	fmt.Printf("uint32 length: %v range: 0 ~ %v\n", unsafe.Sizeof(uint32(1)), math.MaxUint32)
	fmt.Printf("uint64 length: %v range: 0 ~ %v\n", unsafe.Sizeof(uint64(1)), uint64(math.MaxUint64))
	fmt.Printf("uint   length: %v\n", unsafe.Sizeof(uint(1)))
}
