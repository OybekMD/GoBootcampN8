package main

import (
	"fmt"
)

func findWordsContaining(words []string, x byte) []int {
	arr := []int{}
	bigl := len(words)
	for i := 0; i < bigl; i++ {
		l := len(words[i])
		for j := 0; j < l; j++ {
			if words[i][j] == x {
				arr = append(arr, i)
				break
			}
		}
	}
	return arr
}

func main()  {
	var arrstr []string = []string{"leet", "code"}
	var arrstr2 []string = []string{"abc", "bcd", "aaaa", "cbc"}
	var wordbyte byte = 'e'
	var wordbyte2 byte = 'a'
    r := findWordsContaining(arrstr, wordbyte)
    r2 := findWordsContaining(arrstr2, wordbyte2)
	fmt.Println(r)
	fmt.Println(r2)
}