package main

import "fmt"

func minimumFromFour() int {
    var a, b, c, d int
    fmt.Scan(&a, &b, &c, &d)
    min := a
    if min > b{
        min = b
    }
    if min > c{
        min = c
    }
    if min > d{
        min = d
    }
    return min
}

func main()  {
	fmt.Println(minimumFromFour())
}




