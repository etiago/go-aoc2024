package goaoc2024lib

import "testing"

func TestDay4Part1RealInput(t *testing.T) {
	input_path := "../input/day4.txt"
	input := ReadFileLines(&input_path)
	charMatrix := stringLinesToCharMatrix(input)
	result := day4Part1(charMatrix)

	if result != 2464 {
		t.Errorf("Expected 2464, got %d", result)
	}
}

func TestDay4Part1Example(t *testing.T) {
	input_path := "../input/day4-example.txt"
	input := ReadFileLines(&input_path)
	charMatrix := stringLinesToCharMatrix(input)
	result := day4Part1(charMatrix)

	if result != 18 {
		t.Errorf("Expected 18, got %d", result)
	}
}

func TestDay4Part2RealInput(t *testing.T) {
	input_path := "../input/day4.txt"
	input := ReadFileLines(&input_path)
	charMatrix := stringLinesToCharMatrix(input)
	result := day4Part2(charMatrix)

	if result != 1982 {
		t.Errorf("Expected 1982, got %d", result)
	}
}

func TestDay4Part2Example(t *testing.T) {
	input_path := "../input/day4-example.txt"
	input := ReadFileLines(&input_path)
	charMatrix := stringLinesToCharMatrix(input)
	result := day4Part2(charMatrix)

	if result != 9 {
		t.Errorf("Expected 9, got %d", result)
	}
}
