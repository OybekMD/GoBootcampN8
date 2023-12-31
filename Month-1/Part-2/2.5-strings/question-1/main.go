package main

import (
	"bufio"
	"fmt"
	"os"
	"unicode"
	"unicode/utf8"
)

func main() {
	text, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	ws := []rune(text)
	if unicode.IsUpper(ws[0]) && string(ws[utf8.RuneCountInString(text)-1]) == "." {
		fmt.Println("Right")
	} else {
		fmt.Println("Wrong")
	}
}
