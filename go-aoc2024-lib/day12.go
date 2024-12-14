package goaoc2024lib

import "log"

func getNextValidPositions(x int, y int, char rune, mapArray [][]rune, visited PointSet) []Point {
	deltas := [4]Point{
		Point{1, 0},
		Point{0, 1},
		Point{-1, 0},
		Point{0, -1},
	}
	validPositions := make([]Point, 0)

	for _, delta := range deltas {
		newX := x + delta.x
		newY := y + delta.y

		if newX < 0 || newX >= len(mapArray[0]) {
			continue
		}
		if newY < 0 || newY >= len(mapArray) {
			continue
		}
		if mapArray[newY][newX] == char {
			validPositions = append(validPositions, Point{newX, newY})

		}
	}

	return validPositions
}

func exploreValidMoves(x int, y int, char rune, mapArray [][]rune, visited PointSet, validMoves []Point) {
	// nextValidMoves := make([]Point, 0)
	for _, validMove := range validMoves {
		visited[validMove] = struct{}{}
		mapArray[validMove.y][validMove.x] = '.'
		nextValid := getNextValidPositions(validMove.x, validMove.y, char, mapArray, visited)
		exploreValidMoves(validMove.x, validMove.y, char, mapArray, visited, nextValid)
	}
	// return nextValidMoves
}

func exploreStartingAt(x int, y int, char rune, mapArray [][]rune) (int, PointSet) {
	visited := make(PointSet)
	visited[Point{x, y}] = struct{}{}

	validMoves := getNextValidPositions(x, y, char, mapArray, visited)
	exploreValidMoves(x, y, char, mapArray, visited, validMoves)

	perimeter := perimeterForRegion(visited)

	return perimeter, visited
}

func unionPointSets(set1 PointSet, set2 PointSet) PointSet {
	union := make(PointSet)

	for point := range set1 {
		union[point] = struct{}{}
	}

	for point := range set2 {
		union[point] = struct{}{}
	}

	return union
}

func pointVisitedInAnyRegion(point Point, regions []PointSet) bool {
	for _, region := range regions {
		if _, ok := region[point]; ok {
			return true
		}
	}

	return false
}
func processMapArray(mapArray [][]rune) int {
	total := 0
	for y, row := range mapArray {
		for x, cell := range row {
			if cell == '.' {
				continue
			}
			perimeter, visited := exploreStartingAt(x, y, cell, mapArray)
			total += len(visited) * perimeter
		}
	}

	return total
}

func concavities(point Point, points PointSet) int {
	deltas := [12]Point{
		{-1, 0},
		{0, 1},
		{-1, 1},

		{-1, 0},
		{0, -1},
		{-1, -1},

		{1, 0},
		{0, 1},
		{1, 1},

		{1, 0},
		{0, -1},
		{1, -1},
	}

	concavitiesSum := 0
	for i := 0; i < len(deltas)-2; i += 3 {
		p1 := Point{point.x + deltas[i].x, point.y + deltas[i].y}
		p2 := Point{point.x + deltas[i+1].x, point.y + deltas[i+1].y}
		p3 := Point{point.x + deltas[i+2].x, point.y + deltas[i+2].y}

		hasP1 := false
		hasP2 := false
		hasP3 := false
		if _, ok := points[p1]; !ok {
			hasP1 = true
		}
		if _, ok := points[p2]; !ok {
			hasP2 = true
		}
		if _, ok := points[p3]; !ok {
			hasP3 = true
		}

		if hasP1 && hasP2 && !hasP3 {
			concavitiesSum++
		}

	}

	return concavitiesSum
}

func isTopLeftCorner(point Point, points PointSet) bool {
	deltaLeft := Point{-1, 0}
	deltaTop := Point{0, -1}

	if _, ok := points[Point{point.x + deltaLeft.x, point.y + deltaLeft.y}]; !ok {
		if _, ok := points[Point{point.x + deltaTop.x, point.y + deltaTop.y}]; !ok {
			return true
		}
	}
	return false
}

func isInnerTopLeftCorner(point Point, points PointSet) bool {
	deltaRight := Point{1, 0}
	deltaBottom := Point{0, 1}
	deltaDiagonal := Point{1, 1}

	if _, ok := points[Point{point.x + deltaRight.x, point.y + deltaRight.y}]; ok {
		if _, ok := points[Point{point.x + deltaBottom.x, point.y + deltaBottom.y}]; ok {
			if _, ok := points[Point{point.x + deltaDiagonal.x, point.y + deltaDiagonal.y}]; !ok {
				return true
			}
		}
	}
	return false
}

func isTopRightCorner(point Point, points PointSet) bool {
	deltaRight := Point{1, 0}
	deltaTop := Point{0, -1}

	if _, ok := points[Point{point.x + deltaRight.x, point.y + deltaRight.y}]; !ok {
		if _, ok := points[Point{point.x + deltaTop.x, point.y + deltaTop.y}]; !ok {
			return true
		}
	}
	return false
}

func isInnerTopRightCorner(point Point, points PointSet) bool {
	deltaLeft := Point{-1, 0}
	deltaBottom := Point{0, 1}
	deltaDiagonal := Point{-1, 1}

	if _, ok := points[Point{point.x + deltaLeft.x, point.y + deltaLeft.y}]; ok {
		if _, ok := points[Point{point.x + deltaBottom.x, point.y + deltaBottom.y}]; ok {
			if _, ok := points[Point{point.x + deltaDiagonal.x, point.y + deltaDiagonal.y}]; !ok {
				return true
			}
		}
	}
	return false
}

