package day01

import (
	"bufio"
	"os"
	"strings"
	"unicode"
)

var digitMap = map[string]int{
	"one":   1,
	"two":   2,
	"three": 3,
	"four":  4,
	"five":  5,
	"six":   6,
	"seven": 7,
	"eight": 8,
	"nine":  9,
}

func getWordPrefix(s string) int {
	for k, v := range digitMap {
		if strings.HasPrefix(s, k) {
			return v
		}
	}

	return -1
}

func CalibrationValueWithWord(s string) int {
	value := 0

	for i := 0; i < len(s); i++ {
		if unicode.IsDigit(rune(s[i])) {
			value += int(s[i]-'0') * 10
			break
		} else if w := getWordPrefix(s[i:]); w != -1 {
			value += w * 10
			break
		}
	}

	for i := len(s) - 1; i >= 0; i-- {
		if unicode.IsDigit(rune(s[i])) {
			value += int(s[i] - '0')
			break
		} else if w := getWordPrefix(s[i:]); w != -1 {
			value += w
			break
		}
	}

	return value
}

func SolveB(name string) int {
	f, err := os.Open(name)
	if err != nil {
		panic(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)
	sum := 0

	for scanner.Scan() {
		line := scanner.Text()
		sum += CalibrationValueWithWord(line)
	}

	return sum
}
