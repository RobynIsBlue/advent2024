package main

import (
	"bufio"
	"fmt"
	"os"
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
	file, err := os.Open(`.\test.txt`)
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
		count += ruleChecker(line, rulebook, false)
	}
	fmt.Println(count)
}

func ruleMaker(line string, rulebook map[int][]int) {
	nums := strings.Split(line, "|")
	num1, _ := strconv.Atoi(nums[0])
	num2, _ := strconv.Atoi(nums[1])
	rulebook[num1] = append(rulebook[num1], num2)
}

func ruleChecker(line string, rulebook map[int][]int, bad bool) int {
	nums := strings.Split(line, ",")
	checkedNums := []int{}
	for indFirstNum, strNum := range nums {
		num, _ := strconv.Atoi(strNum)
		if k, ok := rulebook[num]; ok {
			for _, ruleNum := range k {
				for indSliceNum, sliceNum := range checkedNums {
					if sliceNum == ruleNum {
						newList := swapInd(indFirstNum, indSliceNum, nums)
						return ruleChecker(newList, rulebook, true)
					}
				}
			}
		}
		checkedNums = append(checkedNums, num)
	}
	if bad {
		return checkedNums[(len(checkedNums) / 2)]
	}
	return 0
}

func swapInd(i, j int, list []string) string {
	val1 := list[i]
	val2 := list[j]
	list[i] = val2
	list[j] = val1
	return strings.Join(list, ",")
}

func main() {
	reader()
}
