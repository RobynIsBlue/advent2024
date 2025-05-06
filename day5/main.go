package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

/*
first split by empty line to get rules and list

for each rule, split
take first value and make it a key in a map (if it doesn't already exist)
take second value and map it to the key

for each list, split
check each value as a key for other ma
*/

func reader() {
	file, err := os.Open(`.\input.txt`)
	if err != nil {
		fmt.Println("couldn't open file")
		return
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	rulebook := make(map[int][]int)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			break
		}
		ruleMaker(line, rulebook)
	}

	count := 0
	for scanner.Scan() {
		line := scanner.Text()
		count += ruleChecker(line, rulebook)
	}
	fmt.Println(count)
}

func ruleMaker(line string, rulebook map[int][]int) {
	nums := strings.Split(line, "|")
	num1, _ := strconv.Atoi(nums[0])
	num2, _ := strconv.Atoi(nums[1])
	rulebook[num1] = append(rulebook[num1], num2)
}

func ruleChecker(line string, rulebook map[int][]int) int {
	nums := strings.Split(line, ",")
	checkedNums := []int{}
	for _, p := range nums {
		n, _ := strconv.Atoi(p)
		if k, ok := rulebook[n]; ok {
			for _, j := range k {
				if slices.Contains(checkedNums, j) {
					return 0
				}
			}
		}
		checkedNums = append(checkedNums, n)
	}
	return checkedNums[(len(checkedNums) / 2)]
}

func main() {
	reader()
}
