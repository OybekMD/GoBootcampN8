package main

import "fmt"

func main() {
	var n int32
	fmt.Scan(&n)
	if n < 10 {
		fmt.Println(n)
	} else if n <= 99 {
		fmt.Println((n / 10) % 10)
	} else if n <= 999 {
		fmt.Println((n / 100) % 10)
	} else if n <= 9999 {
		fmt.Println((n / 1000) % 10)
	}
	if n == 10000 {
		fmt.Println(1)
	}
}
