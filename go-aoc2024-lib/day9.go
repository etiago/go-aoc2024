package goaoc2024lib

import (
	"log"
	"math"
	"sort"
)

type ByFirstIndex []Block

func (a ByFirstIndex) Len() int           { return len(a) }
func (a ByFirstIndex) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByFirstIndex) Less(i, j int) bool { return a[i].indices[0] < a[j].indices[0] }

// I'm not proud of this and I'm sure there's a better way to do it.
func defragBlockToFirstFreeSpaces(disk *Disk) {
	// Get last block
	blockToDefrag := disk.dataBlocks[len(disk.dataBlocks)-1]

	positionsToMove := int(math.Min(float64(len(blockToDefrag.indices)), float64(len(disk.freeBlock.indices))))

	if len(disk.freeBlock.indices) == 0 {
		return
	}

	startIndex := len(blockToDefrag.indices) - 1
	blockLen := len(blockToDefrag.indices) - 1
	for positionsToMove > 0 {
		// Take the index of first free block
		freeIndex := disk.freeBlock.indices[0]
		// Get the last index of lastBlock
		newlyFreedIndex := blockToDefrag.indices[startIndex]

		if freeIndex > newlyFreedIndex {
			startIndex--
			positionsToMove--
			continue
		}

		// Remove the first free block from the free block list
		disk.freeBlock.indices = disk.freeBlock.indices[1:]

		// Remove last index of lastBlock
		blockToDefrag.indices = blockToDefrag.indices[:blockLen]

		// Prepend freeIndex to indices of lastBlock
		blockToDefrag.indices = append([]int{freeIndex}, blockToDefrag.indices...)

		// Append newlyFreedIndex to free block list
		disk.freeBlock.indices = append(disk.freeBlock.indices, newlyFreedIndex)

		// Decrement positionsToMove
		positionsToMove--
	}

	// Rather than sorting which is costly, we could have also just added
	// the indices at the correct position in O(n) time. On the other hand,
	// we would have been resizing the slice multiple times which is also
	// costly. So, I chose to sort the slices.
	sort.Ints(blockToDefrag.indices)
	sort.Ints(disk.freeBlock.indices)

	disk.dataBlocks[len(disk.dataBlocks)-1] = blockToDefrag

	// Move the last block to the beginning of the disk.dataBlocks
	disk.dataBlocks = append([]Block{disk.dataBlocks[len(disk.dataBlocks)-1]}, disk.dataBlocks[:len(disk.dataBlocks)-1]...)

	if blockToDefrag.id == 0 {
		sort.Sort(ByFirstIndex(disk.dataBlocks))
		return
	}
	defragBlockToFirstFreeSpaces(disk)
}

func day9Part1(inputFilePath *string) {
	disk := LoadDay9Disk(inputFilePath)

	defragBlockToFirstFreeSpaces(&disk)

	sum := 0
	for _, block := range disk.dataBlocks {
		for _, index := range block.indices {
			sum += block.id * index
		}
	}
	log.Println("Sum:", sum)
}

func freeBlockHasNumOfContiguousIndices(disk *Disk, num int) (int, bool) {
	for i := 0; i < len(disk.freeBlock.indices)-num-1; i++ {
		firstFreeIndex := disk.freeBlock.indices[i]
		if disk.freeBlock.indices[i+num-1] == firstFreeIndex+num-1 {
			return i, true
		}
	}
	return -1, false
}

func defragWholeBlocks(disk *Disk) {
	// Get last block
	blockToDefrag := disk.dataBlocks[len(disk.dataBlocks)-1]

	blockCountToMove := len(blockToDefrag.indices)

	firstFree, ok := freeBlockHasNumOfContiguousIndices(disk, blockCountToMove)
	if !ok || disk.freeBlock.indices[firstFree] > blockToDefrag.indices[0] {
		if blockToDefrag.id == 0 {
			sort.Sort(ByFirstIndex(disk.dataBlocks))
			return
		}
		// Move last block to the beginning of the disk
		disk.dataBlocks = append([]Block{disk.dataBlocks[len(disk.dataBlocks)-1]}, disk.dataBlocks[:len(disk.dataBlocks)-1]...)
		defragWholeBlocks(disk)
		return
	}

	// Take blockCountToMove indices from the free block, starting at firstFree
	freeIndices := disk.freeBlock.indices[firstFree : firstFree+blockCountToMove]

	// Take old indices of the last block
	oldIndices := blockToDefrag.indices

	// Set the indices of the last block to the free indices
	blockToDefrag.indices = make([]int, len(freeIndices))
	copy(blockToDefrag.indices, freeIndices)

	// Append the old indices to the free block
	disk.freeBlock.indices = append(disk.freeBlock.indices, oldIndices...)

	// Remove the firstFree to firstFree+blockCountToMove indices from the free block
	disk.freeBlock.indices = append(disk.freeBlock.indices[:firstFree], disk.freeBlock.indices[firstFree+blockCountToMove:]...)

	sort.Ints(blockToDefrag.indices)
	sort.Ints(disk.freeBlock.indices)

	disk.dataBlocks[len(disk.dataBlocks)-1] = blockToDefrag

	// Move the last block to the beginning of the disk.dataBlocks
	disk.dataBlocks = append([]Block{disk.dataBlocks[len(disk.dataBlocks)-1]}, disk.dataBlocks[:len(disk.dataBlocks)-1]...)

	if blockToDefrag.id == 0 {
		sort.Sort(ByFirstIndex(disk.dataBlocks))
		return
	}

	defragWholeBlocks(disk)
}

func day9Part2(inputFilePath *string) {
	disk := LoadDay9Disk(inputFilePath)

	defragWholeBlocks(&disk)

	sum := 0
	for _, block := range disk.dataBlocks {
		for _, index := range block.indices {
			sum += block.id * index
		}
	}
	log.Println("Sum:", sum)
}

func Day9(inputFilePath *string) {
	day9Part1(inputFilePath)
	day9Part2(inputFilePath)
}
