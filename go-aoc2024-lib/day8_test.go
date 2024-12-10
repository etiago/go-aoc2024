package goaoc2024lib

import (
	"fmt"
	"testing"
)

func TestGetAntinodesPart1(t *testing.T) {
	newPoints := getAntinodes(Point{27, 22}, Point{48, 2})
	fmt.Println(newPoints)
	// log.Println(newPoints)
}
