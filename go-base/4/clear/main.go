package main

import "fmt"

func main() {
	myMap := map[string]int{
		"A": 1,
		"B": 2,
		"C": 3,
	}

	fmt.Println(myMap)
	for k := range myMap {
		delete(myMap, k)
	}

	fmt.Println(myMap)

	myMap = map[string]int{
		"A": 1,
		"B": 2,
		"C": 3,
	}
	fmt.Println(myMap)

	clear(myMap)
	fmt.Println(myMap)

	mySlice := []int{1, 2, 3, 4, 5, 6, 7, 8}
	fmt.Println(mySlice)

	clear(mySlice)
	fmt.Println(mySlice)

}
