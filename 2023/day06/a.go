package day06

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func waysToBeat(time int, record int) int {
	count := 0

	for i := 0; i <= time; i++ {
		distance := i * (time - i)
		if distance > record {
			count++
		}
	}

	return count
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
		lines = append(lines, scanner.Text())
	}

	timesStr := strings.Fields(lines[0][6:])
	distancesStr := strings.Fields(lines[1][9:])
	times := []int{}
	distances := []int{}
	for i := range timesStr {
		time, _ := strconv.Atoi(timesStr[i])
		times = append(times, time)
		distance, _ := strconv.Atoi(distancesStr[i])
		distances = append(distances, distance)
	}

	total := 1
	for i := range times {
		total *= waysToBeat(times[i], distances[i])
	}

	return total
}
