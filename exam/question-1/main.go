package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// HeLLo -> hEllO
	text, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	ws := []rune(text)
	reply := ""
	for _, v := range ws {
		if v >= 65 && v <= 90 {
			reply = reply + string(v+32)
		} else if v >= 97 && v <= 122 {
			reply = reply + string(v-32)
		} else {
			reply = reply + string(v)
		}
	}
	fmt.Println(reply)
}
