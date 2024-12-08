package main

import (
	goaoc2024lib "etiago/go-aoc2024-lib"
	"flag"
	"log"
)

func main() {
	day := flag.Int("day", 1, "Day to run")
	input_file_path := flag.String("input-file", "", "Input file")
	flag.Parse()

	switch *day {
	case 1:
		goaoc2024lib.Day1(input_file_path)
	case 2:
		goaoc2024lib.Day2(input_file_path)
	case 3:
		goaoc2024lib.Day3(input_file_path)
	case 4:
		goaoc2024lib.Day4(input_file_path)
	case 5:
		goaoc2024lib.Day5(input_file_path)
	case 6:
		goaoc2024lib.Day6(input_file_path)
	case 7:
		goaoc2024lib.Day7(input_file_path)
	case 8:
		goaoc2024lib.Day8(input_file_path)
	default:
		log.Fatalf("Day %d not implemented", *day)
	}
}
