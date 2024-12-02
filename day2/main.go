package main

import (
	"adventofcode/shared"
	"fmt"
	"log"
	"strings"
)

func main() {
	lines, err := shared.ReadLines("input.txt")
	if err != nil {
		log.Fatalf("Failed to read input %v", err)
	}
	inputAsInts := make([][]int, len(lines))
	for i := range inputAsInts {
		spaceSeparatedLine := strings.Split(lines[i], " ")
		inputAsInts[i], _ = shared.ConvertStringToInts(spaceSeparatedLine)
	}

	Part1(inputAsInts)
	Part2(inputAsInts)
}

func Part1(input [][]int) {
	numberOfSafe := 0
	for _, line := range input {
		confirmedUnsafe := false
		isAscending := line[0] < line[1]
		for j := 1; j < len(line); j++ {
			diff := line[j] - line[j-1]
			if isAscending && diff > 0 && diff < 4 {
				continue
			} else if !isAscending && diff < 0 && diff > -4 {
				continue
			} else {
				confirmedUnsafe = true
				break
			}
		}
		if !confirmedUnsafe {
			numberOfSafe++
		}
	}
	fmt.Println("Result Part 1: ", numberOfSafe)
}

func Part2(input [][]int) {
	numberOfSafe := 0
	for _, line := range input {
		confirmedUnsafe := CheckLine(line, false)
		if confirmedUnsafe {
			confirmedUnsafe = CheckLine(line[1:], true)
		}

		if !confirmedUnsafe {
			numberOfSafe++
		}
	}
	fmt.Println("Result Part 2: ", numberOfSafe)
}

func CheckLine(line []int, withoutFirstNumber bool) bool {
	hasRemovedElement := withoutFirstNumber
	isAscending := line[0] < line[1]
	confirmedUnsafe := false
	for j := 1; j < len(line); j++ {
		confirmedUnsafe = !CheckIfSafe(line[j-1], line[j], isAscending)
		if confirmedUnsafe && !hasRemovedElement {
			couldBeRemoved := true
			if j != len(line)-1 {
				couldBeRemoved = CheckIfSafe(line[j-1], line[j+1], isAscending)
			} else {
				couldBeRemoved = true
			}
			if couldBeRemoved {
				confirmedUnsafe = false
				hasRemovedElement = true
				j++
			}
		}
		// If it is still unsafe, we should just exit the loop as line is no longer valid
		if confirmedUnsafe {
			break
		}
	}
	return confirmedUnsafe
}

func CheckIfSafe(input1 int, input2 int, isAscending bool) bool {
	diff := input2 - input1
	if isAscending && diff > 0 && diff < 4 {
		return true
	} else if !isAscending && diff < 0 && diff > -4 {
		return true
	} else {
		return false
	}
}
