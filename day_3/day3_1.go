package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func GetInput() (data []string) {
	file, err := os.Open("input/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		data = append(data, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return
}

type Cable struct {
	PosX      int
	PosY      int
	StepCount int
	Visited   map[string]int
}

func (cable *Cable) position() string {
	return fmt.Sprintf("%dx%d", cable.PosX, cable.PosY)
}

func (cable *Cable) walk(direction string) {
	method := direction[:1]
	steps, _ := strconv.Atoi(direction[1:])
	for i := 0; i < steps; i++ {
		switch method {
		case "U":
			cable.PosY++
		case "R":
			cable.PosX++
		case "D":
			cable.PosY--
		case "L":
			cable.PosX--
		}
		cable.StepCount++
		if _, ok := cable.Visited[cable.position()]; ok {
			continue
		} else {
			cable.Visited[cable.position()] = cable.StepCount
		}

	}
}
func ManhattanDistance(coord string) int {
	coords := strings.Split(coord, "x")
	x, _ := strconv.Atoi(coords[0])
	y, _ := strconv.Atoi(coords[1])
	return x + y
}

func DayThreeOne() int {
	data := GetInput()
	var cables []Cable
	for _, row := range data {
		cable := Cable{0, 0, 0, make(map[string]int)}
		for _, direction := range strings.Split(string(row), ",") {
			cable.walk(direction)
		}
		cables = append(cables, cable)
	}
	result := 9999999999

	for k, _ := range cables[0].Visited {
		if _, ok := cables[1].Visited[k]; ok {
			distance := ManhattanDistance(k)
			if distance < result {
				result = distance
			}
		}
	}
	return result
}
