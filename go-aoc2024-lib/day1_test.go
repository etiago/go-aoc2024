package goaoc2024lib

import "testing"

func TestDay1Part1(t *testing.T) {
	input_path := "../input/day1.txt"
	input := ReadFileLines(&input_path)
	result := day1Part1(input)

	if result != 1938424 {
		t.Errorf("Expected 1938424, got %d", result)
	}
}

func TestDay1Part2(t *testing.T) {
	input_path := "../input/day1.txt"
	input := ReadFileLines(&input_path)
	result := day1Part2(input)

	if result != 22014209 {
		t.Errorf("Expected 22014209, got %d", result)
	}
}
