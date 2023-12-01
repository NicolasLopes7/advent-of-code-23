package day1

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

func Run() {
	body, err := ioutil.ReadFile("./day-1/input.txt")

	if err != nil {
		log.Fatalf("Error reading file: %v", err)
	}

	lines := strings.Split(string(body), "\n")

	sum := 0

	for _, line := range lines {
		first := "0"
		last := "0"

		for i, c := range line {
			_, err := strconv.ParseInt(string(c), 10, 64)

			if err == nil {
				if i == 0 || first == "0" {
					first = string(c)
				}

				last = string(c)
			}
		}
		calibrationValue, err := strconv.Atoi(first + last)

		if err != nil {
			log.Fatalf("Error converting string to int: %v", err)
		}

		sum += calibrationValue
	}

	fmt.Println("Result: %v", sum)
}
