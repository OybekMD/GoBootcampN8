package main

import "fmt"

var ctb int = 1
var ctm int = 2
var cts int = 3

type Parking struct {
    sp map[int]int
}

type ParkingSystem struct {
    parking *Parking
}

func Constructor(big int, medium int, small int) ParkingSystem {
    ParkingSystem {
        return ParkingSystem {
            parking: NewParking(big, medium, small)
        }
    }
}

func (p *ParkingSystem) AddCar(carType int) bool {
    t := carType(c)
    if !p.parking.
}

func NewParking(big, medium, small int) *Parking {
    st := make(map[int]int)
}


func (p Parking) Park(c int){

}

func main()  {
    parkingSystem := ParkingSystem{
        Type1: 1,
        Type2: 2,
        Type3: 0,
    }
    obj := Constructor(1,2,3)
    // Cartype -> 1, 2, 3
    param_1 := obj.AddCar(1)
    Constructor(3, 2, 1)
    fmt.Println(parkingSystem)
    fmt.Println(ctb)
	s := []string{"ParkingSystem", "addCar", "addCar", "addCar", "addCar"}


}