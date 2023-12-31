package main

import (
	"fmt"
	"strconv"
)

func main() {
	var str string
	fmt.Scan(&str)
	reply := ""
	for i := 0; i < len(str); i++ {
		num, err := strconv.Atoi(string(str[i]))
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
		wc := num * num
		reply += strconv.Itoa(wc)
	}
	fmt.Println(reply)
}
