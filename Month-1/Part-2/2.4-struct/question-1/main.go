package main

import "fmt"

type Game struct {
	On    bool
	Ammo  int
	Power int
}

func (g *Game) Shoot() bool {

	if g.On == true && g.Ammo > 0 {
		g.Ammo = g.Ammo - 1
		return true
	} else {
		return false
	}
}

func (g *Game) RideBike() bool {

	if g.On == true && g.Power > 0 {
		g.Power = g.Power - 1
		return true
	} else {
		return false
	}
}

func main() {
	testStruct := new(Game)
	fmt.Println(testStruct)
}