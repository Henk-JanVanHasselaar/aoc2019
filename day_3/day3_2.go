package main

import (
	"fmt"
	"strings"
)

func DayThreeTwo() int {
	data := GetInput()
	var cables []Cable
	for _, row := range data {
		cable := Cable{0, 0, 0, make(map[string]int)}
		for _, direction := range strings.Split(string(row), ",") {
			cable.walk(direction)
		}
		cables = append(cables, cable)
	}
	cable1 := cables[0]
	cable2 := cables[1]
	total_steps := 99999999
	for k, v := range cable1.Visited {
		if v2, ok := cable2.Visited[k]; ok {
			steps := v + v2
			if steps < total_steps {
				total_steps = steps
			}
		}
	}
	return total_steps
}

func main() {
	result_1 := DayThreeOne()
	fmt.Println("Day 3 puzzle 1 result: ", result_1)
	result_2 := DayThreeTwo()
	fmt.Println("Day 3 puzzle 2 result: ", result_2)
}
