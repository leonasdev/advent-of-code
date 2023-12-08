package day03

import (
	"bufio"
	"os"
	"strconv"
	"strings"
	"unicode"
)

func isGear(slice []string, x, y int) (bool, [2]int) {
	for i := -1; i <= 1; i++ {
		for j := -1; j <= 1; j++ {
			if i == 0 && j == 0 {
				continue
			}
			if slice[y+i][x+j] == '*' {
				return true, [2]int{y + i, x + j}
			}
		}
	}

	return false, [2]int{}
}

func gearRatios(slice []string) []int {
	res := []int{}
	gearPositionToNumber := map[[2]int]string{}

	for y := 1; y < len(slice)-1; y++ {
		continuousNumber := ""
		isGears := false
		gearPosition := [2]int{}
		for x := 1; x < len(slice[y])-1; x++ {
			isDigit := unicode.IsDigit(rune(slice[y][x]))
			if isDigit {
				continuousNumber += string(rune(slice[y][x]))
				if !isGears {
					isGears, gearPosition = isGear(slice, x, y)
				}
			} else if continuousNumber != "" && isGears {
				if number, exists := gearPositionToNumber[gearPosition]; exists {
					num1, _ := strconv.Atoi(continuousNumber)
					num2, _ := strconv.Atoi(number)

					res = append(res, num1*num2)
				} else {
					gearPositionToNumber[gearPosition] = continuousNumber
				}
				isGears = false
				continuousNumber = ""
			} else if continuousNumber != "" {
				continuousNumber = ""
			}
		}
	}

	return res
}

func SolveB(name string) int {
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
	for _, s := range gearRatios(inputSlice) {
		sum += s
	}

	return sum
}
