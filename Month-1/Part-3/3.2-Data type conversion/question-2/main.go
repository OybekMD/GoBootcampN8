package main

import (
	"fmt"
	"strconv"
)

func adding(x, y string) int64 {
	arrx := []rune(x)
	arry := []rune(y)
	resx, resy := 0, 0
	for i := 0; i < len(arrx); i++ {
		n, err := strconv.Atoi(string(arrx[i]))
		if err != nil {
			continue
		}
		resx = resx*10 + n
	}
	for i := 0; i < len(arry); i++ {
		n, err := strconv.Atoi(string(arry[i]))
		if err != nil {
			continue
		}
		resy = resy*10 + n
	}
	return int64(resx + resy)
}

func main() {
	fmt.Println(adding("%^80", "hhhhh20&&&&nd"))
}
