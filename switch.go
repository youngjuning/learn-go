package main

import "fmt"

func learnSwitch() {
	var command = "go east"

	switch command {
	case "go east":
		fmt.Println("You head further up the mountain.")
	case "go inside":
		fmt.Println("You enter the cave where you live out the rest of your life.")
	default:
		fmt.Println("Didn't quite get that.")
	}
}
