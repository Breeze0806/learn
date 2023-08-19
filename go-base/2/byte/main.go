package main

import (
	"fmt"
	"unsafe"
)

func main() {
	fmt.Printf("byte length: %v %v %c\n", unsafe.Sizeof(byte('a')), byte('a'), byte('a'))
	fmt.Printf("rune length: %v %v %c\n", unsafe.Sizeof(rune('中')), rune('中'), rune('中'))
}
