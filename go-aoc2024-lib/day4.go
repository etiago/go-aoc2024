package goaoc2024lib

import (
	"log"
)

func stringLinesToCharMatrix(lines []string) [][]rune {
	matrix := make([][]rune, len(lines))
	for i, line := range lines {
		matrix[i] = []rune(line)
	}
	return matrix
}

func getCornersValuesInOrder(x, y int, matrix [][]rune) [4]rune {
	return [4]rune{matrix[y-1][x-1], matrix[y-1][x+1], matrix[y+1][x-1], matrix[y+1][x+1]}
}

func getAdjacencyDeltas() [][]int {
	return [][]int{
		{-1, -1},
		{0, -1},
		{1, -1},
		{-1, 0},
		{1, 0},
		{-1, +1},
		{0, 1},
		{1, 1},
	}
}

func indexIsOutOfBounds(x, y, maxX, maxY int) bool {
	return x < 0 || x >= maxX || y < 0 || y >= maxY
}

// target word is "XMAS", define const
const targetWordPart1 = "XMAS"

func findXmas(x, y int, matrix [][]rune, currentWord string, possibleDeltas [][]int) int {
	lenCurrentWord := len(currentWord)
	nextChar := rune(targetWordPart1[lenCurrentWord])

	if nextChar != matrix[y][x] {
		return 0
	} else {
		newWord := targetWordPart1[:lenCurrentWord+1]
		if newWord == targetWordPart1 {
			return 1
		}

		total := 0
		for _, deltas := range possibleDeltas {
			newX, newY := x+deltas[0], y+deltas[1]
			if indexIsOutOfBounds(newX, newY, len(matrix[0]), len(matrix)) {
				continue
			}

			total += findXmas(newX, newY, matrix, newWord, [][]int{deltas})
		}
		return total
	}
}
func day4Part1(input [][]rune) {
	// Part 1
	total := 0
	for y, line := range input {
		for x := range line {
			adjacencyDeltas := getAdjacencyDeltas()
			total += findXmas(x, y, input, "", adjacencyDeltas)
		}
	}
	log.Println("Total Part 1:", total)
}

var masPossibilties = map[[4]rune]struct{}{
	{'M', 'S', 'M', 'S'}: {},
	{'S', 'M', 'S', 'M'}: {},
	{'S', 'S', 'M', 'M'}: {},
	{'M', 'M', 'S', 'S'}: {},
}

func day4Part2(input [][]rune) {
	total := 0
	// Skip iterating over all the edges since the center
	// of a MAS is always at least 1 unit away from the edges.
	for y := 1; y < len(input)-1; y++ {
		for x := 1; x < len(input[y])-1; x++ {
			// If the current cell is not an 'A', definitely not a MAS.
			if input[y][x] != 'A' {
				continue
			}
			corners := getCornersValuesInOrder(x, y, input)
			if _, ok := masPossibilties[corners]; ok {
				total++
			}
		}
	}
	log.Println("Total Part 2:", total)
}
func Day4(input_file_path *string) {
	lines := ReadFileLines(input_file_path)
	charMatrix := stringLinesToCharMatrix(lines)

	day4Part1(charMatrix)
	day4Part2(charMatrix)
}
