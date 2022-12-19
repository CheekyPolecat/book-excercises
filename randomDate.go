package main //program losuje losowe daty

import (
	"fmt"
	"math/rand"
)

func main() {

	var era = "AD"

	for numberOfDates := 10; numberOfDates > 0; numberOfDates-- { //funkcja podaje ile losowych dat ma wybrać komputer  numberOfDates - ilosść dat

		year := rand.Intn(2200)
		month := rand.Intn(12) + 1
		daysInMonth := 31

		leap := year%400 == 0 || (year%4 == 0 && year%100 != 0) //sprawdza przestepnosc roku

		switch month {
		case 2:
			if leap {
				daysInMonth = 29
			} else {
				daysInMonth = 28
			}

		case 4, 6, 9, 11:
			daysInMonth = 30
		}

		if leap {
			print("przestepny ")
		} else {
			print("nieprzestepny")
		}

		day := rand.Intn(daysInMonth) + 1
		fmt.Printf("%v %v, miesiac %v, dzien %v\n", era, year, month, day)

	}
}
