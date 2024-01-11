package main

import "fmt"

func main() {
	// Rim to int
    var value, last int
    data := map[byte]int {
        'I' : 1,
        'V' : 5,
        'X' : 10,
        'L' : 50,
        'C' : 100,
        'D' : 500,
        'M' : 1000,
    }
    var strdata string
	fmt.Scan(&strdata)
	l := len(strdata)
    for i := l-1; i>= 0;i-- {
        if data[strdata[i]] < last {
            value -= data[strdata[i]]
        }
        if data[strdata[i]] >= last {
            value += data[strdata[i]]
        }
        last = data[strdata[i]]
    }
    fmt.Println(value)
}