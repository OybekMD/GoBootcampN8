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
	isOn := true
	for i := 0; i <= l/2; i++ {
		if string(ws[i]) != string(ws[l-i]) {
			fmt.Println("Нет")
			isOn = false
			break
		}

	}
	if isOn {
		fmt.Println("Палиндром")
	}
}
