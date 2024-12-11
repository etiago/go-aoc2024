package goaoc2024lib

import (
	"log"
	"os"
	"strconv"
	"strings"
)

func ReadFile(input_file_path *string) string {
	content, err := os.ReadFile(*input_file_path)
	if err != nil {
		log.Fatal(err)
	}

	return string(content)
}

func ReadFileLines(input_file_path *string) []string {
	content := ReadFile(input_file_path)
	return strings.Split(content, "\n")
}

func LoadDay5Rules(input_file_path *string) map[int]map[int]struct{} {
	content := ReadFileLines(input_file_path)

	rules := make(map[int]map[int]struct{})
	for _, line := range content {
		if line == "" {
			break
		}

		parts := strings.Split(line, "|")

		left, _ := strconv.Atoi(parts[0])
		right, _ := strconv.Atoi(parts[1])

		if _, ok := rules[left]; !ok {
			rules[left] = make(map[int]struct{})
			rules[left][right] = struct{}{}
		} else {
			rules[left][right] = struct{}{}
		}
	}

	return rules
}

func LoadDay5Updates(input_file_path *string) [][]int {
	content := ReadFileLines(input_file_path)

	skipRules := true
	updates := make([][]int, 0)

	for _, line := range content {
		if line != "" && skipRules {
			continue
		}

		if line == "" {
			skipRules = false
			continue
		}

		parts := strings.Split(line, ",")
		update := make([]int, 0)
		for part := range parts {
			num, _ := strconv.Atoi(parts[part])
			update = append(update, num)
		}

		updates = append(updates, update)
	}

	return updates
}

type GuardPosition struct {
	x int
	y int
}

type GuardPositionWithOrientation struct {
	guardPosition GuardPosition
	orientation   Orientation
}

func LoadDay6Map(inputFilePath *string) ([][]rune, *GuardPosition) {
	content := ReadFileLines(inputFilePath)

	mapArray := make([][]rune, len(content))
	var guardPosition *GuardPosition = nil

	for i, line := range content {
		mapArray[i] = []rune(line)

		maybeGuardX := strings.Index(line, "^")
		if maybeGuardX != -1 {
			guardPosition = &GuardPosition{maybeGuardX, i}
		}
	}

	return mapArray, guardPosition
}

type Equation struct {
	result   int64
	operands []int64
}

func LoadDay7Equations(inputFilePath *string) []Equation {
	content := ReadFileLines(inputFilePath)

	equations := make([]Equation, len(content))
	for i, line := range content {
		chunks := strings.Split(line, ": ")
		result, _ := strconv.ParseInt(chunks[0], 10, 64)

		operands := strings.Split(chunks[1], " ")
		operandsInt := make([]int64, len(operands))
		for i, operand := range operands {
			operandsInt[i], _ = strconv.ParseInt(operand, 10, 64)
		}

		equation := Equation{result, operandsInt}
		equations[i] = equation
	}
	return equations
}

type Point struct {
	x int
	y int
}
type AntennaMapWithMetadata struct {
	mapArray           [][]rune
	antennaToLocations map[rune][]Point
}
type AtennaMapWithAntinodeSet struct {
	mapArray    [][]rune
	antinodeSet map[Point]struct{}
}

func (antennaMapWithMetadata AntennaMapWithMetadata) String() string {
	var sb strings.Builder
	sb.WriteString("-------\n")
	for _, line := range antennaMapWithMetadata.mapArray {
		sb.WriteString(string(line))
		sb.WriteString("\n")
	}
	sb.WriteString("-------\n")
	return sb.String()
}

func (antennaMapWithAntinodeSet AtennaMapWithAntinodeSet) String() string {
	var sb strings.Builder
	sb.WriteString("-------\n")
	for y, line := range antennaMapWithAntinodeSet.mapArray {
		for x, char := range line {
			if _, ok := antennaMapWithAntinodeSet.antinodeSet[Point{x, y}]; ok && char == '.' {
				sb.WriteString("#")
			} else {
				sb.WriteString(string(char))
			}
		}
		sb.WriteString("\n")
	}
	sb.WriteString("-------\n")
	return sb.String()
}

func LoadDay8Map(inputFilePath *string) AntennaMapWithMetadata {
	content := ReadFileLines(inputFilePath)
	antennaToLocations := make(map[rune][]Point)

	mapArray := make([][]rune, len(content))
	for i, line := range content {
		mapArray[i] = []rune(line)

		for j, char := range line {
			if char != '.' {
				if _, ok := antennaToLocations[char]; !ok {
					antennaToLocations[char] = make([]Point, 0)
				}
				antennaToLocations[char] = append(antennaToLocations[char], Point{j, i})
			}
		}
	}

	return AntennaMapWithMetadata{mapArray, antennaToLocations}
}

type Block struct {
	id      int
	indices []int
}

type Disk struct {
	dataBlocks []Block
	freeBlock  Block
	length     int
}

func (disk Disk) String() string {
	retVal := make([]rune, disk.length)

	for _, block := range disk.dataBlocks {
		for _, index := range block.indices {
			retVal[index] = '0' + rune(block.id)
		}

		for _, index := range disk.freeBlock.indices {
			retVal[index] = '.'
		}
	}

	return string(retVal)
}

func LoadDay9Disk(inputFilePath *string) Disk {
	content := ReadFile(inputFilePath)

	dataBlocks := make([]Block, 0)
	freeBlock := Block{-1, make([]int, 0)}
	index := 0
	dataBlockIndex := 0
	for i := 0; i < len(content); i++ {
		dataChar := content[i]
		intChar, _ := strconv.Atoi(string(dataChar))
		block := Block{dataBlockIndex, make([]int, 0)}

		for j := index; j < index+intChar; j++ {
			block.indices = append(block.indices, j)
		}

		dataBlocks = append(dataBlocks, block)
		dataBlockIndex++

		index += intChar

		i++

		if i >= len(content) {
			break
		}

		spaceChar := content[i]
		intSpaceChar, _ := strconv.Atoi(string(spaceChar))

		for j := index; j < index+intSpaceChar; j++ {
			freeBlock.indices = append(freeBlock.indices, j)
		}

		index += intSpaceChar
	}

	log.Println("Data blocks:", dataBlocks)

	return Disk{dataBlocks, freeBlock, index + 1}
}

const ImpassableTile = -2

type Day10Map [][]int

type Day10MapWithMetadata struct {
	mapArray    Day10Map
	startPoints []Point
}

func LoadDay10Map(inputFilePath *string) Day10MapWithMetadata {
	content := ReadFileLines(inputFilePath)

	mapArray := make(Day10Map, len(content))
	startPoints := make([]Point, 0)
	for i, line := range content {
		mapArray[i] = make([]int, len(line))
		for j, char := range line {
			num, err := strconv.Atoi(string(char))
			if err != nil {
				mapArray[i][j] = ImpassableTile
				continue
			}
			if num == 0 {
				startPoints = append(startPoints, Point{j, i})
			}
			mapArray[i][j] = num
		}
	}

	return Day10MapWithMetadata{mapArray, startPoints}
}

type Stone uint64

func LoadDay11Stones(inputFilePath *string) []Stone {
	content := ReadFile(inputFilePath)

	stoneStrs := strings.Split(content, " ")

	stones := make([]Stone, len(stoneStrs))
	for i, stoneStr := range stoneStrs {
		stone, _ := strconv.ParseUint(stoneStr, 10, 64)
		stones[i] = Stone(stone)
	}
	return stones
}
