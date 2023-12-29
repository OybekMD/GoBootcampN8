package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Scanf("%d", &n)
	fmt.Printf("It is %d hours %d minutes.", n/3600, n%3600/60)
}
