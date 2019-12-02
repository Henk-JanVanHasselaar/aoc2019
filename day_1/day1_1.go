package main

import (
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func readFile(filePath string) (numbers []int) {
	fd, err := os.Open(filePath)
	if err != nil {
		panic(fmt.Sprintf("open %s: %v", filePath, err))
	}
	var line int
	for {

		_, err := fmt.Fscanf(fd, "%d\n", &line)

		if err != nil {
			if err == io.EOF {
				return
			}
			panic(fmt.Sprintf("Scan Failed %s: %v", filePath, err))

		}
		numbers = append(numbers, line)
	}
	return
}

func FloatToString(input_num float64) string {
	// to convert a float number to a string
	return strconv.FormatFloat(input_num, 'f', 6, 64)
}

func DayOneOne() int {
	numbers := readFile("input/input.txt")

	final := 0
	for _, item := range numbers {
		result := item / 3
		// parse float to string, so we can split on "."
		str_result := FloatToString(float64(result))
		res := strings.Split(str_result, ".")[0]
		i1, _ := strconv.Atoi(res)
		final += i1 - 2
	}
	return final
}
