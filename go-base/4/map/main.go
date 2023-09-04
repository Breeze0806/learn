package main

import "fmt"

func main() {
	// 创建一个映射
	m := make(map[string]int)
	//赋值
	m = map[string]int{
		"我是水货": 1,
		"水货":   2,
		"水":    3,
	}
	// 添加键值对到映射
	m["货"] = 4

	// 访问映射中的元素
	fmt.Println(m["我是水货"]) // 输出: 1
	fmt.Println(m["水货"])   // 输出: 2
	fmt.Println(m["水"])    // 输出: 3

	// 修改映射中的元素
	m["水货"] = 10
	fmt.Println(m["水货"]) // 输出: 10

	// 删除映射中的元素
	delete(m, "水")
	fmt.Println(m["水"]) // 输出: 0

	// 检查映射中是否存在某个键
	_, ok := m["水货"]
	fmt.Println(ok) // 输出: true
	_, ok = m["水"]
	fmt.Println(ok) // 输出: false
}
