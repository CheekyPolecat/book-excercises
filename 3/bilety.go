package main // Program losuje 10 biletow lotniczych na marsa

import (
	"fmt"
	"math/rand"
	"time"
)

const distance = 62100100

func main() { //komentarz

	//losuje predkosc statku od 16 do 30 km/s i czas lotu

	const (
		speedMin = 16
		speedMax = 30
		cenaMin  = 36
		cenaMax  = 50
	)
	var ticketPrice int
	var ticketType string
	var linia string

	rand.Seed(time.Now().UnixNano())

	fmt.Printf("%-15v %-4v %-17v %-5v\n", "Linia", "Dni", "Rodzaj biletu", "Cena")
	fmt.Println("=============================================")

	
	for i := 10; i > 0; i-- {

		speed := rand.Intn(speedMax-speedMin+1) + speedMin
		travelTime := distance / speed / (60 * 60 * 24) //czas lotu

		//Losuje linie lotnicza

		switch al := rand.Intn(3); al {
		case 0:
			linia = "Virgin Atlantic"
		case 1:
			linia = "SpaceX"
		case 2:
			linia = "Nasa"
		}

		//wybiera cene biletu
		if speed >= 16 && speed < 20 {
			ticketPrice = 36
		} else if speed >= 20 && speed <= 25 {
			ticketPrice = 42
		} else {
			ticketPrice = 50
		}

		//Losuje typ biletu
		switch tt := rand.Intn(2); tt {
		case 0:
			ticketType = "w jedną strone"
		case 1:
			ticketType = "Powrotny"
			ticketPrice *= 2
		}

		fmt.Printf("%-15v %-4v %-17v $ %4v\n", linia, travelTime, ticketType, ticketPrice)
	}

}
