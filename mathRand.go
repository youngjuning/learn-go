package main

import (
	"fmt"
	"math/rand"
)

func leanMathRand() {
	var num = rand.Intn(10) + 1
	fmt.Println(num)

	num = rand.Intn(10) * 1
	fmt.Println(num)
}
