package main

import (
	"fmt"
)

func main() {
	var a, b float64
	fmt.Scanf("%f %f", &a, &b)
	mean := (a + b) / 2.0
	var mean2 int64 = int64(mean)
	if mean < float64(mean2) || mean > float64(mean2) {
		fmt.Printf("%.1f\n", mean)
	} else {
		fmt.Printf("%.f\n", mean)
	}
}
