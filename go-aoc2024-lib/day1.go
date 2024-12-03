package goaoc2024lib

import (
	"log"
	"sort"
	"strconv"
	"strings"
)

func part1(lines []string) {
	// Need two lists: left and right, length is known to be
	// size of input
	left := make([]int, len(lines))
	right := make([]int, len(lines))

	// Iterate through lines of input
	for i, line := range lines {
		chunks := strings.Fields(line)
		left[i], _ = strconv.Atoi(chunks[0])
		right[i], _ = strconv.Atoi(chunks[1])
	}

	// Sort both lists
	sort.Ints(left)
	sort.Ints(right)

	difference := 0

	for i := 0; i < len(left); i++ {
		// Add the absolute difference
		if left[i] > right[i] {
			difference += left[i] - right[i]
		} else {
			difference += right[i] - left[i]
		}
	}

	log.Println("Difference:", difference)
}

func part2(lines []string) {
	// Part 2

	left := make([]int, len(lines))
	right_freq := make(map[int]int)

	for i, line := range lines {
		chunks := strings.Fields(line)
		left[i], _ = strconv.Atoi(chunks[0])
		right_value, _ := strconv.Atoi(chunks[1])
		right_freq[right_value]++
	}
	sort.Ints(left)

	score := 0
	// Iterate through left
	for _, l := range left {
		// Get number of times l appears in right
		n_times := right_freq[l]
		score += l * n_times
	}

	log.Println("Score:", score)
}

func Day1(input_file_path *string) {
	// Day 1 code here
	log.Println("Day 1")

	input := ReadFileLines(input_file_path)

	part1(input)
	part2(input)
}
