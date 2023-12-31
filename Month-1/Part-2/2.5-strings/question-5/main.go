package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode/utf8"
)

func main() {
	text, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	ws := []rune(text)
	l := utf8.RuneCountInString(text) - 1
	newWord := ""
	for i := 0; i <= l; i++ {
		if strings.Count(text, string(ws[i])) < 2 {
			newWord += string(ws[i])
		}
	}
	fmt.Println(newWord)
}
