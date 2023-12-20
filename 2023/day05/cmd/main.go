package main

import (
	"fmt"
	"time"

	"github.com/leonasdev/advent-of-code/2023/day05"
)

func trackTime(start time.Time) {
	fmt.Println(time.Since(start))
}

func main() {
	fmt.Println(day05.SolveA("input.txt"))

	defer trackTime(time.Now())
	fmt.Println(day05.SolveB("input.txt"))
}
