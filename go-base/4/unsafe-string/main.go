package main

import (
	"fmt"
	"unsafe"
)

func StringToBytes(s string) []byte {
	return unsafe.Slice(unsafe.StringData(s), len(s))
}

func BytesToString(b []byte) string {
	return unsafe.String(&b[0], len(b))
}

func main() {
	data := StringToBytes("我是水货")
	fmt.Println(data)
	fmt.Println(BytesToString(data))
}
