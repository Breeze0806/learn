package main

import "fmt"

func main() {
	s := []int{1, 2, 3, 4}
	var prints []func()
	for _, v := range s {
		prints = append(prints, func() { fmt.Println(v) })
	}

	for _, print := range prints {
		print()
	}

	var ps []*int
	for _, v := range s {
		ps = append(ps, &v)
	}

	for _, v := range ps {
		fmt.Println(*v)
	}

	prints = nil
	for i := 1; i <= 4; i++ {
		prints = append(prints, func() { fmt.Println(i) })
	}

	for _, print := range prints {
		print()
	}
}
