package main

import "fmt"

func calculator(arguments <-chan int, done <-chan struct{}) <-chan int {
    channel:=make(chan int)
    sum := 0
    
    go func() {
        defer close(channel)
        for {
            select{
            case n:= <-arguments:
                sum += n
            case <-done:
                channel <-sum
                return
            }
        }       
    }()
    return channel
}

func main(){
    arguments := make(chan int)
    done := make(chan struct{})
    result := calculator(arguments, done)
    for i := 0; i < 10; i++ {
       arguments <- 1
    }
    close(done)
    fmt.Println(<-result)
 }