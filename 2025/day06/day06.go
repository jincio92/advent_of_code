package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	file, err := os.Open("2025/day06/day06.input")
	if err != nil {
		println(err)
		panic(err)
	}
	defer file.Close()
	lines := []string{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		row := scanner.Text()
		lines = append(lines, row)

	}
	partOne(lines)
	partTwo(lines)
}

func partOne(lines []string) {
	splittedLines := [][]string{}
	for _, v := range lines {
		splittedLines = append(splittedLines, strings.Fields(v))
	}
	sum := 0
	for col := 0; col < len(splittedLines[0]); col++ {
		var colTotal int
		operation := splittedLines[len(splittedLines)-1][col]
		if operation == "+" {
			colTotal = 0
		} else {
			colTotal = 1
		}
		for row := 0; row < len(splittedLines)-1; row++ {
			val, err := strconv.Atoi(splittedLines[row][col])
			if err != nil {
				fmt.Printf("row: %v", splittedLines[row])
				panic(err)
			}
			if operation == "+" {
				colTotal += val
			} else {
				colTotal *= val
			}
		}
		sum += colTotal
	}
	fmt.Printf("partOne: %v \n", sum)
}

func partTwo(lines []string) {

	operations := strings.Fields(lines[len(lines)-1])
	index := len(operations)
	sum := 0
	var colTotal int
	operation := operations[index-1]
	if operation == "+" {
		colTotal = 0
	} else {
		colTotal = 1
	}
	for col := len(lines[0]) - 1; col >= 0; col-- {
		number := ""
		for row := 0; row < len(lines)-1; row++ {
			number += lines[row][col : col+1]
		}
		if strings.TrimSpace(number) == "" {
			// fmt.Printf("colTotal: %v, index: %v\n", colTotal, index)
			sum += colTotal
			index--
			operation = operations[index-1]
			if operation == "+" {
				colTotal = 0
			} else {
				colTotal = 1
			}
			continue
		}
		val, err := strconv.Atoi(strings.TrimSpace(number))
		if err != nil {
			panic(err)
		}
		if operation == "+" {
			colTotal += val
		} else {
			colTotal *= val
		}

		// fmt.Printf("number: \"%v\"\n", strings.TrimSpace(number))
	}
	// fmt.Printf("colTotal: %v, index: %v\n", colTotal, index)
	sum += colTotal
	fmt.Printf("partTwo: %v \n", sum)
}
