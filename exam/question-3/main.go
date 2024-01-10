package main

import (
	"fmt"
)

func main()  {
    arr := []int{8,11,15,99,1,25}
	max := arr[0]
	l := len(arr)
	for i := 0; i < l; i++ {
		if arr[i] > max {
			max = arr[i]
		}
	}
	fmt.Println(max)
}