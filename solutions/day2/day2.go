package day2

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
	var twoAccumulator int64
	var threeAccumulator int64
	for _, line := range inputLines {
		if len(line) == 0 {
			continue
		}

		letterCounter := map[rune]int{}
		for _, c := range line {
			val, ok := letterCounter[c]
			if ok {
				letterCounter[c] = val + 1
			} else {
				letterCounter[c] = 1
			}
		}

		hasTwo := false
		hasThree := false
		for _, v := range letterCounter {
			if v == 2 {
				hasTwo = true
			} else if v == 3 {
				hasThree = true
			}
		}

		if hasTwo {
			twoAccumulator++
		}
		if hasThree {
			threeAccumulator++
		}
	}

	return strconv.FormatInt(twoAccumulator*threeAccumulator, 10)
}

func part2(inputLines []string) string {
	for k, line1 := range inputLines {
		for _, line2 := range inputLines[k:] {
			if len(line1) != len(line2) {
				log.Fatal("Mismatch line length, violating question assumption")
			}
			textDistance := 0
			lastMismatchIndex := 0
			for n, c := range line1 {
				if line2[n] != byte(c) {
					textDistance++
					lastMismatchIndex = n
				}
				if textDistance == 2 {
					break
				}
			}

			if textDistance != 1 {
				continue
			} else {
				return line1[:lastMismatchIndex] + line1[lastMismatchIndex+1:]
			}
		}
	}

	return "(Not Found)"
}
