package day04

import (
	"bufio"
	"math"
	"os"
	"regexp"
	"strings"
)

func winningPoints(card string) int {
	re := regexp.MustCompile(`^Card.*: (.*) \|`)
	winningNumbers := strings.Fields(re.FindStringSubmatch(card)[1])
	winningNumbersSet := map[string]bool{}
	for _, n := range winningNumbers {
		winningNumbersSet[n] = true
	}

	re = regexp.MustCompile(`\| (.*)`)
	myNumbers := strings.Fields(re.FindStringSubmatch(card)[1])

	count := 0
	for _, n := range myNumbers {
		if winningNumbersSet[n] {
			count++
		}
	}

	return int(math.Pow(2, float64(count-1)))
}

func SolveA(name string) int {
	f, err := os.Open(name)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	total := 0

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		total += winningPoints(line)
	}

	return total
}
