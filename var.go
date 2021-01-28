package main

import (
	"fmt"
)

func learnVar() {
	var (
		name     = "俊宁"
		age      = 27
		language = "Go"
	)
	fmt.Printf("My name is %v.\n", name)
	fmt.Printf("My age is %v.\n", age)
	fmt.Printf("I am learning %v.\n", language)
}
