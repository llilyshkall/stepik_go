package main

import "fmt"

type Hero struct {
	On    bool
	Ammo  int
	Power int
}

func (h *Hero) RideBike() bool {
	if !h.On {
		return false
	} else if h.Power > 0 {
		h.Power--
		return true
	}
	return false
}

func (h *Hero) Shoot() bool {
	if !h.On {
		return false
	} else if h.Ammo > 0 {
		h.Ammo--
		return true
	}
	return false
}

func main() {
	testStruct := new(Hero)
	fmt.Println(testStruct.RideBike())
	fmt.Println(testStruct.Shoot())
}
