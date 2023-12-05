package day01

import (
	"bufio"
	"os"
	"unicode"
)

func CalibrationValue(s string) int {
	value := 0

	for i := 0; i < len(s); i++ {
		if unicode.IsDigit(rune(s[i])) {
			value += int(s[i]-'0') * 10
			break
		}
	}

	for i := len(s) - 1; i >= 0; i-- {
		if unicode.IsDigit(rune(s[i])) {
			value += int(s[i] - '0')
			break
		}
	}

	return value
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
		sum += CalibrationValue(line)
	}

	return sum
}
