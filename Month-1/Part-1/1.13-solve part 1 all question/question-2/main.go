package main

import (
	"fmt"
)

func main() {
	var n, sum int
	fmt.Scanf("%d", &n)
	for n != 0 {
		sum = sum*10 + n%10
		n /= 10
	}
	fmt.Println(sum)
}
