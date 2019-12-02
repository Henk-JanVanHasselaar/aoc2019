package main

import (
	"fmt"
	"strconv"
	"strings"
)

func CalcFuelCost(input_num int) int {
	// 8 or less is no fuel cost. (would end up negative)
	if input_num <= 8 {
		return 0
	}
	result := input_num / 3
	// parse float to string, so we can split on "."
	str_result := FloatToString(float64(result))
	res := strings.Split(str_result, ".")[0]
	i1_int, _ := strconv.Atoi(res)
	i1 := i1_int - 2
	if i1 <= 0 {
		return input_num
	} else {
		i1 += CalcFuelCost(i1)
	}

	return i1
}

func DayOneTwo() int {
	numbers := readFile("input/input.txt")

	final := 0
	for _, item := range numbers {
		final += CalcFuelCost(item)
	}
	return final
}

func main() {
	day_1_1 := DayOneOne()
	fmt.Println("day one first puzzle result: ", day_1_1)
	day_1_2 := DayOneTwo()
	fmt.Println("day one second puzzle result: ", day_1_2)
}
