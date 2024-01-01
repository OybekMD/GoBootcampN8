package main

import (
	"bufio"
	"io"
	"os"
	"strconv"
)


func main() {
    sum := 0
    s := bufio.NewScanner(os.Stdin)
    for s.Scan() {
        n , err:= strconv.Atoi(string(s.Text()))
        if err != nil{
            break
        }
        sum += n    
    }
    t := strconv.Itoa(sum)
    io.WriteString(os.Stdout , t) 
}