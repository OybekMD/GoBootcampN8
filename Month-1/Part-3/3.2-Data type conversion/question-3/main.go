package main

import (
    "bufio"
	"os"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	text, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	splitString := strings.Split(text, ";")
	splitString[0] = strings.ReplaceAll(splitString[0], ",", ".")
	splitString[0] = strings.ReplaceAll(splitString[0], " ", "")
	splitString[1] = strings.ReplaceAll(splitString[1], ",", ".")
	splitString[1] = strings.ReplaceAll(splitString[1], " ", "")

	numoneFloat, err1 := strconv.ParseFloat(splitString[0], 64)
	numtwoFloat, err2 := strconv.ParseFloat(splitString[1], 64)
	if err1 != nil || err2 != nil {
		fmt.Println("Error parsing number:", err1, err2)
	} else {
		fmt.Printf("%.4f", numoneFloat/numtwoFloat)
	}
}
