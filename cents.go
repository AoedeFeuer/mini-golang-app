package main

import (
	"fmt"
	"math/rand"
)

func main() {
	var cents = 0

	for cents < 2500 {
		switch rand.Intn(3) {
		case 0:
			cents += 5
		case 1:
			cents += 10
		case 2:
			cents += 25
		}
		fmt.Printf("Balance: %v.%v$\n", cents/100, cents%100)
	}
}
