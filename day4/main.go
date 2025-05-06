package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	matrix := reader()
	fmt.Println(waves(matrix))
}

func errChecker(text string, err error) {
	if err != nil {
		fmt.Println(text)
		panic(err)
	}
}

// reads the input file
func reader() [][]string {
	file, err := os.Open(`.\input1.txt`)
	errChecker("error opening file", err)
	defer file.Close()

	var matrix [][]string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		splitLine := strings.Split(line, "")
		matrix = append(matrix, splitLine)
	}
	return matrix
}

// checks every value
func waves(matrix [][]string) int {

	value := 0
	// for every line in the matrix
	for i, line := range matrix {

		// for every value in line
		for j, val := range line {

			// if the value is X, give directionChecker the
			// matrix and pos. of X
			if val == "X" {
				value += directionChecker(matrix, i, j)
			}
		}
	}
	return value
}

// checks viable directions
func directionChecker(matrix [][]string, index1, index2 int) int {
	count := 0

	// check which directions are viable to check
	var left, right, up, down bool
	if index2 >= 3 {
		left = true
	}
	if index2+3 < len(matrix[0]) {
		right = true
	}
	if index1 >= 3 {
		up = true
	}
	if index1+3 < len(matrix) {
		down = true
	}

	// check every index above, equal to, and below
	for i := -1; i <= 1; i++ {
		if !up && i == -1 {
			continue
		}
		if !down && i == 1 {
			continue
		}

		//check every value to the left of, in the same column of, and to the right of
		for j := -1; j <= 1; j++ {
			if !left && j == -1 {
				continue
			}
			if !right && j == 1 {
				continue
			}
			if i == 0 && j == 0 {
				continue
			}
			if masChecker(matrix, index1, index2, i, j, 0) {
				count += 1
			}
		}
	}
	return count
}

func masChecker(matrix [][]string, index1, index2, plusI, plusJ, calls int) bool {
	mas := []string{"M", "A", "S"}
	letter := mas[calls]

	if matrix[index1+plusI][index2+plusJ] != letter {
		return false
	}

	if calls == 2 {
		return true
	}

	switch {
	case plusI < 0:
		plusI -= 1
	case plusI > 0:
		plusI += 1
	}

	switch {
	case plusJ < 0:
		plusJ -= 1
	case plusJ > 0:
		plusJ += 1
	}

	return masChecker(matrix, index1, index2, plusI, plusJ, calls+1)

}
