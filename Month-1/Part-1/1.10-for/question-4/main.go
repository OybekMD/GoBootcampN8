package main

import (
	"fmt"
)

func main() {
	var sum, temp, max int
	for {
		fmt.Scan(&temp)
		if temp == 0 {
			break
		}
		if max < temp {
			sum = 0
		}
		if temp > max {
			max = temp
		}
		if temp == max {
			sum += 1
		}

	}
	fmt.Println(sum)
}
