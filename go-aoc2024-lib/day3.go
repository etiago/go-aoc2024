package goaoc2024lib

import (
	"log"
	"regexp"
	"strconv"
)

type Instruction struct {
	left  int
	right int
}

// This felt a bit like cheating, so since I ended up doing proper parsing
// for part 2, I converted part 1 to use the parsing approach. But leaving this
// here for posterity.
func parseMulInstructions(input string) []Instruction {
	instructions := make([]Instruction, 0)

	r := regexp.MustCompile("mul\\((?P<left>[0-9]+),(?P<right>[0-9]+)\\)")
	matches := r.FindAllStringSubmatch(input, -1)

	for _, match := range matches {
		left, _ := strconv.Atoi(match[1])
		right, _ := strconv.Atoi(match[2])
		instruction := Instruction{left, right}
		instructions = append(instructions, instruction)
	}

	return instructions
}

func day3Part1(input_file_path *string) {
	// Part 1
	input := ReadFile(input_file_path)
	instructions := parseMulInstructions(input)
	// instructions := parseMulWithDoAndDont(input, false)

	total := 0

	for _, instruction := range instructions {
		total += instruction.left * instruction.right
	}

	log.Println("Total:", total)
}

type ParsingState int

const (
	FindMul ParsingState = iota
	FindLeft
	FindRight
	FindDo
)

func parseMulWithDoAndDont(input string, checkDoAndDont bool) []Instruction {
	instructions := make([]Instruction, 0)

	state := FindMul
	var left, right int
	numStartIdx := 0
	for i := 0; i < len(input); i++ {
		// First check if we have a dont()
		if checkDoAndDont && i+7 < len(input) && input[i:i+7] == "don't()" {
			state = FindDo
			i += 6
			continue
		}

		switch state {
		case FindDo:
			if checkDoAndDont && i+4 < len(input) && input[i:i+4] == "do()" {
				state = FindMul
				i += 3
			}
		case FindMul:
			if i+4 < len(input) && input[i:i+4] == "mul(" {
				state = FindLeft
				numStartIdx = i + 4
				i += 3
			}
		case FindLeft:
			if input[i] == ',' {
				substr := input[numStartIdx:i]
				leftLocal, err := strconv.Atoi(substr)
				if err == nil {
					left = leftLocal
					state = FindRight
					numStartIdx = i + 1
				} else {
					state = FindMul
				}
			} else if input[i] < '0' || input[i] > '9' {
				state = FindMul
			}
		case FindRight:
			if input[i] == ')' {
				substr := input[numStartIdx:i]
				rightLocal, err := strconv.Atoi(substr)
				if err == nil {
					right = rightLocal
					instruction := Instruction{left, right}
					instructions = append(instructions, instruction)
					state = FindMul
					numStartIdx = i + 1
				} else {
					state = FindMul
				}
			} else if input[i] < '0' || input[i] > '9' {
				state = FindMul
			}
		}
	}
	return instructions
}

func day3Part2(input_file_path *string) {
	// Part 2
	input := ReadFile(input_file_path)
	instructions := parseMulWithDoAndDont(input, true)

	total := 0

	for _, instruction := range instructions {
		total += instruction.left * instruction.right
	}

	log.Println("Total:", total)
}

func Day3(input_file_path *string) {
	// Part 1
	// day3Part1(input_file_path)
	// Part 2
	day3Part2(input_file_path)
}
