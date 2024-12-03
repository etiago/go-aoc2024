package goaoc2024lib

import (
	"log"
	"strconv"
	"strings"
)

func getDelta(a int, b int) int {
	if a > b {
		return a - b
	} else {
		return b - a
	}
}

func isIncreasing(a int, b int) bool {
	return a < b
}

// Brute force and feels like there should be a better way...
// But it works.
func reportIsSafeWithDampening(levels []string) bool {
	// Early return if the report is already safe.
	if reportIsSafe(levels) {
		return true
	}

	levelCount := len(levels)
	foundSafe := false
	for i := 0; i < levelCount; i++ {
		// Slice the levels, removing index i
		levelsCopy := make([]string, levelCount-1)
		copy(levelsCopy, levels[:i])
		copy(levelsCopy[i:], levels[i+1:])
		if reportIsSafe(levelsCopy) {
			foundSafe = true
			break
		}
	}
	return foundSafe
}

func reportIsSafe(levels []string) bool {
	lastLevel, _ := strconv.Atoi(levels[1])
	beforeLastLevel, _ := strconv.Atoi(levels[0])
	wasIncreasing := isIncreasing(beforeLastLevel, lastLevel)
	initialDelta := getDelta(beforeLastLevel, lastLevel)
	if initialDelta < 1 || initialDelta > 3 {
		return false
	}

	for i := 2; i < len(levels); i++ {
		level, _ := strconv.Atoi(levels[i])
		delta := getDelta(lastLevel, level)
		increasing := isIncreasing(lastLevel, level)

		if increasing != wasIncreasing {
			return false
		}

		if delta < 1 {
			return false
		} else if delta > 3 {
			return false
		}

		lastLevel = level
	}

	return true
}

func day2Part1(input_file_path *string) {
	log.Println("Part 1")

	lines := ReadFileLines(input_file_path)

	safe_reports := 0
	for _, line := range lines {
		levels := strings.Fields(line)
		if reportIsSafe(levels) {
			safe_reports++
		}
	}

	log.Println("Safe reports:", safe_reports)
}

func day2Part2(input_file_path *string) {
	log.Println("Part 2")

	lines := ReadFileLines(input_file_path)

	// Parse input
	safe_reports := 0
	for _, line := range lines {
		levels := strings.Fields(line)
		if reportIsSafeWithDampening(levels) {
			safe_reports++
		}
	}

	log.Println("Safe reports:", safe_reports)

}
func Day2(input_file_path *string) {
	log.Println("Day 1")

	day2Part1(input_file_path)
	day2Part2(input_file_path)
}
