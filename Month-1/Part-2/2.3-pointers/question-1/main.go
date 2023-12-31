package main

import "fmt"

func test(x1 *int, x2 *int) {
	*x1, *x2 = *x1**x2, *x1**x2
	fmt.Println(*x1)
}

func main() {
	var a, b int
	fmt.Scanf("%d", &a)
	fmt.Scanf("%d", &b)
	test(&a, &b)
}
