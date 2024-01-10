package main

import (
	"fmt"
)

func main() {
	arr := []int{1, 2, 3, 4, 6, 7, 8}
	l := len(arr)
	start := arr[0]
	for i := 0; i < l; i++ {
		if start != arr[i] {
			fmt.Println(start)
			break
		}
		start++
	}
}
