package goaoc2024lib

import "log"

func isValidUpdate(update []int, rules map[int]map[int]struct{}) bool {
	seen := make(map[int]struct{})

	for _, updateValue := range update {
		ruleSetForValue := rules[updateValue]
		for ruleValue := range ruleSetForValue {
			if _, ok := seen[ruleValue]; ok {
				return false
			}
		}

		seen[updateValue] = struct{}{}
	}

	return true
}

func isValidUpdateWithInvalidIndex(update []int, rules map[int]map[int]struct{}) (int, bool) {
	seen := make(map[int]struct{})

	for i, updateValue := range update {
		ruleSetForValue := rules[updateValue]
		for ruleValue := range ruleSetForValue {
			if _, ok := seen[ruleValue]; ok {
				return i, false
			}
		}

		seen[updateValue] = struct{}{}
	}

	return -1, true
}

func day5Part1(inputFilePath *string) int {
	log.Println("Day 5 Part 1")
	rules := LoadDay5Rules(inputFilePath)
	updates := LoadDay5Updates(inputFilePath)

	goodUpdates := make([][]int, 0)
	for _, update := range updates {
		isValid := isValidUpdate(update, rules)
		if isValid {
			goodUpdates = append(goodUpdates, update)
		}

	}

	sumMiddleValues := 0

	for _, update := range goodUpdates {
		middleIndex := len(update) / 2
		sumMiddleValues += update[middleIndex]
	}

	log.Println("Part 1 - Sum of middle values:", sumMiddleValues)
	return sumMiddleValues
}

func day5Part2(inputFilePath *string) int {
	log.Println("Day 5 Part 2")
	rules := LoadDay5Rules(inputFilePath)
	updates := LoadDay5Updates(inputFilePath)

	badUpdates := make([][]int, 0)
	for _, update := range updates {
		isValid := isValidUpdate(update, rules)
		if !isValid {
			badUpdates = append(badUpdates, update)
		}
	}

	newlyGoodUpdates := make([][]int, 0)
	for _, badUpdate := range badUpdates {
		invalidIndex, ok := isValidUpdateWithInvalidIndex(badUpdate, rules)

		// Brute force... feels like there should be a better way to do this.
		for !ok {
			// Take element at invalid index and move it to the beginning
			// of the slice.
			element := badUpdate[invalidIndex]
			copy(badUpdate[1:invalidIndex+1], badUpdate[0:invalidIndex])
			badUpdate[0] = element

			invalidIndex, ok = isValidUpdateWithInvalidIndex(badUpdate, rules)
		}

		newlyGoodUpdates = append(newlyGoodUpdates, badUpdate)
	}

	sumMiddleValues := 0

	for _, update := range newlyGoodUpdates {
		middleIndex := len(update) / 2
		sumMiddleValues += update[middleIndex]
	}

	log.Println("Part 2 - Sum of middle values:", sumMiddleValues)
	return sumMiddleValues
}
func Day5(inputFilePath *string) {
	log.Println("Day 5")
	day5Part1(inputFilePath)
	day5Part2(inputFilePath)
}
