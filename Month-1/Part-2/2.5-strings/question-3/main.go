package main

import "fmt"

func main() {
	var x, s string
	fmt.Scan(&x, &s)
	isOn := true
	xr := []rune(x)
	for i := 0; i < len([]rune(x)); i++ {
		if string(xr[i:len([]rune(s))+i]) == s {
			fmt.Print(i)
			isOn = false
			break
		}
	}
	if isOn {
		fmt.Print(-1)
	}
}
