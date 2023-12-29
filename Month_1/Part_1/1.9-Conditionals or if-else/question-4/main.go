package main

import "fmt"

func main() {
	var n int32
	fmt.Scan(&n)
	l := (n % 10) + ((n / 10) % 10) + ((n / 100) % 10)
	r := ((n / 1000) % 10) + ((n / 10000) % 10) + ((n / 100000) % 10)
	if r == l {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
