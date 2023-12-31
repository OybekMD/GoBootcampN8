package main

import (
	"fmt"
)

func main() {
	var word string
	fmt.Scan(&word)
	newst := ""
	for i := 0; i < len(word)-1; i++ {
		newst += string(word[i]) + "*"
	}
	newst += string(word[len(word)-1])
	fmt.Println(newst)
}
