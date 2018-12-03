package day3

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
	rows := process(inputLines)
	return fmt.Sprintf("Part 1: %v\nPart 2: %v\n", part1(rows), part2(rows))
}

type rowsMap struct {
	rows      map[int64]map[int64]int
	claimants map[int64]map[int64][]int64
}

func (r *rowsMap) registerClaim(row, start, end, claimant int64) {
	_, ok := r.rows[row]
	if !ok {
		r.rows[row] = map[int64]int{}
		r.claimants[row] = map[int64][]int64{}
	}

	for i := start; i <= end; i++ {
		if val, ok := r.rows[row][i]; ok {
			r.rows[row][i] = val + 1
		} else {
			r.rows[row][i] = 1
		}
		r.claimants[row][i] = append(r.claimants[row][i], claimant)
	}
}

func process(inputLines []string) *rowsMap {
	var rows = rowsMap{
		rows:      map[int64]map[int64]int{},
		claimants: map[int64]map[int64][]int64{},
	}

	for _, line := range inputLines {
		if len(line) == 0 {
			continue
		}

		lineSplit := strings.Split(strings.Trim(line, " \t\r"), " ")
		if len(lineSplit) != 4 {
			log.Fatal("Bad split of input line " + line)
		}
		id, err := strconv.ParseInt(lineSplit[0][1:], 10, 64)
		if err != nil {
			log.Fatal("Bad ID on input line " + line)
		}

		offsets := strings.Split(lineSplit[2], ",")
		if len(offsets) != 2 {
			log.Fatal("Bad split of offset on input line " + line)
		}
		xOffset, errx := strconv.ParseInt(offsets[0], 10, 64)
		yOffset, erry := strconv.ParseInt(offsets[1][:len(offsets[1])-1], 10, 64)
		if errx != nil || erry != nil {
			log.Fatal("Bad offsets on input line " + line)
		}

		lengths := strings.Split(lineSplit[3], "x")
		if len(lengths) != 2 {
			log.Fatal("Bad split of lengths on input line " + line)
		}
		xLength, errx := strconv.ParseInt(lengths[0], 10, 64)
		yLength, erry := strconv.ParseInt(lengths[1], 10, 64)
		if errx != nil || erry != nil {
			log.Fatal("Bad lengths on input line " + line)
		}

		for x := xOffset; x < xOffset+xLength; x++ {
			rows.registerClaim(x, yOffset, yOffset+yLength-1, id)
		}
	}

	return &rows
}

func part1(rows *rowsMap) string {
	var accumulator int64
	for _, row := range rows.rows {
		for _, sq := range row {
			if sq > 1 {
				accumulator++
			}
		}
	}
	return strconv.FormatInt(accumulator, 10)
}

func part2(rows *rowsMap) string { // hideous
	goodClaimants := map[int64]bool{}
	for _, crow := range rows.claimants {
		for _, cs := range crow {
			for _, claimant := range cs {
				if len(cs) > 1 {
					goodClaimants[claimant] = false
				} else {
					if _, ok := goodClaimants[claimant]; !ok {
						goodClaimants[claimant] = true
					}
				}
			}
		}
	}

	for i, goodClaim := range goodClaimants {
		if goodClaim {
			return strconv.FormatInt(i, 10)
		}
	}

	return "(Not Found)"
}
