package main

import "fmt"

func main() {
	var i int
	fmt.Println("请输入分数：")
	fmt.Scanf("%d\n", &i)
	switch {
	case i < 60: //单个逻辑表达式
		fmt.Println("不合格")
	case i < 70:
		fmt.Println("合格")
	case i < 85:
		fmt.Println("良好")
	default:
		fmt.Println("优秀")
	}

	var c byte
	fmt.Println("请输入等级：")
	fmt.Scanf("%c\n", &c)

	switch c {
	case 'E', 'e': //可以有多个选择
		fmt.Println("1.不合格")
	case 'D', 'd':
		fmt.Println("1.基本合格")
	case 'C', 'c':
		fmt.Println("1.合格")
	case 'B', 'b':
		fmt.Println("1.良好")
	case 'A', 'a':
		fmt.Println("1.优秀")
	default:
		fmt.Println("1.错误的输入")
	}

	switch {
	case c == 'E', c == 'e': //可以有多个表达式
		fmt.Println("2.不合格")
	case c == 'D', c == 'd':
		fmt.Println("2.基本合格")
	case c == 'C', c == 'c':
		fmt.Println("2.合格")
	case c == 'B', c == 'b':
		fmt.Println("2.良好")
	case c == 'A', c == 'a':
		fmt.Println("2.优秀")
	default:
		fmt.Println("2.错误的输入")
	}

	switch {
	case c == 'E':
		fmt.Println("3.不合格")
		fallthrough //fallthrough会执行下一个case区块
	case c == 'e':
		fmt.Println("3.真的不合格")
	case c == 'D', c == 'd':
		fmt.Println("3.基本合格")
	case c == 'C', c == 'c':
		fmt.Println("3.合格")
	case c == 'B', c == 'b':
		fmt.Println("3.良好")
	case c == 'A', c == 'a':
		fmt.Println("3.优秀")
	default:
		fmt.Println("3.错误的输入")
	}

	var in interface{} = i
	switch data := in.(type) { //类型推断
	case int:
		fmt.Printf("int: %v\n", data)
	case uint:
		fmt.Printf("uint: %v\n", data)
	default:
		fmt.Printf("type: %T\n", data)
	}
}
