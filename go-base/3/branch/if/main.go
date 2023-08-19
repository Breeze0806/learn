package main

import "fmt"

func main() {
	var input int
	fmt.Println("请输入分数：")
	fmt.Scanf("%d", &input)
	if input < 60 { //if
		fmt.Println("不合格")
	}

	if input < 60 { //if-else
		fmt.Println("不合格")
	} else {
		fmt.Println("合格")
	}

	if input < 60 { //if-else if-else
		fmt.Println("不合格")
	} else if input < 70 {
		fmt.Println("合格")
	} else if input < 85 {
		fmt.Println("良好")
	} else {
		fmt.Println("优秀")
	}
}
