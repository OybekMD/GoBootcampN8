package main

import (
	"fmt"
	"strconv"
)

func main() {
	var str string
	fmt.Scanln(&str)

	maxDigit := 0

	for _, char := range str {
		digit, err := strconv.Atoi(string(char))
		if err != nil {
			continue
		}

		if digit > maxDigit {
			maxDigit = digit
		}
	}

	fmt.Println(maxDigit)
}
