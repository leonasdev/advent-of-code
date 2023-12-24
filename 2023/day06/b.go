package day06

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

// func waysToBeat(time int, record int) int {
// 	count := 0
//
// 	for i := 0; i <= time; i++ {
// 		distance := i * (time - i)
// 		if distance > record {
// 			count++
// 		}
// 	}
//
// 	return count
// }

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

	timesStr := strings.Join(strings.Fields(lines[0][6:]), "")
	distancesStr := strings.Join(strings.Fields(lines[1][9:]), "")
	time, _ := strconv.Atoi(timesStr)
	distance, _ := strconv.Atoi(distancesStr)

	return waysToBeat(time, distance)
}
