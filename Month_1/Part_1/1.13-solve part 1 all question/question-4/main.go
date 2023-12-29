package main

import (
	"fmt"
)

func main() {
	var a, b, c int
	fmt.Scanf("%d %d %d", &a, &b, &c)
	if a*a+b*b == c*c {
		fmt.Println("Прямоугольный")
	} else {
		fmt.Println("Непрямоугольный")
	}
}
