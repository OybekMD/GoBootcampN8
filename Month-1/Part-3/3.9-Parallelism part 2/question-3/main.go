package main

import "fmt"

func calculator(firstChan <-chan int, secondChan <-chan int, stopChan <-chan struct{}) <-chan int {
	channel := make(chan int)

	go func() {
		defer close(channel)
		select {
		case m1 := <-firstChan:
			channel <- (m1 * m1)
		case m2 := <-secondChan:
			channel <- (m2 * 3)
		case <-stopChan:
		}
	}()
	return channel
}

func main() {
	type баста struct{}
	fir, sec := make(chan int), make(chan int)
	stop := make(chan struct{})
	go func() {
		// fir <- 5
		// close(fir)
		sec <- 4
		stop <- баста{}
	}()
	fmt.Print(<-calculator(fir, sec, stop))
}
