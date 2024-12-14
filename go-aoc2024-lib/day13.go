package goaoc2024lib

import "log"

func getCheapestCombination(game Day13Game) (int, int, bool) {
	upOnA := false
	upOnB := true
	// downOnA := false
	downOnB := false

	pressesOnA := 0
	pressesOnB := 0
	sumX := 0
	sumY := 0

	for game.prize.x != sumX || game.prize.y != sumY {
		if upOnB {
			pressesOnB++
			sumX += game.buttonBStep.x
			sumY += game.buttonBStep.y

			if sumX > game.prize.x && sumY > game.prize.y {
				upOnB = false
				downOnB = true
			}
			continue
		}

		if downOnB {
			pressesOnB--
			sumX -= game.buttonBStep.x
			sumY -= game.buttonBStep.y

			if pressesOnB < 0 {
				return 0, 0, false
			}
			if sumX < game.prize.x && sumY < game.prize.y {
				upOnA = true
				downOnB = false
			}
			continue
		}

		if upOnA {
			if sumX > game.prize.x+game.buttonAStep.x || sumY > game.prize.y+game.buttonAStep.y {
				upOnA = false
				downOnB = true
				continue
			}

			pressesOnA++
			sumX += game.buttonAStep.x
			sumY += game.buttonAStep.y

			if sumX > game.prize.x || sumY > game.prize.y {
				upOnA = false
				downOnB = true
			}
		}
	}

	return pressesOnA, pressesOnB, true
}

func day13Part1(inputFilePath *string) {
	games := LoadDay13Games(inputFilePath, false)
	cost := 0

	for _, game := range games {
		pressesOnA, pressesOnB, found := getCheapestCombination(game)
		if found {
			cost += pressesOnA*3 + pressesOnB
		}
	}
	log.Println("Cost:", cost)
}

func getPresses(game Day13Game) (int, int, bool) {
	denominator := game.buttonBStep.x*game.buttonAStep.y - (game.buttonBStep.y * game.buttonAStep.x)
	aPresses := (game.buttonBStep.x*game.prize.y - game.buttonBStep.y*game.prize.x) / denominator
	bPresses := (game.buttonAStep.y*game.prize.x - game.buttonAStep.x*game.prize.y) / denominator

	if aPresses*game.buttonAStep.x+bPresses*game.buttonBStep.x != game.prize.x {
		return 0, 0, false
	}
	if aPresses*game.buttonAStep.y+bPresses*game.buttonBStep.y != game.prize.y {
		return 0, 0, false
	}
	return aPresses, bPresses, true
}
func day13Part2(inputFilePath *string) {
	games := LoadDay13Games(inputFilePath, true)
	cost := 0

	for _, game := range games {
		pressesOnA, pressesOnB, found := getPresses(game)
		if found {
			cost += pressesOnA*3 + pressesOnB
		}
	}
	log.Println("Cost:", cost)
}

func Day13(inputFilePath *string) {
	day13Part1(inputFilePath)
	day13Part2(inputFilePath)
}
