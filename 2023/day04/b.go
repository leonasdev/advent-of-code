package day04

import (
	"bufio"
	"os"
	"regexp"
	"strings"
)

func winningCopies(card string, number int) []int {
	copies := []int{}

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
			copies = append(copies, number+count)
		}
	}

	return copies
}

func winningCards(orignalCards []string) int {
	allCards := make([]int, len(orignalCards))
	for i := 0; i < len(orignalCards); i++ {
		allCards[i]++
	}

	for number, times := range allCards {
		copies := winningCopies(orignalCards[number], number)
		for _, number := range copies {
			allCards[number] += times
		}
	}

	total := 0
	for _, v := range allCards {
		total += v
	}

	return total
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
		lines = append(lines, scanner.Text())
	}

	return winningCards(lines)
}
