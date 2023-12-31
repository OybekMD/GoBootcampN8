package main

import "fmt"

func main() {
	var n, m, last int
	fmt.Scan(&n, &m)
	//n, m = 100, 101
	var isWorked bool = false
	
	for i := n; i <= m; i++ {
		if i%7 == 0 {
			last = i
			isWorked = true
		}
	}
	if n != m && isWorked {
		fmt.Println(last)
	} else {
		fmt.Println("NO")
	}
}
