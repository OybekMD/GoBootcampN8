package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	for i := 1; i <= n; i = i + i {
		fmt.Print(i, " ")
	}

}
