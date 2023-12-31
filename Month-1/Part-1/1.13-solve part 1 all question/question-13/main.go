package main

import "fmt"


func main() {
	var n int
	fmt.Scan(&n)
	fib1 := 0
	fib2 := 1
	next := 0
	c := 1
	for next < n{
		fib1, fib2 = fib2 , fib1 + fib2
		next = fib2
		c++
	}

	if next == n {
		fmt.Println(c)
	} else {
		fmt.Println(-1)
	}
}