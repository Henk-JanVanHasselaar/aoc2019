package main

import (
	"strconv"
)

func ValidateNumber(number int) bool {
	double := false
	previous := 0
	for _, char := range strconv.Itoa(number) {
		digit, _ := strconv.Atoi(string(char))
		if digit < previous {
			return false
		}
		if digit == previous {
			double = true
		}
		previous = digit

	}

	return double
}

func DayFourOne() int {
	combos := 0
	for i := 171309; i < 643603; i++ {
		if ValidateNumber(i) {
			combos++
		}
	}

	return combos

}
