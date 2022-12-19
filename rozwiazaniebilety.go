package main

import (
    "fmt"
    "math/rand"
)

func main() {

    // distance in km
    const distanceToMars = 62100000
    // speed in km/s
    const speedMin = 16
    const speedMax = 30
    // a 16 speed price min
    const priceMin = 36
    // a 30 speed price max
    const priceMax = 50

    i := 0

    fmt.Printf("%-17v %-4v %-15v %v \n", "Spaceline", "Days", "Trip type", "Price")
    fmt.Println("==============================================")

    for i < 10 {

        spaceLineId := rand.Intn(3)
        spaceLine := ""
        tripType := "One-way"

        speed := speedMin + rand.Intn(speedMax-speedMin+1)

        // using proportion to calculate price
        price := priceMin + (speed-speedMin)*(priceMax-priceMin)/(speedMax-speedMin)

        // seconds
        tripDuration := distanceToMars / speed
        // days
        tripDuration = tripDuration / (60 * 60 * 24)

        if rand.Intn(2) == 1 {
            tripType = "Round-trip"
            price *= 2
        }

        switch spaceLineId {
        case 0:
            spaceLine = "Space Adventures"
        case 1:
            spaceLine = "SpaceX"
        case 2:
            spaceLine = "Virgin Galactic"
        }

        fmt.Printf("%-17v %-4v %-15v $ %4v \n", spaceLine, tripDuration, tripType, price)

        i++
    }
}
