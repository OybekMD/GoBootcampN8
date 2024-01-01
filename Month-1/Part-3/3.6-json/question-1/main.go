package main

import (
  "encoding/json"
  "fmt"
  "io/ioutil"
  "os"
)

type Student struct {
  LastName   string `json:"LastName"`
  FirstName  string `json:"FirstName"`
  MiddleName string `json:"MiddleName"`
  Birthday   string `json:"Birthday"`
  Address    string `json:"Address"`
  Phone      string `json:"Phone"`
  Rating     []int  `json:"Rating"`
}

type Group struct {
  ID       int       `json:"ID"`
  Number   string    `json:"Number"`
  Year     int       `json:"Year"`
  Students []Student `json:"Students"`
}

func main() {
  data, err := ioutil.ReadAll(os.Stdin)
  if err != nil {
    fmt.Println("Error reading input:", err)
    return
  }
  var group Group
  err = json.Unmarshal(data, &group)
  if err != nil {
    fmt.Println("Error decoding JSON:", err)
    return
  }

  average := calculateAverageRating(group.Students)

  result := map[string]float64{"Average": average}
  resultJSON, err := json.MarshalIndent(result, "", "    ")
  if err != nil {
    fmt.Println("Error encoding result:", err)
    return
  }

  fmt.Println(string(resultJSON))
}

func calculateAverageRating(students []Student) float64 {
  totalRating := 0
  totalCount := 0

  for _, student := range students {
    totalRating += len(student.Rating)
    totalCount++
  }

  if totalCount == 0 {
    return 0.0
  }

  return float64(totalRating) / float64(totalCount)
}