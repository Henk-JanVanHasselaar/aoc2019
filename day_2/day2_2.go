package main

import (
	"fmt"
)

func DayTwoTwo() int {
	for noun := 0; noun < 100; noun++ {
		for verb := 0; verb < 100; verb++ {
			result := DayTwoOne(noun, verb)
			if result == 19690720 {
				return 100*noun + verb
			}
		}
	}
	return 0
}

func main() {
	result_1 := DayTwoOne(12, 2)
	fmt.Println("result for day 2 puzzle 1: ", result_1)
	result_2 := DayTwoTwo()
	fmt.Println("result for day 2 puzzle 2: ", result_2)
}
