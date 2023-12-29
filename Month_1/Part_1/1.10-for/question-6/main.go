package main

import (
	"fmt"
)

func main() {
	var temp int
	for {
		fmt.Scan(&temp)
		if temp < 10 {
			continue
		}
		if temp > 100 {
			break
		} else {
			fmt.Println(temp)
		}
	}
}
