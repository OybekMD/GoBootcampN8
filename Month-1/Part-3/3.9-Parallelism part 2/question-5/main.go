package main

import (
	"fmt"  // пакет используется для проверки выполнения условия задачи, не удаляйте его
	"time" // пакет используется для проверки выполнения условия задачи, не удаляйте его
)

func main() {
	in1 := make(chan int, 10)
	in2 := make(chan int, 10)
	out := make(chan int, 10)
	n := 10
	for i := 0; i < 10; i++ {
		in1 <- i
		in2 <- 10 + i
	}
	merge2Channels(fn, in1, in2, out, n)
	fmt.Println("not blocked if printed first")
	time.Sleep(time.Second * 3)
	//fmt.Println(r)

}

func fn(x int) int {
	time.Sleep(time.Millisecond * 100)
	return x
}

func merge2Channels(fn func(int) int, in1 <-chan int, in2 <-chan int, out chan<- int, n int) {
	go func() {
		for i := range in1 {
			out <- i * 2
		}
	}()
}
