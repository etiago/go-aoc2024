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
	default:
		log.Fatalf("Day %d not implemented", *day)
	}
}
