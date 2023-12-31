package main

import "fmt"

func vote(x int, y int, z int) int {
	if x == y {
		return x
	}
	return z
}

func main() {
	fmt.Println(vote(0, 0, 1))
}
