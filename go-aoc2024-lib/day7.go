package goaoc2024lib

import (
	"log"
	"math"
	"strconv"

	"gonum.org/v1/gonum/stat/combin"
)

const (
	OperationAdd int = iota
	OperationMultiply
	OperationConcatenate
)

// I don't love this, but it works.
func generateCombinations(numbers []int, length int) [][]int {
	dims := make([]int, length)
	for i := range dims {
		dims[i] = len(numbers)
	}

	cart := combin.NewCartesianGenerator(dims)
	results := [][]int{}

	prod := make([]int, length)
	for cart.Next() {
		cart.Product(prod)

		combo := make([]int, length)
		for i, idx := range prod {
			if idx < len(numbers) {
				combo[i] = numbers[idx]
			}
		}
		results = append(results, combo)
	}

	return results
}

type ComboIsValidChecker func(expectedSum int64, operands []int64, combo []int) bool

func part1ComboIsValid(expectedSum int64, operands []int64, combo []int) bool {
	sum := operands[0]
	for i := 1; i < len(operands); i++ {
		if combo[i-1] == OperationAdd {
			sum += operands[i]
		} else if combo[i-1] == OperationMultiply {
			sum *= operands[i]
		}
	}
	return sum == expectedSum
}

func hasValidCombo(checker ComboIsValidChecker, expectedSum int64, operands []int64, combos [][]int) bool {
	for _, combo := range combos {
		if checker(expectedSum, operands, combo) {
			return true
		}
	}

	return false
}
func day7Part1(inputFilePath *string) int64 {
	equations := LoadDay7Equations(inputFilePath)
	sum := int64(0)
	for _, equation := range equations {
		gapCount := len(equation.operands) - 1

		possibleVals := []int{0, 1}
		combos := generateCombinations(possibleVals, gapCount)

		if hasValidCombo(part1ComboIsValid, equation.result, equation.operands, combos) {
			sum += equation.result
		}
	}

	log.Println("Sum:", sum)
	return sum
}

func part2ComboIsValid(expectedSum int64, operands []int64, combo []int) bool {
	sum := operands[0]
	for i := 1; i < len(operands); i++ {
		if combo[i-1] == OperationAdd {
			sum += operands[i]
		} else if combo[i-1] == OperationMultiply {
			sum *= operands[i]
		} else if combo[i-1] == OperationConcatenate {
			digitCount := len(strconv.FormatInt(operands[i], 10))
			sum = sum*int64(math.Pow10(digitCount)) + operands[i]
		}
	}
	return sum == expectedSum
}

func day7Part2(inputFilePath *string) int64 {
	equations := LoadDay7Equations(inputFilePath)
	sum := int64(0)
	for _, equation := range equations {
		gapCount := len(equation.operands) - 1

		possibleVals := []int{0, 1, 2}
		combos := generateCombinations(possibleVals, gapCount)

		if hasValidCombo(part2ComboIsValid, equation.result, equation.operands, combos) {
			sum += equation.result
		}
	}

	log.Println("Sum:", sum)
	return sum
}

func Day7(inputFilePath *string) {
	day7Part1(inputFilePath)
	day7Part2(inputFilePath)
}
