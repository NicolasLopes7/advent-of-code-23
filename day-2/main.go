package day2

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

type CubeSet struct {
	Blue  int
	Red   int
	Green int
}
type Game struct {
	Id   int
	sets []CubeSet
}

type Limit struct {
	Blue  int
	Red   int
	Green int
}

func Run() {
	body, err := ioutil.ReadFile("./day-2/input.txt")

	if err != nil {
		log.Fatalf("Error reading file: %v", err)
	}

	lines := strings.Split(string(body), "\n")
	games, err := parseInput(lines)

	if err != nil {
		log.Fatalf("Error parsing input: %v", err)
	}

	// part1(games)
	part2(games)
}

func part1(games []Game) {
	limit := Limit{
		Blue:  14,
		Red:   12,
		Green: 13,
	}

	winnerIds := []int{}

	for _, game := range games {
		isPossible := true
		for _, set := range game.sets {
			if set.Blue > limit.Blue || set.Red > limit.Red || set.Green > limit.Green {
				isPossible = false
			}
		}

		if isPossible {
			winnerIds = append(winnerIds, game.Id)
		}
	}

	idsSum := 0
	for _, id := range winnerIds {
		idsSum += id
	}

	fmt.Println("Winner ids sum: ", idsSum)
}

func part2(games []Game) {
	powers := []int{}

	for _, game := range games {
		fewCubesLimit := Limit{
			Blue:  0,
			Red:   0,
			Green: 0,
		}
		for _, set := range game.sets {
			if set.Blue > fewCubesLimit.Blue {
				fewCubesLimit.Blue = set.Blue
			}

			if set.Red > fewCubesLimit.Red {
				fewCubesLimit.Red = set.Red
			}

			if set.Green > fewCubesLimit.Green {
				fewCubesLimit.Green = set.Green
			}
		}
		powers = append(powers, fewCubesLimit.Blue*fewCubesLimit.Red*fewCubesLimit.Green)
	}

	powersSum := 0
	for _, power := range powers {
		powersSum += power
	}
	fmt.Println("Powers sum: ", powersSum)
}

func parseInput(lines []string) ([]Game, error) {
	var games []Game

	for _, line := range lines {
		game := Game{}

		parts := strings.Split(line, ":")

		// Parse Game ID
		id, err := extractGameID(parts[0])
		if err != nil {
			return nil, fmt.Errorf("error extracting game ID: %v", err)
		}
		game.Id = id

		// Parse Cube Sets
		cubeSets, err := parseCubeSets(parts[1])
		if err != nil {
			return nil, fmt.Errorf("error parsing cube sets: %v", err)
		}
		game.sets = cubeSets

		games = append(games, game)
	}

	return games, nil
}

func extractGameID(part string) (int, error) {
	idStr := strings.Trim(strings.Split(part, "Game")[1], " ")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return 0, fmt.Errorf("error converting id to int: %v", err)
	}
	return id, nil
}

func parseCubeSets(part string) ([]CubeSet, error) {
	var cubeSets []CubeSet

	subsets := strings.Split(part, ";")

	for _, subset := range subsets {
		cubes := strings.Split(subset, ",")
		cubeSet := CubeSet{}

		for _, cube := range cubes {
			cubeParts := strings.Split(cube, " ")
			color := cubeParts[2]

			value, err := strconv.Atoi(cubeParts[1])
			if err != nil {
				return nil, fmt.Errorf("error converting string to int: %v", err)
			}

			switch color {
			case "blue":
				cubeSet.Blue = value
			case "green":
				cubeSet.Green = value
			case "red":
				cubeSet.Red = value
			}
		}

		cubeSets = append(cubeSets, cubeSet)
	}

	return cubeSets, nil
}
