package main

import (
	f "fmt"

	"github.com/bajalnyt/go_oreilly/functions"
	"github.com/bajalnyt/go_oreilly/greetingspackage"
)

// Constant grouping declaration
const (
	// ignore first value by assigning to blank identifier
	_ = iota
	TrafficLightStateRedLight
	TrafficLightStateGreenLight
	TrafficLightStateYellowLight
)

// Constant grouping declaration
const (
	LosAngeles = 1984 + (iota * 4)
	Seoul
	Barcelona
)

func main() {
	f.Println(Barcelona)

	greetingspackage.GopherGreetings()

	f.Println(functions.SumAndDiffOperations(2, 3))
	f.Println(functions.MultiSum(2, 3, 4, 5, 6))

}
