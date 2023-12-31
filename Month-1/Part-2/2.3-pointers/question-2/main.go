package main

import "fmt"

func test(x1 *int, x2 *int) {
	*x1, *x2 = *x2, *x1
	fmt.Println(*x1, *x2)
} 

func main() {
	var a, b int
	fmt.Scanf("%d", &a)
	fmt.Scanf("%d", &b)
	test(&a, &b)
}
