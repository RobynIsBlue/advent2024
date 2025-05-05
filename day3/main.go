package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func main() {
	readInput2()
}

// func readInput1() {
// 	output, err := os.Open("C:\\codingsh\\github.com\\RobynIsBlue\\advent2024\\day1\\input1.txt")
// 	if err != nil {
// 		fmt.Println("failed reading file")
// 	}
// 	defer output.Close()
// 	scanner := bufio.NewScanner(output)

// 	val := 0
// 	for scanner.Scan() {
// 		line := scanner.Text()
// 		val += parser(doDont(line))
// 		fmt.Println(val)
// 	}
// 	fmt.Println(val)

// }

func readInput2() {
	dat, err := os.ReadFile("C:\\codingsh\\github.com\\RobynIsBlue\\advent2024\\day1\\input1.txt")
	if err != nil {
		fmt.Println("failed reading file")
	}
	data := string(dat)
	fmt.Println(parser(doDont(data)))
	fmt.Println(parser(data))
	fmt.Println(parser(data) - parser(doDont(data)))
}

func parser(line string) int {
	match, _ := regexp.Compile(`mul\((\d\w{0,3}),(\d\w{0,3})\)`)
	subs := match.FindAllStringSubmatch(line, -1)
	val := 0
	for _, n := range subs {
		x, err := strconv.Atoi(n[1])
		y, err := strconv.Atoi(n[2])
		if err != nil {
			fmt.Println("error converting numbers")
		}
		val += x * y
	}
	return val
}

func doDont(line string) string {
	match := regexp.MustCompile(`don't\(\)(.*?)((do\(\))|\n)`)
	hey := match.ReplaceAllLiteralString(line, "")
	os.WriteFile("output.txt", []byte(hey), 0644)
	return hey
}