func isBottomLeftCorner(point Point, points PointSet) bool {
	deltaLeft := Point{-1, 0}
	deltaBottom := Point{0, 1}

	if _, ok := points[Point{point.x + deltaLeft.x, point.y + deltaLeft.y}]; !ok {
		if _, ok := points[Point{point.x + deltaBottom.x, point.y + deltaBottom.y}]; !ok {
			return true
		}
	}
	return false
}

func isInnerBottomLeftCorner(point Point, points PointSet) bool {
	deltaTop := Point{0, -1}
	deltaRight := Point{1, 0}
	deltaDiagonal := Point{1, -1}

	if _, ok := points[Point{point.x + deltaTop.x, point.y + deltaTop.y}]; ok {
		if _, ok := points[Point{point.x + deltaRight.x, point.y + deltaRight.y}]; ok {
			if _, ok := points[Point{point.x + deltaDiagonal.x, point.y + deltaDiagonal.y}]; !ok {
				return true
			}
		}
	}
	return false
}

func isBottomRightCorner(point Point, points PointSet) bool {
	deltaRight := Point{1, 0}
	deltaBottom := Point{0, 1}

	if _, ok := points[Point{point.x + deltaRight.x, point.y + deltaRight.y}]; !ok {
		if _, ok := points[Point{point.x + deltaBottom.x, point.y + deltaBottom.y}]; !ok {
			return true
		}
	}
	return false
}

func isInnerBottomRightCorner(point Point, points PointSet) bool {
	deltaLeft := Point{-1, 0}
	deltaTop := Point{0, -1}
	deltaDiagonal := Point{-1, -1}

	if _, ok := points[Point{point.x + deltaTop.x, point.y + deltaTop.y}]; ok {
		if _, ok := points[Point{point.x + deltaLeft.x, point.y + deltaLeft.y}]; ok {
			if _, ok := points[Point{point.x + deltaDiagonal.x, point.y + deltaDiagonal.y}]; !ok {
				return true
			}
		}
	}
	return false
}

func getSideCount(points PointSet) int {
	// - If the point has exactly two neighbors to the left, right, up or down, it is a side
	// - if the point has 4 neighbours to left, right, up and down, but one of the corners is
	// empty, it is a side

	sidesCount := 0
	for point := range points {
		if isTopLeftCorner(point, points) {
			sidesCount++
		}
		if isTopRightCorner(point, points) {
			sidesCount++
		}
		if isBottomLeftCorner(point, points) {
			sidesCount++
		}
		if isBottomRightCorner(point, points) {
			sidesCount++
		}
		if isInnerTopLeftCorner(point, points) {
			sidesCount++
		}
		if isInnerTopRightCorner(point, points) {
			sidesCount++
		}
		if isInnerBottomLeftCorner(point, points) {
			sidesCount++
		}
		if isInnerBottomRightCorner(point, points) {
			sidesCount++
		}
		// neighborsLeftRight := 0
		// neighborsTopDown := 0

		// deltasLeftRight := [2]Point{
		// 	{-1, 0},
		// 	{1, 0},
		// }

		// deltasTopDown := [2]Point{
		// 	{0, 1},
		// 	{0, -1},
		// }

		// for _, delta := range deltasLeftRight {
		// 	newPoint := Point{point.x + delta.x, point.y + delta.y}
		// 	if _, ok := points[newPoint]; ok {
		// 		neighborsLeftRight++
		// 	}
		// }
		// for _, delta := range deltasTopDown {
		// 	newPoint := Point{point.x + delta.x, point.y + delta.y}
		// 	if _, ok := points[newPoint]; ok {
		// 		neighborsTopDown++
		// 	}
		// }

		// neighbors := neighborsLeftRight + neighborsTopDown
		// if neighbors == 0 {
		// 	log.Println("Position xy contributed 4 sides ", point)
		// 	sidesCount += 4
		// 	continue
		// }
		// concavitiesSum := concavities(point, points)
		// sidesCount += concavitiesSum * 2
		// if neighborsLeftRight == 1 && neighborsTopDown == 1 {
		// 	log.Println("Position xy contributed 1 sides ", point)
		// 	sidesCount++
		// 	continue
		// }

	}

	return sidesCount
}
func processMapArrayPart2(mapArray [][]rune) int {
	total := 0
	for y, row := range mapArray {
		for x, cell := range row {
			if cell == '.' {
				continue
			}
			_, visited := exploreStartingAt(x, y, cell, mapArray)

			sidesCount := getSideCount(visited)

			total += len(visited) * sidesCount
		}
	}

	return total
}

func perimeterSides(region PointSet, point Point) int {
	perimeter := 0
	deltas := [4]Point{
		Point{1, 0},
		Point{0, 1},
		Point{-1, 0},
		Point{0, -1},
	}

	for _, delta := range deltas {
		newPoint := Point{point.x + delta.x, point.y + delta.y}
		if _, ok := region[newPoint]; !ok {
			perimeter++
		}
	}

	return perimeter
}

func perimeterForRegion(region PointSet) int {
	perimeter := 0
	for point := range region {
		perimeter += perimeterSides(region, point)
	}

	return perimeter
}

func day12Part1(inputFilePath *string) int {
	mapArray := ReadFileLinesAsRuneMatrix(inputFilePath)

	cost := processMapArray(mapArray)

	log.Println("Cost:", cost)
	return 0
}

func day12Part2(inputFilePath *string) {
	mapArray := ReadFileLinesAsRuneMatrix(inputFilePath)

	cost := processMapArrayPart2(mapArray)

	log.Println("Cost:", cost)

}

func Day12(inputFilePath *string) {
	day12Part1(inputFilePath)
	day12Part2(inputFilePath)
}
