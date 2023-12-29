package main

import (
	"fmt"
)

func main() {
	var a, b, c int
	fmt.Scanf("%d %d %d", &a, &b, &c)
	if a+b > c && a+c > b && c+b > a {
		fmt.Println("Существует")
	} else {
		fmt.Println("Не существует")
	}
}
