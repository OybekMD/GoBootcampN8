package main

import (
	"fmt"
	"math"
)

type Sizer interface {
	Area() float64
}

type Square struct {
	width float64
	height float64	
}

func (s Square) Area() float64 {
	return s.width * s.height
}

type Circle struct {
	Radius float64
}

func (c Circle) Area() float64  {
	return math.Pi * math.Pow(c.Radius, 2)
}


func main()  {
	a := Square{
		width: 10,
		height: 10,
	}
	b := Circle{
		Radius: 10,
	}
	l := Less(a, b)
	fmt.Printf("%+v\n",l)


}

func Less(s1, s2 Sizer) Sizer {
	if s1.Area() < s2.Area() {
		return s1
	}
	return s2
}
