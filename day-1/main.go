package day1

import (
	"fmt"
	"io/ioutil"
	"log"
)

func Run() {
	body, err := ioutil.ReadFile("./day-1/input.txt")

	if err != nil {
		log.Fatalf("Error reading file: %v", err)
	}

	fmt.Println(string(body))
}
