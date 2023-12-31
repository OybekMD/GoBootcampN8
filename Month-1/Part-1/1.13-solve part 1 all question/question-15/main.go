package main

import "fmt"

func main() {
	var n, m int
	fmt.Scan(&n, &m)
	nlen := 0
	tempn := n
	for tempn != 0 {
		nlen++
		tempn /= 10
	}
	reply := 0
	tendiv := 1
	for i := 1; i < nlen; i++ {
		tendiv *= 10
	}
	for i := 0; i < nlen; i++ {
		if n/tendiv != m {
			reply = reply*10 + (n / tendiv)
		}
		n = n % tendiv
		tendiv = tendiv / 10
	}
	fmt.Println(reply)
}
