package main

import (
	"fmt"
	"time"
)

func work(n int) int {
	if n > 3 {
		time.Sleep(time.Millisecond * 500)
		return n + 1
	} else {
		time.Sleep(time.Millisecond * 500)
		return n - 1
	}
}

func main() {
	var a int
	hashmap := make(map[int]int)

	for i := 0; i < 10; i++ {
		fmt.Scan(&a)
		if el, ok := hashmap[a]; ok {
			fmt.Print(el, " ")
		} else {
			hashmap[a] = work(a)
			fmt.Print(hashmap[a], " ")
		}
	}
}
