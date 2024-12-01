package goaoc2024lib

import (
	"log"
	"os"
	"strings"
)

func ReadFile(input_file_path string) string {
	content, err := os.ReadFile(input_file_path)
	if err != nil {
		log.Fatal(err)
	}

	return string(content)
}

func ReadFileLines(input_file_path string) []string {
	content := ReadFile(input_file_path)
	return strings.Split(content, "\n")
}
