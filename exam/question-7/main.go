package main

import "fmt"

func main() {
	// text := "ABC"
	var text string
	fmt.Scan(&text)
	we := []byte(text)
	l := len(we)
	reply := 0
	for i := 0; i < l; i++ {
		reply = reply * 10 + int(we[i] % 64)
	}
    fmt.Println(reply)
}

