package day05

import (
	"bufio"
	"math"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type Seed struct {
	Start int
	Range int
}

func lastLocationWithRangeSeed(lines []string) int {
	seedStrings := strings.Fields(regexp.MustCompile(`seeds: (.*)`).FindStringSubmatch(lines[0])[1])
	seeds := []Seed{}
	for i := 0; i < len(seedStrings)-1; i += 2 {
		start, _ := strconv.Atoi(seedStrings[i])
		ranges, _ := strconv.Atoi(seedStrings[i+1])
		seed := Seed{
			Start: start,
			Range: ranges,
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

	// locations := []int{}

	out := make(chan int)

	// brute-force with 3 hours lol
	for _, seed := range seeds {
		go func(start int, ranges int) {
			minVal := math.MaxInt
			for i := 0; i < ranges; i++ {
				source := start + i
				source = newSource(mapLines, source)
				// locations = append(locations, source)
				minVal = min(minVal, source)
			}
			out <- minVal
		}(seed.Start, seed.Range)
	}

	minVal := math.MaxInt

	for i := 0; i < len(seeds); i++ {
		minVal = min(minVal, <-out)
	}

	// return slices.Min(locations)
	return minVal
}

func newSource(mapLines [][]string, source int) int {
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

	return source
}

func SolveB(name string) int {
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

	return lastLocationWithRangeSeed(lines)
}
