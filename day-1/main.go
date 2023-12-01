package day1

import (
	"fmt"
	"io/ioutil"
	"log"
	"regexp"
	"strconv"
	"strings"
	"unicode"
)

func Run() {
	body, err := ioutil.ReadFile("./day-1/input.txt")

	if err != nil {
		log.Fatalf("Error reading file: %v", err)
	}

	lines := strings.Split(string(body), "\n")

	// part1(lines)
	part2(lines)
}

func part1(lines []string) {
	sum := 0

	for _, line := range lines {
		first := "0"
		last := "0"

		for i, c := range line {
			isDigit := unicode.IsDigit(c)

			if isDigit {
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

	fmt.Println("Result:", sum)
}

func part2(lines []string) {
	parsedReverse := reverseSlice(lines)

	stringDigits := []string{
		"zero", "one", "two", "three", "four", "five", "six", "seven", "eight", "nine",
	}
	stringDigitsReverse := reverseSlice(stringDigits)

	firstNumREQuery := `\d`
	for _, digit := range stringDigits {
		firstNumREQuery += `|` + digit
	}

	lastNumREQuery := `\d`
	for _, digit := range stringDigitsReverse {
		lastNumREQuery += `|` + digit
	}

	firstNumRE := regexp.MustCompile(firstNumREQuery)
	lastNumRE := regexp.MustCompile(lastNumREQuery)

	var sum = 0

	for i, line := range lines {
		lineReverse := parsedReverse[i]

		first := digitToStr(firstNumRE.FindString(line), stringDigits)
		last := digitToStr(lastNumRE.FindString(lineReverse), stringDigitsReverse)

		calibrationValue, err := strconv.Atoi(first + last)

		if err != nil {
			log.Fatalf("Error: %v", err)
		}

		sum += calibrationValue
	}

	fmt.Println("Result:", sum)
}

func reverseSlice(s []string) (result []string) {
	result = make([]string, len(s))
	for i, v := range s {
		for _, c := range v {
			result[i] = string(c) + result[i]
		}
	}
	return
}

func digitToStr(s string, m []string) string {
	parsedValue, err := strconv.Atoi(s)
	if err == nil {
		return strconv.Itoa(parsedValue)
	}
	for i, value := range m {
		if value == s {
			return strconv.Itoa(i)
		}
	}

	panic("No digit found")
}
