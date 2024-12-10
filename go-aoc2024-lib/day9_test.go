package goaoc2024lib

import "testing"

func TestDefragLastBlockToFirstFreeSpaces(t *testing.T) {
	path := "../input/day9-example.txt"
	disk := LoadDay9Disk(&path)

	defragBlockToFirstFreeSpaces(&disk)
}

func TestDefragWholeBlocks(t *testing.T) {
	path := "../input/day9-example.txt"
	disk := LoadDay9Disk(&path)

	defragWholeBlocks(&disk)

}
