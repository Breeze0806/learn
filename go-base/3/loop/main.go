package main

import "fmt"

func main() {
	var input int
	fmt.Printf("请输入分数(0-5)：")
	fmt.Scanf("%d", &input)

	for i := 0; i < 5; i++ { //正常的for
		fmt.Println("loop1:", i)

		if i < 2 {
			continue //跳过本次执行
		}
		if i == input {
			fmt.Println("loop1: break") //跳出本层循环
			break
		}
	}
	i := 0
Loop:
	for i < 5 { //去掉;的for循环
		fmt.Println("loop2:", i)
		for j := 0; j < 5; j++ {
			if j == input {
				fmt.Println("loop2: break Loop") //跳出Loop标记的循环
				break Loop
			}
			if j == 1 {
				break //跳出本层循环
			}
		}
		i++
	}
	i = 0
	for ; i < 5; i++ { //空缺一个元素并带有；的for循环
		fmt.Println("loop3:", i)
		for j := 0; j < 5; j++ {
			if j < 2 {
				continue
			}
			if j == input {
				goto Exit //跳出到Exit
			}
		}
	}
	return
Exit:
	fmt.Println("loop3: Exit")
}
