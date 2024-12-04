package main

import (
	"adventofcode/shared"
	"fmt"
	"log"
)

func main() {
	lines, err := shared.ReadLines("input.txt")
	if err != nil {
		log.Fatalf("Failed to read input %v", err)
	}

	lines = AddPaddingToArray(lines)

	Part1(lines)
	Part2(lines)
	// Part2(singleLine)

}

func AddPaddingToArray(input []string) []string {
	result := make([]string, len(input[0])+2)
	for i := 0; i < len(input[0])+2; i++ {
		// Add padding to first and last lines
		result[0] += "."
		result[len(input)+1] += "."
	}

	for i := 1; i < len(input)+1; i++ {
		result[i] = "." + input[i-1] + "."
	}
	return result
}

func Part1(input []string) {
	occurances := 0
	for i := 0; i < len(input); i++ {
		for j := 0; j < len(input[i]); j++ {
			if input[i][j] == 'X' {
				if input[i][j+1] == 'M' && input[i][j+2] == 'A' && input[i][j+3] == 'S' {
					occurances++
				}
				if input[i][j-1] == 'M' && input[i][j-2] == 'A' && input[i][j-3] == 'S' {
					occurances++
				}
				if input[i+1][j] == 'M' && input[i+2][j] == 'A' && input[i+3][j] == 'S' {
					occurances++
				}
				if input[i-1][j] == 'M' && input[i-2][j] == 'A' && input[i-3][j] == 'S' {
					occurances++
					/// Diagonally
				}
				if input[i+1][j+1] == 'M' && input[i+2][j+2] == 'A' && input[i+3][j+3] == 'S' {
					occurances++
				}
				if input[i+1][j-1] == 'M' && input[i+2][j-2] == 'A' && input[i+3][j-3] == 'S' {
					occurances++
				}
				if input[i-1][j-1] == 'M' && input[i-2][j-2] == 'A' && input[i-3][j-3] == 'S' {
					occurances++
				}
				if input[i-1][j+1] == 'M' && input[i-2][j+2] == 'A' && input[i-3][j+3] == 'S' {
					occurances++
				}
			}
		}
	}
	fmt.Println("Result Part 1: ", occurances)
}

func Part2(input []string) {
	occurances := 0
	for i := 0; i < len(input); i++ {
		for j := 0; j < len(input[i]); j++ {
			if input[i][j] == 'A' {
				crossingAmount := 0
				if input[i+1][j+1] == 'M' && input[i-1][j-1] == 'S' {
					crossingAmount++
				}
				if input[i+1][j+1] == 'S' && input[i-1][j-1] == 'M' {
					crossingAmount++
				}
				if input[i+1][j-1] == 'M' && input[i-1][j+1] == 'S' {
					crossingAmount++
				}
				if input[i+1][j-1] == 'S' && input[i-1][j+1] == 'M' {
					crossingAmount++
				}
				if crossingAmount > 1 {
					occurances++
				}
			}
		}
	}
	fmt.Println("Result Part 2: ", occurances)
}
