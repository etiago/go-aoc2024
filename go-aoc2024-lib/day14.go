package goaoc2024lib

import (
	"log"
)

func mathModulo(a int, b int) int {
	return ((a % b) + b) % b
}

func moveRobotNTimes(robot *Robot, times int, mapWidth int, mapHeight int) {
	newX := mathModulo((robot.position.x + (times * robot.velocity.x)), mapWidth)
	newY := mathModulo((robot.position.y + (times * robot.velocity.y)), mapHeight)

	robot.position = Point{newX, newY}
}

func getCountsPerQuadrant(robots []Robot, mapWidth int, mapHeight int) [4]int {
	countsPerQuadrant := [4]int{0, 0, 0, 0}
	midMapX := mapWidth / 2
	midMapY := mapHeight / 2

	widthIsOdd := mapWidth%2 != 0
	heightIsOdd := mapHeight%2 != 0
	padMiddleX := 0
	padMiddleY := 0
	if mapWidth%2 != 0 {
		padMiddleX = 1
	}
	if mapHeight%2 != 0 {
		padMiddleY = 1
	}

	for _, robot := range robots {
		// If robot is exactly in the middle on X or Y, skip
		if (robot.position.x == midMapX && widthIsOdd) || (robot.position.y == midMapY && heightIsOdd) {
			continue
		}
		quadrant := 0
		if robot.position.x < midMapX && robot.position.y < midMapY {
			quadrant = 0
		} else if robot.position.x >= midMapX+padMiddleX && robot.position.y < midMapY {
			quadrant = 1
		} else if robot.position.x < midMapX && robot.position.y >= midMapY+padMiddleY {
			quadrant = 2
		} else if robot.position.x >= midMapX+padMiddleX && robot.position.y >= midMapY+padMiddleY {
			quadrant = 3
		}
		countsPerQuadrant[quadrant]++
	}

	return countsPerQuadrant
}
func day14Part1(inputFilePath *string) {
	robots := LoadDay14Robots(inputFilePath)

	for i := 0; i < len(robots); i++ {
		moveRobotNTimes(&robots[i], 100, 101, 103)
	}
	quadrants := getCountsPerQuadrant(robots, 101, 103)
	safetyFactor := quadrants[0] * quadrants[1] * quadrants[2] * quadrants[3]

	log.Println("Safety factor:", safetyFactor)

	robots = LoadDay14Robots(inputFilePath)

	// var input string
	times := 0
	for times < 100000000 {
		for i := 0; i < len(robots); i++ {
			moveRobotNTimes(&robots[i], 1, 101, 103)
		}
		found := printRobotsAsMatrixIfGoodCandidate(robots, 101, 103)
		times++
		if found {
			log.Println("Times:", times)
			return
		}
	}
}

func hasVerticalLineLengthN(robotSet map[Point]bool, robot Point, n int) bool {
	for i := 1; i < n; i++ {
		if _, ok := robotSet[Point{robot.x, robot.y + i}]; !ok {
			return false
		}
	}
	return true
}
func printRobotsAsMatrixIfGoodCandidate(robots []Robot, mapWidth int, mapHeight int) bool {
	robotSet := make(map[Point]bool)
	for _, robot := range robots {
		robotSet[robot.position] = true
	}

	// Hated this :( but don't care enough about AoC to make it better.
	// What a horrible part 2 this was.
	found := false
	for robot := range robotSet {
		if hasVerticalLineLengthN(robotSet, robot, 10) {
			found = true
		}
	}

	if !found {
		return false
	}

	matrix := make([][]rune, mapHeight)
	matrixCounts := make([][]int, mapHeight)
	for i := 0; i < mapHeight; i++ {
		matrix[i] = make([]rune, mapWidth)
		matrixCounts[i] = make([]int, mapWidth)
		for j := 0; j < mapWidth; j++ {
			matrix[i][j] = '.'
		}
	}

	for _, robot := range robots {
		matrixCounts[robot.position.y][robot.position.x]++
		matrix[robot.position.y][robot.position.x] = '#'
	}

	for i := 0; i < mapHeight; i++ {
		log.Println(string(matrix[i]))
	}
	return true
}

func Day14(inputFilePath *string) {
	day14Part1(inputFilePath)
}
