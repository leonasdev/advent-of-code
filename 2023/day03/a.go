package day03

import (
	"bufio"
	"os"
	"slices"
	"strconv"
	"strings"
	"unicode"
)

var symbols []rune = []rune{'@', '#', '$', '%', '&', '*', '/', '+', '-', '='}

func isSymbol(r rune) bool {
	return slices.Contains(symbols, r)
}

func matchContinuousNumber(slice []string) []int {
	res := []int{}

	for y := 1; y < len(slice)-1; y++ {
		isPartNumber := false
		continuousNumber := ""
		for x := 1; x < len(slice[y])-1; x++ {
			if unicode.IsDigit(rune(slice[y][x])) {
				continuousNumber += string(rune(slice[y][x]))
				if isSymbol(rune(slice[y-1][x])) || isSymbol(rune(slice[y+1][x])) || isSymbol(rune(slice[y][x-1])) || isSymbol(rune(slice[y][x+1])) || isSymbol(rune(slice[y-1][x-1])) || isSymbol(rune(slice[y+1][x-1])) || isSymbol(rune(slice[y-1][x+1])) || isSymbol(rune(slice[y+1][x+1])) {
					isPartNumber = true
				}
			} else if continuousNumber != "" {
				if isPartNumber {
					num, err := strconv.Atoi(continuousNumber)
					if err != nil {
						panic(err)
					}

					res = append(res, num)
				}
				isPartNumber = false
				continuousNumber = ""
			}
		}
	}

	return res
}

func SolveA(name string) int {
	f, err := os.Open(name)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	inputSlice := []string{}

	flag := true
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		line = ".." + line + ".."
		if flag {
			inputSlice = append(inputSlice, strings.Repeat(".", len(line)))
			flag = false
		}
		inputSlice = append(inputSlice, line)
	}
	inputSlice = append(inputSlice, strings.Repeat(".", len(inputSlice[0])))

	sum := 0
	for _, s := range matchContinuousNumber(inputSlice) {
		sum += s
	}

	return sum
}
