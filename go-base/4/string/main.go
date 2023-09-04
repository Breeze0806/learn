package main

import (
	"fmt"
	"unsafe"
)

func main() {
	// 声明一个字符串变量
	str := "Breeze0806"

	// 获取字符串的长度
	fmt.Println(len(str)) // 输出: 10

	// 将字符串转换为字节数组
	bytes := []byte(str)

	// 修改字节数组中的某个元素
	bytes[0] = 'M'

	// 将修改后的字节数组转换回字符串
	str = string(bytes)
	fmt.Println(str) // 输出: Mreeze0806

	//string和[]byte转化
	str = "Breeze0806"
	strPtr := unsafe.Pointer(&str)
	bytesPtr := *(*[]byte)(unsafe.Pointer(uintptr(strPtr)))
	fmt.Println(bytesPtr)

	b := []byte{66, 114, 101, 101, 122, 101, 48, 56, 48, 54}
	str = *(*string)(unsafe.Pointer(&b))
	fmt.Println(str) // Output: Breeze0806
}
