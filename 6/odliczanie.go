package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	var count=0
		
	for 
	 {
		fmt.Println(count)
		time.Sleep(time.Second)
		count++
		
		if count>=10 {
		count=0
		
			if rand.Intn(2)==0 {
				break}
		}
		
		}
	
	}

 