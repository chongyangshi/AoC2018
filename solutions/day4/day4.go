package day4

import (
	"fmt"
	"log"
	"sort"
	"strconv"
	"strings"
	"time"
)

var totalSleepByGuard = map[int64]int{}
var minuteCounterByGuard = map[int64]map[int]int{}

func Run(input string) string {
	var inputLines = strings.Split(input, "\n")
	for k, v := range inputLines {
		inputLines[k] = strings.Trim(v, " \t")
	}
	process(inputLines)

	return fmt.Sprintf("Part 1: %v\nPart 2: %v\n", part1(), part2())
}

func process(inputLines []string) {
	var eventsByDay = map[string][]int{}
	var guardSchedule = map[int64][]string{}
	var dateToGuard = map[string]int64{}
	for _, line := range inputLines {
		lineSplit := strings.Split(strings.Trim(line, " "), " ")

		date, err := time.Parse("2006-01-02 15:04", lineSplit[0][1:]+" "+lineSplit[1][:5])
		if err != nil {
			log.Fatal("Bad date on line " + line)
		}
		dateString := date.Format("2006-01-02")

		if strings.Contains(line, "begins shift") {
			guardNumber, err := strconv.ParseInt(lineSplit[3][1:], 10, 64)
			if err != nil {
				log.Fatal("Bad guard number on line " + line)
			}

			shiftDate := dateString
			if date.Hour() != 0 {
				shiftDate = date.AddDate(0, 0, 1).Format("2006-01-02")
			}
			guardSchedule[guardNumber] = append(guardSchedule[guardNumber], shiftDate)
			dateToGuard[shiftDate] = guardNumber
		} else {
			// A guard always falls asleep first and wakes up within the hour, so we just maintain a sequence of minutes
			eventsByDay[dateString] = append(eventsByDay[dateString], date.Minute())
		}
	}

	for guard, guardDays := range guardSchedule {
		for _, day := range guardDays {
			eventsOfDay := eventsByDay[day]
			sort.Ints(eventsOfDay)
			sleepTime := 0
			for i := 0; i < len(eventsOfDay); i += 2 {
				sleepTime += eventsOfDay[i+1] - eventsOfDay[i]
				for j := eventsOfDay[i]; j < eventsOfDay[i+1]; j++ {
					if _, ok := minuteCounterByGuard[guard]; ok {
						if count, countOK := minuteCounterByGuard[guard][j]; countOK {
							minuteCounterByGuard[guard][j] = count + 1
						} else {
							minuteCounterByGuard[guard][j] = 1
						}
					} else {
						minuteCounterByGuard[guard] = map[int]int{}
						minuteCounterByGuard[guard][j] = 1
					}
				}
			}
			if val, ok := totalSleepByGuard[guard]; ok {
				totalSleepByGuard[guard] = val + sleepTime
			} else {
				totalSleepByGuard[guard] = sleepTime
			}
		}
	}
}

func part1() string {
	var maxSleepGuard int64
	var maxSleepTime = -1
	for guard, sleep := range totalSleepByGuard {
		if sleep > maxSleepTime {
			maxSleepGuard = guard
			maxSleepTime = sleep
		}
	}

	minutesCounter := minuteCounterByGuard[maxSleepGuard]
	var maxSleepMinute int
	var maxSleepMinuteCount = -1
	for minute, count := range minutesCounter {
		if count > maxSleepMinuteCount {
			maxSleepMinuteCount = count
			maxSleepMinute = minute
		}
	}

	return strconv.FormatInt(int64(maxSleepMinute)*maxSleepGuard, 10)
}

func part2() string {
	var maxSleepGuard int64
	var maxSleepMinute int
	var maxSleepFrequency = -1
	for guard, counter := range minuteCounterByGuard {
		for minute, count := range counter {
			if count > maxSleepFrequency {
				maxSleepMinute = minute
				maxSleepGuard = guard
				maxSleepFrequency = count
			}
		}
	}

	return strconv.FormatInt(int64(maxSleepMinute)*maxSleepGuard, 10)
}
