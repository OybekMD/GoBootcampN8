package main

import (
	"fmt"
	"strconv"
	"time"
)

func task(ch chan int, N int) {
    ch <- N + 1
}

func main() {
	theMine := [3]int{1, 2, 3}
	oreChan := make(chan int)
	go func(mine [3]int) {
	   for _, item := range mine {
		  task(oreChan, item)
	   }
	}(theMine)
	go func() {
	   for i := 0; i < 3; i++ {
		  foundOre := <-oreChan
		  fmt.Println("Miner: Received " + strconv.Itoa(foundOre) + " from finder")
	   }
	}()
	<-time.After(time.Second * 5)
 }
