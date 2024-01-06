package main

import "fmt"

func removeDuplicates(nums []int) int {
    l, last, c := len(nums), nums[0], 1
	for i := 1; i < l; i++ {
		if last != nums[i] {
			c++
		}
		last = nums[i]
	}
	return c
}

func main()  {
	a := removeDuplicates([]int{0,0,1,1,1,2,2,3,3,4,6,7,8,6})
	fmt.Println(a)
}