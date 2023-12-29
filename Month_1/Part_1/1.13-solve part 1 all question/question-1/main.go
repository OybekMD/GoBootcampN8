package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Scanf("%d", &n)
	fmt.Println((n % 10) + (n / 10 % 10) + (n / 100 % 10))
}
