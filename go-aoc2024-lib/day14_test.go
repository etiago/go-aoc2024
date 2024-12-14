package goaoc2024lib

import "testing"

func TestMoveRobotExample1(t *testing.T) {
	robot := Robot{
		position: Point{2, 4},
		velocity: Point{2, -3},
	}
	expected := Point{1, 3}
	moveRobotNTimes(&robot, 5, 11, 7)
	if robot.position.x != expected.x || robot.position.y != expected.y {
		t.Errorf("Expected %d, %d, got %d, %d", expected.x, expected.y, robot.position.x, robot.position.y)
	}
}

func TestMoveRobotExample1MoreMoves(t *testing.T) {
	robot := Robot{
		position: Point{2, 4},
		velocity: Point{2, -3},
	}
	expected := Point{10, 6}
	moveRobotNTimes(&robot, 4, 11, 7)
	if robot.position.x != expected.x || robot.position.y != expected.y {
		t.Errorf("Expected %d, %d, got %d, %d", expected.x, expected.y, robot.position.x, robot.position.y)
	}
}

func TestMoveRobotAnotherRobot(t *testing.T) {
	robot := Robot{
		position: Point{9, 3},
		velocity: Point{2, 3},
	}
	moves := 4
	expected := Point{6, 1}
	moveRobotNTimes(&robot, moves, 11, 7)
	if robot.position.x != expected.x || robot.position.y != expected.y {
		t.Errorf("Expected %d, %d, got %d, %d", expected.x, expected.y, robot.position.x, robot.position.y)
	}
}

func TestCountsPerQuadrantExample1(t *testing.T) {
	path := "../input/day14-example.txt"
	robots := LoadDay14Robots(&path)

	// newRobots := make([]Robot, 0)
	// for _, robot := range robots {
	// 	moveRobotNTimes(&robot, 100, 11, 7)
	// }

	for i := 0; i < len(robots); i++ {
		moveRobotNTimes(&robots[i], 100, 11, 7)
	}

	expected := [4]int{1, 3, 4, 1}
	counts := getCountsPerQuadrant(robots, 11, 7)
	if counts != expected {
		t.Errorf("Expected %v, got %v", expected, counts)
	}
}

func TestPrintRobotsAsMatrixIfGoodCandidate(t *testing.T) {

}
