package main

import (
	goaoc2024lib "etiago/go-aoc2024-lib"
	"flag"
	"log"
)

func main() {
	day := flag.Int("day", 1, "Day to run")
	inputFilePath := flag.String("input-file", "", "Input file")
	flag.Parse()

	switch *day {
	case 1:
		goaoc2024lib.Day1(inputFilePath)
	case 2:
		goaoc2024lib.Day2(inputFilePath)
	case 3:
		goaoc2024lib.Day3(inputFilePath)
	case 4:
		goaoc2024lib.Day4(inputFilePath)
	case 5:
		goaoc2024lib.Day5(inputFilePath)
	case 6:
		goaoc2024lib.Day6(inputFilePath)
	case 7:
		goaoc2024lib.Day7(inputFilePath)
	case 8:
		goaoc2024lib.Day8(inputFilePath)
	case 9:
		goaoc2024lib.Day9(inputFilePath)
	case 10:
		goaoc2024lib.Day10(inputFilePath)
	default:
		log.Fatalf("Day %d not implemented", *day)
	}
}
