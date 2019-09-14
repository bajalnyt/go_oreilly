package main

import (
	"fmt"

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
	fmt.Println(Barcelona)

	greetingspackage.GopherGreetings()
}
