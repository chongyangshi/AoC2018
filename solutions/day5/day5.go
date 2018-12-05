package day5

import (
	"fmt"
	"log"
	"strings"
	"unicode"
)

// Was trying to do something cleverer, gave up. So have a very slow solution indeed,
// because of so many array copies.

func Run(input string) string {
	input = strings.Trim(input, " \r\n\t")
	if len(input) < 2 {
		log.Fatal("Invalid input (too short)")
	}
	return fmt.Sprintf("Part 1: %v\nPart 2: %v\n", part1(input), part2(input))
}

func react(polymer []rune) int {
	for {
		reacted := map[int]bool{}
		i := 0
		for i < len(polymer)-1 {
			if unicode.IsUpper(polymer[i]) {
				if unicode.ToLower(polymer[i]) == polymer[i+1] {
					reacted[i] = true
					reacted[i+1] = true
					i += 2
				} else {
					i++
				}
			} else {
				if unicode.ToUpper(polymer[i]) == polymer[i+1] {
					reacted[i] = true
					reacted[i+1] = true
					i += 2
				} else {
					i++
				}
			}
		}

		if len(reacted) == 0 {
			break
		}

		var newPolymer = []rune{}
		for i := 0; i < len(polymer); i++ {
			if _, ok := reacted[i]; !ok {
				newPolymer = append(newPolymer, polymer[i])
			}
		}

		polymer = newPolymer
	}

	return len(polymer)
}

func part1(input string) int {
	return react([]rune(input))
}

func part2(input string) int {
	originalPolymer := []rune(input)
	units := map[rune]bool{}
	for _, r := range originalPolymer {
		unit := unicode.ToLower(r)
		if _, ok := units[unit]; !ok {
			units[unit] = true
		}
	}

	bestLength := len(originalPolymer)
	for unit := range units {
		reductedPolymer := []rune{}
		unitUpper := unicode.ToUpper(unit)
		for _, r := range originalPolymer {
			if r != unit && r != unitUpper {
				reductedPolymer = append(reductedPolymer, r)
			}
		}

		reactedPolymerLength := react(reductedPolymer)
		if reactedPolymerLength < bestLength {
			bestLength = reactedPolymerLength
		}
	}

	return bestLength
}
