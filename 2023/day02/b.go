package day02

import (
	"bufio"
	"os"
	"regexp"
	"strconv"
)

var colors = []string{"red", "green", "blue"}

func powerOfGame(game string) int {
	power := 1

	for _, color := range colors {
		re := regexp.MustCompile(`([0-9]+) ` + color)
		matches := re.FindAllStringSubmatch(game, -1)
		minimalNeed := 0
		for _, match := range matches {
			cube := match[1]
			i, err := strconv.Atoi(cube)
			if err != nil {
				panic(err)
			}

			minimalNeed = max(i, minimalNeed)
		}
		power *= minimalNeed
	}

	return power
}

func SolveB(name string) int {
	f, err := os.Open(name)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	sum := 0

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		sum += powerOfGame(line)
	}

	return sum
}
