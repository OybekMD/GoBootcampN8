package main

import (
	"fmt"
)

func numLenght(x int) int {
	l := 0
	for x > 0 {
		x = x / 10
		l++
	}
	return l
}

func tenMulti(x int, y int) int {
	for i := 1; i < y; i++ {
		x = x / 10
	}
	return x % 10
}

func main() {
	var n, m int
	fmt.Scan(&n, &m)
	nl := numLenght(n)
	ml := numLenght(m)
	for i := nl; i >= 1; i-- {
		for j := ml; j >= 1; j-- {
			if tenMulti(n, i) == tenMulti(m, j) {
				fmt.Print(tenMulti(n, i), " ")
			}
		}
	}
}
