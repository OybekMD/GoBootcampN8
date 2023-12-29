package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	result := 0
	power := 1
	for n > 0 {
		remainder := n % 2
		n /= 2
		result += remainder * power
		power *= 10
	}
	fmt.Println(result)
}
