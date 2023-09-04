package main

import "fmt"

func main() {
	// 创建一个整数数组
	a := [5]int{1, 2, 3, 4, 5}

	// 创建一个基于数组的切片
	s := a[1:3]
	fmt.Println(s) // 输出: [2 3]

	// 使用make函数创建一个长度为5的切片，初始值为0
	m := make([]int, 5)
	fmt.Println(m) // 输出: [0 0 0 0 0]

	// 使用append函数向切片添加元素
	m = append(m, 1, 2, 3)
	fmt.Println(m) // 输出: [0 0 0 0 0 1 2 3]

	// 使用copy函数复制切片
	n := make([]int, len(m))
	copy(n, m)
	fmt.Println(n) // 输出: [0 0 0 0 0 1 2 3]

	// 使用delete函数删除切片元素
	m = append(m[:len(m)-2], m[2:]...)
	fmt.Println(m) // 输出: [0 0 3]
}
