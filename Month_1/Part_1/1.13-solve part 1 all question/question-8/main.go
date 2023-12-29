package main

import "fmt"

func main() {
	var n, min, c, temp int
	c = 1
	fmt.Scan(&n)
	fmt.Scan(&temp)
	min = temp
	for i := 1; i < n; i++ {
		fmt.Scan(&temp)
		if temp < min {
			min = temp
		} else if min == temp {
			c++
		}
	}

	fmt.Println(c)
}
