package goaoc2024lib

import (
	"testing"
)

func TestPushBoxesAndMove(t *testing.T) {
	testFilePath := "../input/day15-test.txt"
	puzzle := LoadDay15Map(&testFilePath, false)

	// Print puzzle matrix
	for _, row := range puzzle.mapArray {
		t.Log(string(row))
	}

	for _, movement := range puzzle.movements {
		tryPushBoxesAndMove(movement, &puzzle)
	}

	// Print puzzle matrix
	for _, row := range puzzle.mapArray {
		t.Log(string(row))
	}

}

func TestPushBoxesAndMoveExpandedPuzzle(t *testing.T) {
	testFilePath := "../input/day15-large.txt"
	puzzle := LoadDay15Map(&testFilePath, true)

	// Print puzzle matrix
	for _, row := range puzzle.mapArray {
		t.Log(string(row))
	}

	for _, movement := range puzzle.movements {
		tryPushBoxesAndMoveExpandedPuzzle(movement, &puzzle)
		for _, row := range puzzle.mapArray {
			t.Log(string(row))
		}
		// pushables := tryPushBoxesAndMoveExpandedPuzzle(Up, &puzzle)
		// pushables = tryPushBoxesAndMoveExpandedPuzzle(Right, &puzzle)
		// pushables = tryPushBoxesAndMoveExpandedPuzzle(Up, &puzzle)
		// pushables = tryPushBoxesAndMoveExpandedPuzzle(Right, &puzzle)
		// pushables = tryPushBoxesAndMoveExpandedPuzzle(Right, &puzzle)
		// pushables = tryPushBoxesAndMoveExpandedPuzzle(Right, &puzzle)

	}

	// t.Log("Pushables:", pushables)
	// Print puzzle matrix

}
