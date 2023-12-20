package day05

import (
	"bufio"
	"os"
	"regexp"
	"slices"
	"strconv"
	"strings"
)

func lastLocation(lines []string) int {
	seedStrings := strings.Fields(regexp.MustCompile(`seeds: (.*)`).FindStringSubmatch(lines[0])[1])
	seeds := []int{}
	for _, seedString := range seedStrings {
		seed, err := strconv.Atoi(seedString)
		if err != nil {
			panic(err)
		}
		seeds = append(seeds, seed)
	}

	mapLines := [][]string{}
	currentMapLines := []string{}
	for i := 2; i < len(lines); i++ {
		if regexp.MustCompile(`^[a-z]`).MatchString(lines[i]) {
			currentMapLines = []string{}
			continue
		} else if lines[i] == "" || i == len(lines)-1 {
			mapLines = append(mapLines, currentMapLines)
			currentMapLines = []string{}
			continue
		}
		currentMapLines = append(currentMapLines, lines[i])
	}

	locations := []int{}

	for _, seed := range seeds {
		source := seed
		for _, m := range mapLines {
			for _, line := range m {
				splits := strings.Fields(line)
				destStart, _ := strconv.Atoi(splits[0])
				sourceStart, _ := strconv.Atoi(splits[1])
				ranges, _ := strconv.Atoi(splits[2])

				if source >= sourceStart && source < sourceStart+ranges {
					source = destStart + source - sourceStart
					break
				}
			}
		}
		locations = append(locations, source)
	}

	return slices.Min(locations)
}

func SolveA(name string) int {
	f, err := os.Open(name)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	lines := []string{}
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}

	return lastLocation(lines)
}
