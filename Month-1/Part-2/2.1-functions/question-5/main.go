package main

import "fmt"

func sumInt(nums ...int) (int, int) {
	sum := 0
	for _, num := range nums {
		sum += num
	}
	return len(nums), sum
}

func main() {
	fmt.Println(sumInt(1, 2, 3, 4))
}
