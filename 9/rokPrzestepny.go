package main  			// program okresla  czy rok jest przestepny

import (
	"fmt"
	"math/rand"
)

func main() {
	var year = rand.Intn(2100) + 1 	//losowo losuje rok
	var przestepnosc = year%400 == 0 || (year%4 == 0 && year%100 != 0)

	fmt.Printf("Rok %v jest", year)
	if przestepnosc {
		fmt.Print(" przestepny")
	} else {
		fmt.Print(" nieprzestepny")
	}
}
