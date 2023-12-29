package main

import "fmt"

func main() {
	var n, m, sum int
	fmt.Scan(&n)
	m = n
	for n > 10 {
		for m != 0 {
			sum = sum + (m % 10)
			m = m / 10
		}
		n = sum
		m = n
		sum = 0
	}
	fmt.Println(n)
}
