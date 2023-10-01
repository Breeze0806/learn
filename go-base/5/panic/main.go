package main

import (
	"fmt"
)

func main() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("error:", err)
		}
		fmt.Println("end")
	}()
	panic(nil)
}
