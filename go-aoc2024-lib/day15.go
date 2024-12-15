package goaoc2024lib

import (
	"log"
	"sort"
)

func movementToDelta(movement Movement) Point {
	switch movement {
	case Up:
		return Point{0, -1}
	case Down:
		return Point{0, 1}
	case Right:
		return Point{1, 0}
	case Left:
		return Point{-1, 0}
	}
	panic("Invalid movement")
}

func movementHitsWall(movement Movement, puzzle *Day15Map) bool {
	delta := movementToDelta(movement)
	newPosition := Point{puzzle.robotPosition.x + delta.x, puzzle.robotPosition.y + delta.y}
	return puzzle.mapArray[newPosition.y][newPosition.x] == '#'
}

func tryMoveToEmptySpace(movement Movement, puzzle *Day15Map) bool {
	delta := movementToDelta(movement)
	newPosition := Point{puzzle.robotPosition.x + delta.x, puzzle.robotPosition.y + delta.y}
	if puzzle.mapArray[newPosition.y][newPosition.x] == '.' {
		puzzle.mapArray[puzzle.robotPosition.y][puzzle.robotPosition.x] = '.'
		puzzle.mapArray[newPosition.y][newPosition.x] = '@'
		puzzle.robotPosition = newPosition
		return true
	}
	return false
}

func tryPushBoxesAndMove(movement Movement, puzzle *Day15Map) {
	delta := movementToDelta(movement)
	currentPosition := Point{puzzle.robotPosition.x + delta.x, puzzle.robotPosition.y + delta.y}
	boxesToMove := 0
	for puzzle.mapArray[currentPosition.y][currentPosition.x] == 'O' {
		boxesToMove++
		currentPosition = Point{currentPosition.x + delta.x, currentPosition.y + delta.y}
	}

	if puzzle.mapArray[currentPosition.y][currentPosition.x] == '.' {
		for i := 0; i < boxesToMove; i++ {
			puzzle.mapArray[currentPosition.y][currentPosition.x] = 'O'
			currentPosition = Point{currentPosition.x - delta.x, currentPosition.y - delta.y}
		}
		puzzle.mapArray[puzzle.robotPosition.y][puzzle.robotPosition.x] = '.'
		puzzle.mapArray[currentPosition.y][currentPosition.x] = '@'
		puzzle.robotPosition = currentPosition
	}
}

type BoxPart2 [2]Point
type BoxSet map[BoxPart2]bool

func getAllBoxesPushedByBoxLaterally(boxesCanBePushed map[BoxPart2]bool, beingPushed []BoxPart2, puzzle *Day15Map, delta Point) {
	newBeingPushed := make([]BoxPart2, 0)
	for _, box := range beingPushed {
		boxNewPositions := BoxPart2{Point{box[0].x + delta.x, box[0].y + delta.y}, Point{box[1].x + delta.x, box[1].y + delta.y}}

		if delta.x == 1 && puzzle.mapArray[boxNewPositions[1].y][boxNewPositions[1].x] == '#' {
			boxesCanBePushed[box] = false
			continue
		}
		if delta.x == -1 && puzzle.mapArray[boxNewPositions[0].y][boxNewPositions[0].x] == '#' {
			boxesCanBePushed[box] = false
			continue
		}
		if delta.x == 1 && puzzle.mapArray[boxNewPositions[1].y][boxNewPositions[1].x] == '.' {
			boxesCanBePushed[box] = true
		}
		if delta.x == -1 && puzzle.mapArray[boxNewPositions[0].y][boxNewPositions[0].x] == '.' {
			boxesCanBePushed[box] = true
		}
		if delta.x == 1 && puzzle.mapArray[boxNewPositions[1].y][boxNewPositions[1].x] == '[' {
			newBeingPushed = append(newBeingPushed, BoxPart2{Point{boxNewPositions[1].x, boxNewPositions[1].y}, Point{boxNewPositions[1].x + 1, boxNewPositions[1].y}})
			boxesCanBePushed[box] = true
		}
		if delta.x == -1 && puzzle.mapArray[boxNewPositions[0].y][boxNewPositions[0].x] == ']' {
			newBeingPushed = append(newBeingPushed, BoxPart2{Point{boxNewPositions[0].x - 1, boxNewPositions[0].y}, Point{boxNewPositions[0].x, boxNewPositions[0].y}})
			boxesCanBePushed[box] = true
		}
	}

	if len(newBeingPushed) > 0 {
		getAllBoxesPushedByBoxLaterally(boxesCanBePushed, newBeingPushed, puzzle, delta)
	}
}
func getAllBoxesPushedByBox(boxesCanBePushed map[BoxPart2]bool, beingPushed []BoxPart2, puzzle *Day15Map, delta Point) {
	newBeingPushed := make([]BoxPart2, 0)
	for _, box := range beingPushed {
		boxNewPositions := BoxPart2{Point{box[0].x + delta.x, box[0].y + delta.y}, Point{box[1].x + delta.x, box[1].y + delta.y}}

		if puzzle.mapArray[boxNewPositions[0].y][boxNewPositions[0].x] == '.' && puzzle.mapArray[boxNewPositions[1].y][boxNewPositions[1].x] == '.' {
			boxesCanBePushed[box] = true
		}

		if puzzle.mapArray[boxNewPositions[0].y][boxNewPositions[0].x] == '#' || puzzle.mapArray[boxNewPositions[1].y][boxNewPositions[1].x] == '#' {
			boxesCanBePushed[box] = false
			continue
		}
		if puzzle.mapArray[boxNewPositions[0].y][boxNewPositions[0].x] == '[' {
			newBeingPushed = append(newBeingPushed, BoxPart2{Point{boxNewPositions[0].x, boxNewPositions[0].y}, Point{boxNewPositions[1].x, boxNewPositions[1].y}})
			boxesCanBePushed[box] = true
		}
		if puzzle.mapArray[boxNewPositions[0].y][boxNewPositions[0].x] == ']' {
			newBeingPushed = append(newBeingPushed, BoxPart2{Point{boxNewPositions[0].x - 1, boxNewPositions[0].y}, Point{boxNewPositions[0].x, boxNewPositions[0].y}})
			boxesCanBePushed[box] = true
		}

		if puzzle.mapArray[boxNewPositions[1].y][boxNewPositions[1].x] == '[' {
			newBeingPushed = append(newBeingPushed, BoxPart2{Point{boxNewPositions[1].x, boxNewPositions[1].y}, Point{boxNewPositions[1].x + 1, boxNewPositions[1].y}})
			boxesCanBePushed[box] = true
		}
	}

	if len(newBeingPushed) > 0 {
		getAllBoxesPushedByBox(boxesCanBePushed, newBeingPushed, puzzle, delta)
	}
}

