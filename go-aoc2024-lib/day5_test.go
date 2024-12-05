package goaoc2024lib

import "testing"

func TestDay5Part1Example(t *testing.T) {
	expectedSum := 143

	inputPath := "../input/day5-example.txt"

	sum := day5Part1(&inputPath)

	if sum != expectedSum {
		t.Errorf("Expected %d, got %d", expectedSum, sum)
	}
}

func TestDay5Part1RealData(t *testing.T) {
	expectedSum := 4957

	inputPath := "../input/day5.txt"

	sum := day5Part1(&inputPath)

	if sum != expectedSum {
		t.Errorf("Expected %d, got %d", expectedSum, sum)
	}
}

func TestDay5Part2Example(t *testing.T) {
	expectedSum := 123

	inputPath := "../input/day5-example.txt"

	sum := day5Part2(&inputPath)

	if sum != expectedSum {
		t.Errorf("Expected %d, got %d", expectedSum, sum)
	}
}

func TestDay5Part2RealData(t *testing.T) {
	expectedSum := 6938

	inputPath := "../input/day5.txt"

	sum := day5Part2(&inputPath)

	if sum != expectedSum {
		t.Errorf("Expected %d, got %d", expectedSum, sum)
	}
}
