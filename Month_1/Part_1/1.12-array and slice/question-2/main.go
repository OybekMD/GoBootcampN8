package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Scan(&n)
	array := make([]int, n, n)
	for i := 0; i < n; i++ {
		var temp int
		fmt.Scan(&temp)
		array[i] = temp
	}
	fmt.Println(array[3])
}
