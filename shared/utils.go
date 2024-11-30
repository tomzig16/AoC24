package shared

import (
	"bufio"
	"os"
	"strconv"
)

func ReadLines(fileName string) ([]string, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return lines, nil
}

func ConvertStringToInts(strings []string) ([]int, error) {
	var ints []int
	for _, s := range strings {
		num, err := strconv.Atoi(s)
		if err != nil {
			return nil, err
		}
		ints = append(ints, num)
	}
	return ints, nil
}

func MaxIntAndIndex(ints []int) (int, int) {
	maxIndex := 0
	max := ints[maxIndex]
	for i, num := range ints {
		if num > max {
			max = num
			maxIndex = i
		}
	}
	return max, maxIndex

}
