package main

import (
	"bufio"
	"fmt"
	"os"
	"unicode/utf8"
)

func main() {
	text, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	ws := []rune(text)
	l := utf8.RuneCountInString(text) - 1
	newWord := ""
	for i := 1; i <= l; i = i + 2 {
		newWord += string(ws[i])
	}
	fmt.Println(newWord)
}
