package main

import "fmt"

func main() {

	var num int
	var result string
	fmt.Scan(&num)

	if num > 4 && num < 21 {
		result = "korov"
	} else if num%10 == 1 {
		result = "korova"
	} else if num%10 > 1 && num%10 < 5 {
		result = "korovy"
	} else {
		result = "korov"
	}
	fmt.Println(num, result)
}
