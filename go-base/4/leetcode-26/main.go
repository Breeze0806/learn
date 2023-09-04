package main

import "fmt"

func main() {
	nums := []int{2, 2, 4, 3, 3}
	k := removeDuplicates(nums)
	fmt.Println(k, nums)
}

func removeDuplicates(nums []int) (k int) {
	numMap := make(map[int]struct{})

	k = 0
	for _, v := range nums {
		if _, ok := numMap[v]; !ok {
			numMap[v] = struct{}{}
			nums[k] = v
			k++
		}
	}

	return
}
