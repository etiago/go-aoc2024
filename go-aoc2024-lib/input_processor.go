package goaoc2024lib

import (
	"log"
	"os"
	"strconv"
	"strings"
)

func ReadFile(input_file_path *string) string {
	content, err := os.ReadFile(*input_file_path)
	if err != nil {
		log.Fatal(err)
	}

	return string(content)
}

func ReadFileLines(input_file_path *string) []string {
	content := ReadFile(input_file_path)
	return strings.Split(content, "\n")
}

func LoadDay5Rules(input_file_path *string) map[int]map[int]struct{} {
	content := ReadFileLines(input_file_path)

	rules := make(map[int]map[int]struct{})
	for _, line := range content {
		if line == "" {
			break
		}

		parts := strings.Split(line, "|")

		left, _ := strconv.Atoi(parts[0])
		right, _ := strconv.Atoi(parts[1])

		if _, ok := rules[left]; !ok {
			rules[left] = make(map[int]struct{})
			rules[left][right] = struct{}{}
		} else {
			rules[left][right] = struct{}{}
		}
	}

	return rules
}

func LoadDay5Updates(input_file_path *string) [][]int {
	content := ReadFileLines(input_file_path)

	skipRules := true
	updates := make([][]int, 0)

	for _, line := range content {
		if line != "" && skipRules {
			continue
		}

		if line == "" {
			skipRules = false
			continue
		}

		parts := strings.Split(line, ",")
		update := make([]int, 0)
		for part := range parts {
			num, _ := strconv.Atoi(parts[part])
			update = append(update, num)
		}

		updates = append(updates, update)
	}

	return updates
}

type GuardPosition struct {
	x int
	y int
}

type GuardPositionWithOrientation struct {
	guardPosition GuardPosition
	orientation   Orientation
}

func LoadDay6Map(inputFilePath *string) ([][]rune, *GuardPosition) {
	content := ReadFileLines(inputFilePath)

	mapArray := make([][]rune, len(content))
	var guardPosition *GuardPosition = nil

	for i, line := range content {
		mapArray[i] = []rune(line)

		maybeGuardX := strings.Index(line, "^")
		if maybeGuardX != -1 {
			guardPosition = &GuardPosition{maybeGuardX, i}
		}
	}

	return mapArray, guardPosition
}

type Equation struct {
	result   int64
	operands []int64
}

func LoadDay7Equations(inputFilePath *string) []Equation {
	content := ReadFileLines(inputFilePath)

	equations := make([]Equation, len(content))
	for i, line := range content {
		chunks := strings.Split(line, ": ")
		result, _ := strconv.ParseInt(chunks[0], 10, 64)

		operands := strings.Split(chunks[1], " ")
		operandsInt := make([]int64, len(operands))
		for i, operand := range operands {
			operandsInt[i], _ = strconv.ParseInt(operand, 10, 64)
		}

		equation := Equation{result, operandsInt}
		equations[i] = equation
	}
	return equations
}
