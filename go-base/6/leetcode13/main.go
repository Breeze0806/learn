package main

import "fmt"

var riMap = map[byte]int{
	'I': 1,
	'V': 5,
	'X': 10,
	'L': 50,
	'C': 100,
	'D': 500,
	'M': 1000,
}

func romanToInt(s string) int {
	ans := 0
	i := 0
	for i = 0; i < len(s)-1; i++ {
		if riMap[s[i]] < riMap[s[i+1]] {
			ans -= riMap[s[i]]
		} else {
			ans += riMap[s[i]]
		}
	}
	ans += riMap[s[i]]
	return ans
}

func main() {
	fmt.Println(romanToInt("MCMXCIV"))
}
