package main

import (
	"fmt"
)

func main() {
	var b, sum int
	fmt.Scan(&b)
	for i := 0; i < b; i++ {
		var temp int
		fmt.Scan(&temp)
		if temp%8 == 0 && (temp > 9 && temp < 100) {
			sum = sum + temp
		}
	}
	fmt.Println(sum)
}
