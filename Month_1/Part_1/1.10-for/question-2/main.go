package main

import "fmt"

func main() {
	var a, b int
	fmt.Scan(&a, &b)
	sum := 0
	for ; a <= b; a++ {
		sum = sum + a
	}
	fmt.Println(sum)
}
