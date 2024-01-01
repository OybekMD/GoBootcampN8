package main

import (
	"fmt"
	"strconv"
)

func main()  {
	fn := func(num uint) uint {
		strNum := strconv.FormatUint(uint64(num), 10)
		resultStr := ""
		for _, digitChar := range strNum {
			digit, _ := strconv.Atoi(string(digitChar))
			if digit != 0 && digit%2 == 0 {
				resultStr += string(digitChar)
			}
		}
	
		result, _ := strconv.Atoi(resultStr)
		if result == 0 {
			return 100
		}
	
		return uint(result)
	}
	fmt.Println(fn(123))	
}