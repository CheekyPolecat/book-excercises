package main

import (
	"fmt"
)
func main() {
	const distance=236000000000000000
	const lightSpeed=299792
	const secPerYear=86400*365
	const lightYears=distance/lightSpeed/secPerYear

	fmt.Print(lightYears)
}