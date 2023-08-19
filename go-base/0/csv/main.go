package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
)

func main() {
	f, err := os.Open("a.csv")
	if err != nil {
		fmt.Println("open a.csv fail. err:", err)
		return
	}
	defer f.Close()
	r := csv.NewReader(f)
	for {
		record, err := r.Read()
		if err == io.EOF {
			return
		}

		if err != nil {
			fmt.Println("read a.csv fail. err:", err)
			return
		}
		fmt.Println(record)
	}
}
