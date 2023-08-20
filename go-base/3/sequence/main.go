package main

import "fmt"

func main() {
	var i1 int     //声明变量,默认值为0
	var i2 int = 1 //声明变量,初始化为1
	i3 := 2        //这是最常用的声明和初始化手段
	fmt.Println(i1, i2, i3)
	i4 := i2 + i3 //使用运算表达式赋值
	i3 *= i4      //使用运算表达式赋值
	fmt.Println(i3, i4)
	const ci1 int = 13 //声明常量，无法变更
	fmt.Println(ci1)

	x, y, z := 100, 101, 102 //多返回值赋值
	fmt.Println(x, y, z)
}
