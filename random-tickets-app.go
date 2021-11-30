package main

import (
	"fmt"
	"math/rand"
)

func main() {
	var distanceFromEarthToMars = 62100000
	var secondsPerDay = 86400
	var spaceLine = ""
	var tripType = ""

	var lenFirstColumn = len("Space Adventures") + 2
	var lenSecondColumn = len("Days") + 2
	var lenThirdColumn = len("Round trip") + 2
	var lenForthColumn = len("Price") + 2
	var fullLength = lenFirstColumn + lenSecondColumn + lenThirdColumn + lenForthColumn

	fmt.Printf("%-0*s | %-0*s | %-0*s | %-0*s |\n", lenFirstColumn, "Spaceline", lenSecondColumn, "Days", lenThirdColumn, "Trip type", lenForthColumn, "Price")
	for printSign := 0; printSign < fullLength+11; printSign++ {
		fmt.Print("=")
	}
	fmt.Print("\n")

	for date := 0; date < 10; date++ {
		var line = rand.Intn(3) + 1
		var types = rand.Intn(2)
		var speed = 16 + 1 + rand.Intn(30-16)
		var price = 20 + speed

		switch line {
		case 1:
			spaceLine = "Space Adventures"
		case 2:
			spaceLine = "SpaceX"
		case 3:
			spaceLine = "Virgin Galactic"
		}

		duration := (distanceFromEarthToMars / speed) / secondsPerDay
		if (distanceFromEarthToMars%speed != 0) || ((distanceFromEarthToMars/speed)%24 != 0) {
			duration += 1
		}

		switch types {
		case 0:
			tripType = "One-way"
		case 1:
			tripType = "Round trip"
			price *= 2
		}
		fmt.Printf("%-0*s | %-0*d | %-0*s | $ %*d |\n", lenFirstColumn, spaceLine, lenSecondColumn, duration, lenThirdColumn, tripType, lenForthColumn-2, price)
	}
}
