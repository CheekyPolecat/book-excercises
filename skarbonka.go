package main

import(
	"fmt"
	"math/rand"
)

func main() {
var monety float32
var skarbonka float32=0


for	skarbonka<20 {
switch coin := rand.Intn(3); coin{
	case 0:
		monety=0.05
	case 1:
		monety=0.1
	case 2:
		monety=0.2
	}
	
	skarbonka=skarbonka+monety


fmt.Printf("Wrzucono %v groszy \n", monety*100)
fmt.Printf("W skarbonce jest %.2f zł \n", skarbonka)
	
}

}
