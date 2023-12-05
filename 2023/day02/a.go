package day02

import (
	"bufio"
	"os"
	"regexp"
	"strconv"
)

var maxCubes = map[string]int{
	"red":   12,
	"green": 13,
	"blue":  14,
}

// isPossible returns game id if the game is possible, 0 if not possible
func isPossible(game string) int {
	re := regexp.MustCompile(`^Game (.*):`)
	idStr := re.FindStringSubmatch(game)[1]

	for color, maxCube := range maxCubes {
		re = regexp.MustCompile(`([0-9]+) ` + color)
		matches := re.FindAllStringSubmatch(game, -1)
		for _, match := range matches {
			cube := match[1]
			i, err := strconv.Atoi(cube)
			if err != nil {
				panic(err)
			}

			if i > maxCube {
				return 0
			}
		}
	}

	id, err := strconv.Atoi(idStr)
	if err != nil {
		panic(err)
	}

	return id
}

func SolveA(name string) int {
	f, err := os.Open(name)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	sum := 0

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		sum += isPossible(line)
	}

	return sum
}
