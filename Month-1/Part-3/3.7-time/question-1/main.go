package main

import (
	"fmt"
	"time"
)

func main() {
	var inputTime string
	fmt.Scan(&inputTime)
	parsedTime, err := time.Parse(time.RFC3339, inputTime)
	if err != nil {
		fmt.Println("Error parsing timestamp:", err)
		return
	}
	unixDateTime := parsedTime.Format(time.UnixDate)
	fmt.Println(unixDateTime)
}
