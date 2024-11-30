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

	fmt.Println("Input lines: ", lines)
}
