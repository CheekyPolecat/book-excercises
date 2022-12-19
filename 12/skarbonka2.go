package main

import(
	"fmt"
	"math/rand"
)

func main() {
skarbonka:=0 
var monety int
for	skarbonka<2000 {
switch coin := rand.Intn(3); coin{
	case 0:
		monety=5
	case 1:
		monety=10
	case 2:
		monety=20
	}
	skarbonka=skarbonka+monety
	groszy:= skarbonka%100

fmt.Printf("Wrzucono %v groszy \n", monety)
fmt.Printf("W skarbonce jest %d zł %02d groszy \n", skarbonka/100, groszy)
	
}

}
