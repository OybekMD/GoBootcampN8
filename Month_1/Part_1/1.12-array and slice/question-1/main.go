package main

import "fmt"

func main() {
	var workArray [10]uint8

	for i := 0; i < 10; i++ {
		fmt.Scanf("%d", &workArray[i])
	}

	for i := 0; i < 3; i++ {
		var firstIndex, secondIndex int
		fmt.Scanf("%d %d", &firstIndex, &secondIndex)

		tmp := workArray[firstIndex]
		workArray[firstIndex] = workArray[secondIndex]
		workArray[secondIndex] = tmp
	}

	for _, element := range workArray {
		fmt.Printf("%d ", element)
	}
}
