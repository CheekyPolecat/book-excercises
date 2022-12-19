package main //Program oblicza wage na Marsie

import (
	"fmt"
)

func main() {
	var (
		age    = 30
		weight = 68.3
	)

	fmt.Print("Moja waga na marsie to: ")
	fmt.Print(weight * 0.3783)
	fmt.Print("kg i mialbym ")
	fmt.Print(age * 365 / 687)
	fmt.Print(" lat")
}
