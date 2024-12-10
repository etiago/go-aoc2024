package goaoc2024lib

import "log"

var possibleDeltas = [4]Point{
	Point{1, 0},
	Point{0, 1},
	Point{-1, 0},
	Point{0, -1},
}

func exploreAndAccumulate(day10MapWithMetadata *Day10MapWithMetadata, x int, y int, endPoints PointSet, sum *int) {
	if day10MapWithMetadata.mapArray[y][x] == 9 {
		endPoints[Point{x, y}] = struct{}{}
		*sum += 1
		return
	}
	nextSlope := day10MapWithMetadata.mapArray[y][x] + 1

	// log.Println("Current position:", x, y, ", Next slope:", nextSlope)
	validDeltas := make([]Point, 0)
	for _, delta := range possibleDeltas {
		newX := x + delta.x
		newY := y + delta.y

		if newX < 0 || newX >= len(day10MapWithMetadata.mapArray[0]) {
			continue
		}
		if newY < 0 || newY >= len(day10MapWithMetadata.mapArray) {
			continue
		}
		if day10MapWithMetadata.mapArray[newY][newX] == nextSlope {
			validDeltas = append(validDeltas, delta)
		}
	}

	// log.Println("Valid deltas:", validDeltas)
	for _, validDelta := range validDeltas {
		// log.Println("Launching exploreAndAccumulate with", x+validDelta.x, y+validDelta.y)
		exploreAndAccumulate(day10MapWithMetadata, x+validDelta.x, y+validDelta.y, endPoints, sum)
	}
}

func day10Part1(inputFilePath *string) int {
	day10MapWithMetadata := LoadDay10Map(inputFilePath)

	// log.Println("Starting points:", day10MapWithMetadata.startPoints)
	sum := 0
	endPointsForStartingPoints := make([]PointSet, 0)
	part2Sum := 0
	for i, startPoint := range day10MapWithMetadata.startPoints {
		endPointsForStartingPoints = append(endPointsForStartingPoints, make(PointSet))
		endPoints := endPointsForStartingPoints[i]
		exploreAndAccumulate(&day10MapWithMetadata, startPoint.x, startPoint.y, endPoints, &part2Sum)
	}

	for _, endPoint := range endPointsForStartingPoints {
		sum += len(endPoint)
	}
	log.Println("Sum:", sum)
	log.Println("Part 2 sum:", part2Sum)
	return 0
}

func Day10(inputFilePath *string) {
	day10Part1(inputFilePath)
}
