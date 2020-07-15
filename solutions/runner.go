package solutions

import (
	"time"

	"github.com/chongyangshi/AoC2018/solutions/day1"
	"github.com/chongyangshi/AoC2018/solutions/day2"
	"github.com/chongyangshi/AoC2018/solutions/day3"
	"github.com/chongyangshi/AoC2018/solutions/day4"
	"github.com/chongyangshi/AoC2018/solutions/day5"
)

var runMap = map[string]interface{}{
	"1": day1.Run,
	"2": day2.Run,
	"3": day3.Run,
	"4": day4.Run,
	"5": day5.Run,
}

func RunSolution(solution string, input string) (string, time.Duration) {
	var results string
	start := time.Now()

	if f, ok := runMap[solution]; ok {
		results = f.(func(string) string)(input)
	} else {
		results = "(Solution code not found)"
	}
	realExecutionTime := time.Now().Sub(start)
	return results, realExecutionTime
}
