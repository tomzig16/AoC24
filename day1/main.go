package main

import (
	"adventofcode/shared"
	"fmt"
	"log"
	"slices"
	"strings"
)

func main() {
	lines, err := shared.ReadLines("input.txt")
	if err != nil {
		log.Fatalf("Failed to read input %v", err)
	}

	leftArr, rightArr := GetTwoArrays(lines)

	Part1(leftArr, rightArr)
	Part2(leftArr, rightArr) // careful, they are sorted
	// Part 2 probably could be solved with simple find function over the array, but I am not sure how to
	// count how many times it finds. Easier to store in the map.
	// Also it could be even better if we could store strings that we read before converting into the map
	// potentially saving some time there
}

func GetTwoArrays(data []string) ([]int, []int) {

	var leftArray []int
	var rightArray []int

	for _, line := range data {
		splitLine := strings.Split(line, "   ")
		lineOfInts, _ := shared.ConvertStringToInts(splitLine)

		leftArray = append(leftArray, lineOfInts[0])
		rightArray = append(rightArray, lineOfInts[1])
	}

	return leftArray, rightArray
}

func Part1(leftArr []int, rightArr []int) {
	slices.Sort(leftArr)
	slices.Sort(rightArr)

	sum := 0
	for i := 0; i < len(leftArr); i++ {
		distance := leftArr[i] - rightArr[i]
		sum += max(distance, -distance) // golang apparently does not have good abs function?
	}

	fmt.Println("Result Part 1: ", sum)
}

func Part2(leftArr []int, rightArr []int) {
	rightMap := make(map[int]int)

	for _, el := range rightArr {
		quantity, ok := rightMap[el]
		if ok {
			rightMap[el] = quantity + 1
		} else {
			rightMap[el] = 1
		}
	}

	sum := 0
	for _, el := range leftArr {
		quantity, ok := rightMap[el]
		if ok {
			sum += (el * quantity)
		}
	}

	fmt.Println("Result Part 2: ", sum)
}
