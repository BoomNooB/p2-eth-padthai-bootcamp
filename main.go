package main

import (
	"errors"
	"log"
)

// define var in global scope
// k have to remain the same value
var x, y, k float32

func main() {
	// Set initial state
	x = 100
	y = 100
	k = 10000

	log.Printf("Initial state: x=%.4f, y=%.4f, k=%.0f", x, y, k)

	// Trade
	err := trade("x", 10)
	if err != nil {
		log.Println("Trade success")
	}

	log.Printf("After trade: x=%.4f, y=%.4f, k=%.0f", x, y, k)

	// Trade not success
	err = trade("x", 1000)
	if err != nil {
		log.Println("Trade not success")
	}
	log.Printf("After trade: x=%.4f, y=%.4f, k=%.0f", x, y, k)

	// Trade Y asset
	err = trade("y", 99)
	if err != nil {
		log.Println("Trade success")
	}
	log.Printf("After trade: x=%.4f, y=%.4f, k=%.0f", x, y, k)

}

func trade(asset string, amount float32) error {
	// copy value of x and y and check if it's negative
	tmpX := x
	tmpY := y

	if asset == "x" {
		tmpX -= amount
		tmpY = k / tmpX
	}

	if asset == "y" {
		tmpY -= amount
		tmpX = k / tmpY
	}

	if tmpX < 0 || tmpY < 0 {
		log.Println("[Trade] Negative balance")
		return errors.New("Negative balance")
	}

	// update value of x and y
	x = tmpX
	y = tmpY

	log.Println("[Trade] Success")
	return nil
}
