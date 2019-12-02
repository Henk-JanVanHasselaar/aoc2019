package main

import (
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

func GetInput() (numbers []int) {
	content, err := ioutil.ReadFile("input/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	for _, numb := range strings.Split(string(content), ",") {
		int_numb, _ := strconv.Atoi(numb)
		numbers = append(numbers, int_numb)
	}
	return
}

func ExecuteCommand(input []int, start int) []int {
	method := input[start]
	val1 := input[start+1]
	val2 := input[start+2]
	dest := input[start+3]

	if method == 1 {
		input[dest] = input[val1] + input[val2]
	} else if method == 2 {
		input[dest] = input[val1] * input[val2]
	} else {
		log.Fatal("something went wrong..")
	}
	return input

}

func DayTwoOne(noun int, verb int) int {
	input := GetInput()
	input[1] = noun
	input[2] = verb
	for i := 0; i < len(input); i += 4 {
		method := input[i]
		if method == 99 {
			break
		}
		input = ExecuteCommand(input, i)
	}
	return input[0]

}
