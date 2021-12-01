package main

import (
	"fmt"
	"math/rand"
)

func main() {
	var fiveCoin = 0.05
	var tenCoin = 0.1
	var twentyFiveCoin = 0.25
	var moneyBox = 0.0

	for step := 1; moneyBox < 20.0; step++ {
		var pickCoin = rand.Intn(3)

		switch pickCoin {
		case 0:
			moneyBox += fiveCoin
		case 1:
			moneyBox += tenCoin
		case 2:
			moneyBox += twentyFiveCoin
		}
		fmt.Printf("Step: %v, %.2f coins\n", step, moneyBox)
	}
}
