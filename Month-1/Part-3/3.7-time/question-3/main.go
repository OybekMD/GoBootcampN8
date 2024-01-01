package main

import (
	"bufio"
	"fmt"
	"strings"
	"os"
	"time"
)

func main() {
	inputTime, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	s1, s2 := strings.Split(inputTime, ",")[0], strings.Split(inputTime, ",")[1]

	s1time, err := time.Parse("02.01.2006 15:04:05", s1)
	if err != nil {
		return 
	}
	s2time, err := time.Parse("02.01.2006 15:04:05", s2)
	if err != nil {
		return 
	}
	
	diff := time.Since(s1time).Round(time.Second) - time.Since(s2time).Round(time.Second)
	diff2 := time.Since(s2time).Round(time.Second) - time.Since(s1time).Round(time.Second)

	if diff < diff2 {
		fmt.Println(diff2)
	} else {
		fmt.Println(diff)
	}

}
