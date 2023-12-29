package main

import "fmt"

func main() {
	var a int32
	fmt.Scan(&a)
	if (a%10 == (a/10)%10) || (a%10 == a/100) || ((a/10)%10 == a/100) {
		fmt.Println("NO")
	} else {
		fmt.Println("YES")
	}
}
