package main

import (
	"fmt"
	"unsafe"
)

func main() {
	fmt.Printf("bool length: %v %v/%v", unsafe.Sizeof(true), 0 == 0, 0 != 0)
}
