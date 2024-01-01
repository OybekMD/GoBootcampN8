package main

import (
	"fmt"
)

func main() {
	groupCity := map[int][]string{
		10:   []string{"City1", "City2"},
		100:  []string{"City3", "City4"},
		1000: []string{"City5", "City6"},
	}

	cityPopulation := map[string]int{
		"City1": 15,
		"City2": 25,
		"City3": 150,
		"City4": 200,
		"City5": 1200,
		"City6": 1800,
	}

	for _, city := range append(groupCity[10], groupCity[1000]...) {
		delete(cityPopulation, city)
	}

	fmt.Println(cityPopulation)
	fmt.Println(groupCity[100])
}
