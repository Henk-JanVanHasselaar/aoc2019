package main

import (
	"fmt"
	"strconv"
)

func ValidateNumberTwo(number int) bool {
	double := make(map[int]int)
	previous := 0
	for _, char := range strconv.Itoa(number) {
		digit, _ := strconv.Atoi(string(char))
		if digit < previous {
			return false
		}
		if digit == previous {
			if _, ok := double[digit]; ok {
				double[digit]++
			} else {
				double[digit] = 1
			}
		}
		previous = digit
	}
	for _, v := range double {
		if v == 1 {
			return true
		}
	}
	return false

}

func DayFourTwo() int {
	combos := 0
	for i := 171309; i < 643603; i++ {
		if ValidateNumberTwo(i) {
			combos++
		}
	}

	return combos

}

func main() {
	result_1 := DayFourOne()
	fmt.Println("Day 4 puzzle one result: ", result_1)
	result_2 := DayFourTwo()
	fmt.Println("Day 4 puzzle two result: ", result_2)
}
