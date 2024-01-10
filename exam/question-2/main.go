package main

import (
	"fmt"
)
func printType(data any){
	fmt.Printf("%T\n", data)
}

func main()  {
	var data bool = true
	printType(data)
}