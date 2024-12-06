package goaoc2024lib

import (
	"log"
	"sync"
)

type Orientation int

const (
	FacingUp Orientation = iota
	FacingRight
	FacingDown
	FacingLeft
)

func guardCanWalkStraight(mapArray [][]rune, guardPosition GuardPosition, orientation Orientation) (bool, *GuardPosition) {
	var nextGuardPosition *GuardPosition = nil

	if orientation == FacingUp {
		nextGuardPosition = &GuardPosition{x: guardPosition.x, y: guardPosition.y - 1}
	} else if orientation == FacingRight {
		nextGuardPosition = &GuardPosition{x: guardPosition.x + 1, y: guardPosition.y}
	} else if orientation == FacingDown {
		nextGuardPosition = &GuardPosition{x: guardPosition.x, y: guardPosition.y + 1}
	} else if orientation == FacingLeft {
		nextGuardPosition = &GuardPosition{x: guardPosition.x - 1, y: guardPosition.y}
	}

	guardCanStep := !isGuardInMap(mapArray, *nextGuardPosition) || mapArray[nextGuardPosition.y][nextGuardPosition.x] != '#'
	return guardCanStep, nextGuardPosition
}

func isGuardInMap(mapArray [][]rune, guardPosition GuardPosition) bool {
	return guardPosition.x >= 0 && guardPosition.x < len(mapArray) && guardPosition.y >= 0 && guardPosition.y < len(mapArray[0])
}

func getVisitedPositions(mapArray [][]rune, guardPosition GuardPosition, stopIfLoop bool) (map[GuardPosition]struct{}, bool) {
	orientation := FacingUp

	visitedPositions := make(map[GuardPosition]struct{})
	visitedPositions[guardPosition] = struct{}{}

	visitedPositionsWithOrientation := make(map[GuardPositionWithOrientation]struct{})

	for {
		canWalkStraight, newPosition := guardCanWalkStraight(mapArray, guardPosition, orientation)
		newPositionIsInMap := isGuardInMap(mapArray, *newPosition)

		if canWalkStraight && !newPositionIsInMap {
			break
		}

		if canWalkStraight {
			guardPosition = *newPosition
			// Have we visited this position with this orientation
			if _, ok := visitedPositionsWithOrientation[GuardPositionWithOrientation{guardPosition: guardPosition, orientation: orientation}]; ok {
				if stopIfLoop {
					return visitedPositions, true
				}
			}
			visitedPositions[guardPosition] = struct{}{}
			visitedPositionsWithOrientation[GuardPositionWithOrientation{guardPosition: guardPosition, orientation: orientation}] = struct{}{}
		} else {
			orientation = (orientation + 1) % 4
		}
	}

	return visitedPositions, false
}

func day6Part1(inputFilePath *string) int {
	mapArray, guardPosition := LoadDay6Map(inputFilePath)

	// Map of visited positions
	visitedPositions, _ := getVisitedPositions(mapArray, *guardPosition, false)

	log.Println("Visited positions: ", len(visitedPositions))

	return len(visitedPositions)
}

func makeMapCopy(mapArray [][]rune) [][]rune {
	mapCopy := make([][]rune, len(mapArray))
	for i := range mapArray {
		mapCopy[i] = make([]rune, len(mapArray[i]))
		copy(mapCopy[i], mapArray[i])
	}

	return mapCopy
}

// I started with a non-concurrent version which was reasonable,
// but decided to try and parallelise it since it lends itself to it.
// It's much faster (by around a factor of 5).
func day6Part2Concurrent(inputFilePath *string) int {
	mapArray, guardPosition := LoadDay6Map(inputFilePath)

	visitedPositions, _ := getVisitedPositions(mapArray, *guardPosition, false)

	results := make(chan bool)
	var wg sync.WaitGroup

	for position := range visitedPositions {
		wg.Add(1)
		go func(p GuardPosition) {
			defer wg.Done()
			mapCopy := makeMapCopy(mapArray)

			mapCopy[position.y][position.x] = '#'

			_, isLoop := getVisitedPositions(mapCopy, *guardPosition, true)
			results <- isLoop
		}(position)

	}

	go func() {
		wg.Wait()
		close(results)
	}()

	blockageCounts := 0
	for isLoop := range results {
		if isLoop {
			blockageCounts++
		}
	}

	log.Println("Blockage counts: ", blockageCounts)

	return blockageCounts
}

func day6Part2(inputFilePath *string) int {
	mapArray, guardPosition := LoadDay6Map(inputFilePath)

	visitedPositions, _ := getVisitedPositions(mapArray, *guardPosition, false)

	blockageCounts := 0

	for position := range visitedPositions {
		// newGuardPosition := GuardPosition{x: guardPosition.x, y: guardPosition.y}
		mapArray[position.y][position.x] = '#'

		_, isLoop := getVisitedPositions(mapArray, *guardPosition, true)
		if isLoop {
			blockageCounts++
		}
		mapArray[position.y][position.x] = '.'
	}
	log.Println("Blockage counts: ", blockageCounts)

	return blockageCounts
}

func Day6(inputFilePath *string) {
	day6Part1(inputFilePath)
	day6Part2Concurrent(inputFilePath)
}
