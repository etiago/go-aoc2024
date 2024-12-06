package goaoc2024lib

import "testing"

func TestDay6Part2Example(t *testing.T) {
	inputFilePath := "../input/day6-example.txt"
	blockageCounts := day6Part2(&inputFilePath)
	expectedBlockageCounts := 6
	if blockageCounts != 6 {
		t.Errorf("Got %d, expected %d", blockageCounts, expectedBlockageCounts)
	}
}

func TestDay6Part2RealData(t *testing.T) {
	inputFilePath := "../input/day6.txt"
	blockageCounts := day6Part2(&inputFilePath)
	expectedBlockageCounts := 1939
	if blockageCounts != 1939 {
		t.Errorf("Got %d, expected %d", blockageCounts, expectedBlockageCounts)
	}
}
