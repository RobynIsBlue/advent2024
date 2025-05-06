package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func reader() [][]string {
	file, err := os.Open(`..\input1.txt`)
	if err != nil {
		fmt.Println("error opening input")
	}
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

func star(matrix [][]string) {
	var count int
	for i, v := range matrix {
		if i == 0 || i == len(matrix)-1 {
			continue
		}
		for j, _ := range v {
			if j == 0 || j == len(matrix[0])-1 {
				continue
			}
			if matrix[i][j] == "A" {
				count += smChecker(i, j, matrix)
			}
		}
	}
	fmt.Print(count)
}

func smChecker(i, j int, matrix [][]string) int {
	var bottomRight string
	switch matrix[i-1][j-1] {
	case "S":
		bottomRight = "M"
	case "M":
		bottomRight = "S"
	default:
		return 0
	}
	if matrix[i+1][j+1] != bottomRight {
		return 0
	}
	bottomLeft := matrix[i+1][j-1]
	topRight := matrix[i-1][j+1]

	if topRight != bottomLeft &&
		(topRight == "S" || topRight == "M") &&
		(bottomLeft == "S" || bottomLeft == "M") {
		return 1
	}
	return 0
}

func main() {
	matrix := reader()
	star(matrix)
}
