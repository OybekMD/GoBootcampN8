package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

func main() {
	inputTime, _ := bufio.NewReader(os.Stdin).ReadString('\n')

	parsedTime, err := time.Parse(time.DateTime, inputTime[:19])
	if err != nil {
		fmt.Println(err)
		return 
	}
	if parsedTime.Hour()>13 {
		parsedTime = parsedTime.AddDate(0,0,1)
	}
	fmt.Println(parsedTime.Format(time.DateTime))

}
