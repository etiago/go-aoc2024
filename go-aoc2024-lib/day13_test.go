package goaoc2024lib

import "testing"

func TestGetCheapestCombinationExample1(t *testing.T) {
	expectedA := 80
	expectedB := 40
	game := Day13Game{
		prize: Point{
			x: 8400,
			y: 5400,
		},
		buttonAStep: Point{
			x: 94,
			y: 34,
		},
		buttonBStep: Point{
			x: 22,
			y: 67,
		},
	}

	pressesOnA, pressesOnB, found := getCheapestCombination(game)
	if pressesOnA != expectedA || pressesOnB != expectedB || !found {
		t.Errorf("Expected %d, %d, got %d, %d", expectedA, expectedB, pressesOnA, pressesOnB)
	}
}

func TestGetCheapestCombinationExample1NewApproach(t *testing.T) {
	expectedA := 80
	expectedB := 40
	game := Day13Game{
		prize: Point{
			x: 8400,
			y: 5400,
		},
		buttonAStep: Point{
			x: 94,
			y: 34,
		},
		buttonBStep: Point{
			x: 22,
			y: 67,
		},
	}

	pressesOnA, pressesOnB, found := getPresses(game)
	if pressesOnA != expectedA || pressesOnB != expectedB || !found {
		t.Errorf("Expected %d, %d, got %d, %d", expectedA, expectedB, pressesOnA, pressesOnB)
	}
}

func TestGetCheapestCombinationExample2(t *testing.T) {
	game := Day13Game{
		prize: Point{
			x: 12748,
			y: 12176,
		},
		buttonAStep: Point{
			x: 26,
			y: 66,
		},
		buttonBStep: Point{
			x: 67,
			y: 21,
		},
	}

	pressesOnA, pressesOnB, found := getCheapestCombination(game)
	if found {
		t.Errorf("Expected no result, got %d, %d", pressesOnA, pressesOnB)
	}
}

func TestGetCheapestCombinationExample2NewApproach(t *testing.T) {
	game := Day13Game{
		prize: Point{
			x: 12748,
			y: 12176,
		},
		buttonAStep: Point{
			x: 26,
			y: 66,
		},
		buttonBStep: Point{
			x: 67,
			y: 21,
		},
	}

	pressesOnA, pressesOnB, found := getPresses(game)
	if found {
		t.Errorf("Expected no result, got %d, %d", pressesOnA, pressesOnB)
	}
}

func TestGetCheapestCombinationExample3(t *testing.T) {
	expectedA := 38
	expectedB := 86
	game := Day13Game{
		prize: Point{
			x: 7870,
			y: 6450,
		},
		buttonAStep: Point{
			x: 17,
			y: 86,
		},
		buttonBStep: Point{
			x: 84,
			y: 37,
		},
	}

	pressesOnA, pressesOnB, found := getCheapestCombination(game)
	if pressesOnA != expectedA || pressesOnB != expectedB || !found {
		t.Errorf("Expected %d, %d, got %d, %d", expectedA, expectedB, pressesOnA, pressesOnB)
	}
}
