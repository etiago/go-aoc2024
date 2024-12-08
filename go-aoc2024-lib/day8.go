package goaoc2024lib

import (
	"log"
	"math"

	"gonum.org/v1/gonum/stat/combin"
)

type PointSet map[Point]struct{}

func getAntinodes(pointA Point, pointB Point) [2]Point {
	deltaX := pointB.x - pointA.x
	deltaY := pointB.y - pointA.y

	absDeltaX := math.Abs(float64(deltaX))
	absDeltaY := math.Abs(float64(deltaY))

	antinodeA := Point{-1, -1}
	antinodeB := Point{-1, -1}

	if deltaX < 0 {
		antinodeA.x = pointA.x + int(absDeltaX)
		antinodeB.x = pointB.x - int(absDeltaX)
	} else {
		antinodeA.x = pointA.x - int(absDeltaX)
		antinodeB.x = pointB.x + int(absDeltaX)
	}

	if deltaY < 0 {
		antinodeA.y = pointA.y + int(absDeltaY)
		antinodeB.y = pointB.y - int(absDeltaY)
	} else {
		antinodeA.y = pointA.y - int(absDeltaY)
		antinodeB.y = pointB.y + int(absDeltaY)
	}

	return [2]Point{antinodeA, antinodeB}
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func gcd(a, b int) int {
	// GCD using Euclidean algorithm
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func getAntinodesPart2(pointA Point, pointB Point, xLength int, yLength int) map[Point]struct{} {
	deltaX := abs(pointB.x - pointA.x)
	deltaY := abs(pointB.y - pointA.y)

	divisor := gcd(deltaX, deltaY)
	normalizedDeltaX := deltaX / divisor
	normalizedDeltaY := deltaY / divisor

	antinodePoints := make(map[Point]struct{})
	antinodePoints[pointA] = struct{}{}
	antinodePoints[pointB] = struct{}{}

	pointAmultiplierX := 1
	pointAmultiplierY := 1
	if pointA.x < pointB.x {
		pointAmultiplierX = -1
	}
	if pointA.y < pointB.y {
		pointAmultiplierY = -1
	}

	pointBmultiplierX := 1
	pointBmultiplierY := 1
	if pointB.x < pointA.x {
		pointBmultiplierX = -1
	}
	if pointB.y < pointA.y {
		pointBmultiplierY = -1
	}

	// First lets expand away from pointA
	newPointA := Point{pointA.x, pointA.y}
	newPointB := Point{pointB.x, pointB.y}
	pointInMap := true
	for pointInMap {
		newPointA.x = newPointA.x + (normalizedDeltaX * pointAmultiplierX)
		newPointA.y = newPointA.y + (normalizedDeltaY * pointAmultiplierY)

		if isAntinodeInMap(xLength, yLength, newPointA) {
			antinodePoints[newPointA] = struct{}{}
		} else {
			pointInMap = false
		}
	}

	pointInMap = true
	for pointInMap {
		newPointB.x = newPointB.x + (normalizedDeltaX * pointBmultiplierX)
		newPointB.y = newPointB.y + (normalizedDeltaY * pointBmultiplierY)

		if isAntinodeInMap(xLength, yLength, newPointB) {
			antinodePoints[newPointB] = struct{}{}
		} else {
			pointInMap = false
		}
	}

	return antinodePoints
}

func isAntinodeInMap(xLength int, yLength int, antinode Point) bool {
	return antinode.x >= 0 && antinode.x < xLength && antinode.y >= 0 && antinode.y < yLength
}

func day8Part1(inputFilePath *string) {
	antennaMapWithMetadata := LoadDay8Map(inputFilePath)
	xLength := len(antennaMapWithMetadata.mapArray[0])
	yLength := len(antennaMapWithMetadata.mapArray)

	// set for antinodes
	antinodeSet := make(PointSet)
	for _, antennaPoints := range antennaMapWithMetadata.antennaToLocations {
		antennaPairPermutationsIndices := combin.Permutations(len(antennaPoints), 2)

		for _, pairIndex := range antennaPairPermutationsIndices {
			pointA := antennaPoints[pairIndex[0]]
			pointB := antennaPoints[pairIndex[1]]

			newAntinodes := getAntinodes(pointA, pointB)

			if isAntinodeInMap(xLength, yLength, newAntinodes[0]) {
				antinodeSet[newAntinodes[0]] = struct{}{}
			}
			if isAntinodeInMap(xLength, yLength, newAntinodes[1]) {
				antinodeSet[newAntinodes[1]] = struct{}{}
			}
		}
	}

	log.Println("Antinodes:", len(antinodeSet))
}

func day8Part2(inputFilePath *string) {
	antennaMapWithMetadata := LoadDay8Map(inputFilePath)
	xLength := len(antennaMapWithMetadata.mapArray[0])
	yLength := len(antennaMapWithMetadata.mapArray)

	antinodeSet := make(PointSet)
	for _, antennaPoints := range antennaMapWithMetadata.antennaToLocations {
		antennaPairCombinationsIndices := combin.Combinations(len(antennaPoints), 2)

		for _, pairIndex := range antennaPairCombinationsIndices {
			pointA := antennaPoints[pairIndex[0]]
			pointB := antennaPoints[pairIndex[1]]

			newAntinodes := getAntinodesPart2(pointA, pointB, xLength, yLength)

			for antinode := range newAntinodes {
				antinodeSet[antinode] = struct{}{}
			}
		}
	}

	log.Println("Antinodes:", len(antinodeSet))
}

func Day8(inputFilePath *string) {
	day8Part1(inputFilePath)
	day8Part2(inputFilePath)
}
