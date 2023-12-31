package main

import (
	"fmt"
	"unicode"
)

func isAlphanumeric(c rune) bool {
	return unicode.Is(unicode.Latin, c) || unicode.IsDigit(c)
}

func checkPassword(password string) string {
	if len(password) < 5 {
		return "Wrong password"
	}

	for _, char := range password {
		if !isAlphanumeric(char) {
			return "Wrong password"
		}
	}

	return "Ok"
}

func main() {
	var pw string
	fmt.Scan(&pw)
	result := checkPassword(pw)
	fmt.Println(result)
}
