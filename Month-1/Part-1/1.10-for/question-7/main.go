package main

import (
	"fmt"
)

func main() {
	var x, p, y, count int
	fmt.Scan(&x, &p, &y)
	for {
		foiz := float32(p) * 0.01
		x = x + int(float32(x)*foiz)
		count += 1
		if x > y {
			break
		}
	}
	fmt.Println(count)
}
