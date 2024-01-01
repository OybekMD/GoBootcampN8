package main

import "fmt"

func removeDuplicates(inputStream <-chan string, outputStream chan<- string) {
    var prev string
    for value := range inputStream {
        if value != prev {
            outputStream <- value
            prev = value
        }
    }
    close(outputStream)
}

func main()  {
	inputStream := make(chan string)
    outputStream := make(chan string)

    go func() {
        inputStream <- "apple"
        inputStream <- "banana"
        inputStream <- "apple"
        inputStream <- "orange"
        inputStream <- "banana"
        close(inputStream)
    }()

    go removeDuplicates(inputStream, outputStream)

    for value := range outputStream {
        fmt.Println(value) // Output: apple banana orange
    }
}