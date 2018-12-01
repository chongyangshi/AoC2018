package day1

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

func Run(input string) string {
	var inputLines = strings.Split(input, "\n")
	for k, v := range inputLines {
		inputLines[k] = strings.Trim(v, " \t")
	}
	return fmt.Sprintf("Part 1: %v\nPart 2: %v\n", part1(inputLines), part2(inputLines))
}

func part1(inputLines []string) string {
	var accumulator int64
	for _, v := range inputLines {
		if len(v) == 0 {
			continue
		}

		i, err := strconv.ParseInt(v, 10, 64)
		if err != nil {
			log.Fatal("Error parsing int: ", err)
		}
		accumulator += i
	}

	return strconv.FormatInt(accumulator, 10)
}

func part2(inputLines []string) string {
	var frequency int64
	var frequencies = map[int64]int{}
	var index int
	for {
		v := inputLines[index]
		if len(v) == 0 {
			continue
		}

		i, err := strconv.ParseInt(v, 10, 64)
		if err != nil {
			log.Fatal("Error parsing int: ", err)
		}
		frequency += i

		if _, ok := frequencies[frequency]; !ok {
			frequencies[frequency] = 1
		} else {
			return strconv.FormatInt(frequency, 10)
		}

		if index == len(inputLines)-1 {
			index = 0
		} else {
			index++
		}
	}
}