func moveAllBoxes(boxes map[BoxPart2]bool, delta Point, puzzle *Day15Map) {
	boxesForMove := make([]BoxPart2, 0)

	// sort boxes by y coordinate, from lower y to larger y
	for box := range boxes {
		boxesForMove = append(boxesForMove, box)
	}

	sort.Slice(boxesForMove, func(i, j int) bool {
		if delta.y == -1 {
			return boxesForMove[i][0].y < boxesForMove[j][0].y
		} else if delta.y == 1 {
			return boxesForMove[i][0].y > boxesForMove[j][0].y
		} else if delta.x == -1 {
			return boxesForMove[i][0].x < boxesForMove[j][0].x
		} else if delta.x == 1 {
			return boxesForMove[i][0].x > boxesForMove[j][0].x
		}
		return false
	})

	for _, box := range boxesForMove {
		puzzle.mapArray[box[0].y][box[0].x] = '.'
		puzzle.mapArray[box[1].y][box[1].x] = '.'
		puzzle.mapArray[box[0].y+delta.y][box[0].x+delta.x] = '['
		puzzle.mapArray[box[1].y+delta.y][box[1].x+delta.x] = ']'
	}
}

type MovementFunc func(map[BoxPart2]bool, []BoxPart2, *Day15Map, Point)

var movementFuncs = map[Movement]MovementFunc{
	Up:    getAllBoxesPushedByBox,
	Down:  getAllBoxesPushedByBox,
	Left:  getAllBoxesPushedByBoxLaterally,
	Right: getAllBoxesPushedByBoxLaterally,
}

func tryPushBoxesAndMoveExpandedPuzzle(movement Movement, puzzle *Day15Map) map[BoxPart2]bool {
	delta := movementToDelta(movement)

	boxesCanBePushed := make(map[BoxPart2]bool)

	beingPushed := make([]BoxPart2, 0)
	nextRobotPosition := Point{puzzle.robotPosition.x + delta.x, puzzle.robotPosition.y + delta.y}

	// Get the boxes being pushed directly by the robot
	if puzzle.mapArray[nextRobotPosition.y][nextRobotPosition.x] == '[' {
		beingPushed = append(beingPushed, BoxPart2{Point{nextRobotPosition.x, nextRobotPosition.y}, Point{nextRobotPosition.x + 1, nextRobotPosition.y}})
	}
	if puzzle.mapArray[nextRobotPosition.y][nextRobotPosition.x] == ']' {
		beingPushed = append(beingPushed, BoxPart2{Point{nextRobotPosition.x - 1, nextRobotPosition.y}, Point{nextRobotPosition.x, nextRobotPosition.y}})
	}

	movementFuncs[movement](boxesCanBePushed, beingPushed, puzzle, delta)

	allBoxesCanBePushed := true
	for _, canBePushed := range boxesCanBePushed {
		if !canBePushed {
			allBoxesCanBePushed = false
			break
		}
	}

	if allBoxesCanBePushed {
		moveAllBoxes(boxesCanBePushed, delta, puzzle)
	}

	// Actually move the robot if possible.
	if puzzle.mapArray[nextRobotPosition.y][nextRobotPosition.x] == '.' {
		puzzle.mapArray[puzzle.robotPosition.y][puzzle.robotPosition.x] = '.'
		puzzle.mapArray[nextRobotPosition.y][nextRobotPosition.x] = '@'
		puzzle.robotPosition = nextRobotPosition
	}

	return boxesCanBePushed
}

func day15Part1(inputFilePath *string) int {
	puzzle := LoadDay15Map(inputFilePath, false)

	for _, movement := range puzzle.movements {
		if movementHitsWall(movement, &puzzle) {
			continue
		}

		if tryMoveToEmptySpace(movement, &puzzle) {
			continue
		}

		tryPushBoxesAndMove(movement, &puzzle)
	}

	coordinateSum := 0
	for y, row := range puzzle.mapArray {
		for x, cell := range row {
			if cell == 'O' {
				coordinateSum += y*100 + x
			}
		}
	}
	log.Println("Coordinate sum:", coordinateSum)
	return coordinateSum
}

func day15Part2(inputFilePath *string) int {
	puzzle := LoadDay15Map(inputFilePath, true)

	for _, movement := range puzzle.movements {
		tryPushBoxesAndMoveExpandedPuzzle(movement, &puzzle)
	}

	coordinateSum := 0
	for y, row := range puzzle.mapArray {
		for x, cell := range row {
			if cell == '[' {
				coordinateSum += y*100 + x
			}
		}
	}
	log.Println("Coordinate sum:", coordinateSum)
	return coordinateSum
}
func Day15(inputFilePath *string) int {
	day15Part1(inputFilePath)
	day15Part2(inputFilePath)
	return 0
}
