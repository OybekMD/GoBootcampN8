package main

import "fmt"

func convert(x int64) uint16 {
	return uint16(x)
}

func main() {
	fmt.Println(convert(123))
}
