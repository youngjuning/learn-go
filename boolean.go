package main

import (
	"fmt"
	"strings"
)

func learnBoolean() {
	fmt.Println("You find yourself in a dimly lit cavern.")

	var command = "walk outside"
	if strings.Contains(command, "outside") {
		fmt.Println("You leave the cave.")
	}
}
