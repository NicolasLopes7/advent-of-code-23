package main

import (
	"log"
	"os"

	day1 "github.com/nicolaslopes7/advent-of-code-23/day-1"
)

func main() {
	args := os.Args

	if len(args) < 2 {
		log.Fatal("Please provide a valid day arg\nEx: day-1")
	}

	day := args[1]

	switch day {
	case "day-1":
		day1.Run()
	default:
		log.Fatal("Please provide a valid day arg. Ex: day-1")
	}
}
