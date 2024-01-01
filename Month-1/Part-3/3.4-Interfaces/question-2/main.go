package main

import (
  "fmt"
)

type Battery struct {
  capacity string
}

func (b Battery) String() string {
  formattedCapacity := ""
  cx := 0
  for _, v := range b.capacity {
    if v == '0' {
      formattedCapacity += " "
    } else {
      cx++

    }
  }
  for i := 0; i < cx; i++ {
    formattedCapacity += "X"
  }
  return "[" + formattedCapacity + "]"
}

func main() {
  var input string
  fmt.Scan(&input)

  batteryForTest := Battery{capacity: input}
  fmt.Print(batteryForTest)
}