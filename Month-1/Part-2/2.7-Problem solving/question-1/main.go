package main

import (
    "fmt"
    "math"
       )

func main(){
    var a, b int
    fmt.Scan(&a, &b)
    r := math.Hypot(float64(a), float64(b))
    fmt.Println(r)
}