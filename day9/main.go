package main

import (
	"adventofcode/shared"
	"log"
	"strconv"
)

type MemoryLayout struct {
	start  int
	length int
}

func main() {
	lines, err := shared.ReadLines("input.txt")
	if err != nil {
		log.Fatalf("Failed to read input %v", err)
	}

	occupiedMemory := make([]MemoryLayout, 0)
	freeMemory := make([]MemoryLayout, 0)
	tempStart := 0
	for i := 0; i < len(lines[0]); i += 2 {
		occupiedSlots, _ := strconv.Atoi(string(lines[0][i]))

		tempLength := occupiedSlots
		occupiedMemory = append(occupiedMemory, MemoryLayout{start: tempStart, length: tempLength})

		if i+1 >= len(lines[0]) {
			break
		}
		freeSlots, _ := strconv.Atoi(string(lines[0][i+1]))
		tempStart += tempLength
		tempLength = freeSlots
		freeMemory = append(freeMemory, MemoryLayout{start: tempStart, length: tempLength})

		tempStart += tempLength
	}

	Part1(occupiedMemory, freeMemory)
	Part2(lines)

}

func Part1(occupiedMemory []MemoryLayout, freeMemory []MemoryLayout) {
	checksum := 0
	checksumCounter := 0

	occupiedIndex := 0

	occupiedIndexBackwards := len(occupiedMemory) - 1
	occupiedFromBackwardsLeftover := occupiedMemory[occupiedIndexBackwards].length

	freeIndex := 0
	freeSpacesInNextSegment := 0
	// allMemoryMoved := false

	for occupiedIndex < occupiedIndexBackwards {
		occupiedSegmentLength := occupiedMemory[occupiedIndex].length
		for i := 0; i < occupiedSegmentLength; i++ {
			checksum += (checksumCounter * occupiedIndex)
			checksumCounter++
		}
		occupiedIndex++

		// Corner case if we run out of free spaces first
		freeSpacesInNextSegment = freeMemory[freeIndex].length

		for i := freeSpacesInNextSegment; i > 0; i-- {
			if occupiedFromBackwardsLeftover == 0 {
				occupiedIndexBackwards--
				occupiedFromBackwardsLeftover = occupiedMemory[occupiedIndexBackwards].length
			}
			checksum += (checksumCounter * occupiedIndexBackwards)
			checksumCounter++
			occupiedFromBackwardsLeftover--
		}
		freeIndex++
	}
	// If we have left any unaccounted numbers
	// usedFromLastSegment := occupiedFromBackwardsLeftover
	for i := 0; i < occupiedFromBackwardsLeftover; i++ {
		checksum += (checksumCounter * occupiedIndexBackwards)
		checksumCounter++
	}

	// for i := 0; i < occupiedMemory[occupiedIndex].length-usedFromLastSegment; i++ {
	// 	checksum += (checksumCounter * occupiedIndex)
	// 	checksumCounter++
	// }
	shared.PrintResultsInt(1, checksum)
}

func Part2(lines []string) {

}
