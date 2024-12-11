package goaoc2024lib

import (
	"log"
	"strconv"
)

const multiplier = uint64(2024)

// On part 2, realised this wasn't going to work.
// Made a new implementation for part 2 which is better in general, so also using it for part 1.
// But leaving this here for reference.
func iterateStones(stones []Stone, times int) []Stone {
	newStones := make([]Stone, 0)
	for _, stone := range stones {
		if stone == 0 {
			newStones = append(newStones, 1)
			continue
		}

		stoneStr := strconv.FormatUint(uint64(stone), 10)
		if len(stoneStr)%2 == 0 {
			// left stone with half of the digits, right stone with the other half
			halfLength := len(stoneStr) / 2
			leftStoneStr := stoneStr[:halfLength]
			rightStoneStr := stoneStr[halfLength:]

			// parse left
			leftStone, _ := strconv.ParseUint(leftStoneStr, 10, 64)
			newStones = append(newStones, Stone(leftStone))
			rightStone, _ := strconv.ParseUint(rightStoneStr, 10, 64)
			newStones = append(newStones, Stone(rightStone))

			continue
		}

		multiplied := uint64(stone) * multiplier
		newStones = append(newStones, Stone(multiplied))
	}

	if times == 0 {
		return newStones
	}

	return iterateStones(newStones, times-1)

}

func iterateStoneCounts(stones map[Stone]uint64, times int) map[Stone]uint64 {
	newStones := make(map[Stone]uint64)
	for stone, count := range stones {
		if stone == 0 {
			newStones[1] += count
			continue
		}

		stoneStr := strconv.FormatUint(uint64(stone), 10)
		if len(stoneStr)%2 == 0 {
			// left stone with half of the digits, right stone with the other half
			halfLength := len(stoneStr) / 2
			leftStoneStr := stoneStr[:halfLength]
			rightStoneStr := stoneStr[halfLength:]

			// parse left
			leftStone, _ := strconv.ParseUint(leftStoneStr, 10, 64)
			newStones[Stone(leftStone)] += count
			rightStone, _ := strconv.ParseUint(rightStoneStr, 10, 64)
			newStones[Stone(rightStone)] += count

			continue
		}

		multiplied := uint64(stone) * multiplier
		newStones[Stone(multiplied)] += count
	}

	if times == 0 {
		return newStones
	}

	return iterateStoneCounts(newStones, times-1)
}

func day11Part1(inputFilePath *string) int {
	stones := LoadDay11Stones(inputFilePath)

	stoneCounts := make(map[Stone]uint64)
	for _, stone := range stones {
		stoneCounts[stone]++
	}

	newStoneCounts := iterateStoneCounts(stoneCounts, 24)
	sum := uint64(0)
	for _, count := range newStoneCounts {
		sum += count
	}
	log.Println("Stone counts:", sum)
	return 0
}

func day11Part2(inputFilePath *string) int {
	stones := LoadDay11Stones(inputFilePath)

	stoneCounts := make(map[Stone]uint64)
	for _, stone := range stones {
		stoneCounts[stone]++
	}

	newStoneCounts := iterateStoneCounts(stoneCounts, 74)
	sum := uint64(0)
	for _, count := range newStoneCounts {
		sum += count
	}
	log.Println("Stone counts:", sum)
	return 0
}

func Day11(inputFilePath *string) {
	day11Part1(inputFilePath)
	day11Part2(inputFilePath)
}
