package main

import "fmt"

func task2(ch chan string, str string) {
    for i := 0; i < 5; i++ {
        ch <- str + " "
    }

	close(ch)
}


func main() {
    ch := make(chan string)

    go task2(ch, "hello")

    for message := range ch {
        fmt.Println(message)
    }
}


