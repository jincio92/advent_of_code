package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("2025/day03/day03.input")
	if err != nil {
		println(err)
		panic(err)
	}
	defer file.Close()
	partOne(file)
	// partTwo(file)
}

func partOne(file *os.File) {
	scanner := bufio.NewScanner(file)
	sum1 := 0
	sum2 := 0
	for scanner.Scan() {
		row := scanner.Text()
		sum1 += findMax(row, 1)
		sum2 += findMax(row, 11)
	}
	fmt.Printf("PartOne: %v\n", sum1)
	fmt.Printf("PartTwo: %v\n", sum2)
}

func findMax(row string, limit int) int {
	indexMax := 0
	max, err := strconv.Atoi(string(row[indexMax]))
	if err != nil {
		panic(err)
	}
	for i := 1; i < len(row)-limit; i++ {
		a, err := strconv.Atoi(string(row[i]))
		if err != nil {
			println(err)
			panic(err)
		}
		if a > max {
			indexMax = i
			max = a
		}
	}
	if limit > 0 {
		sum := max * int(math.Pow10(limit))
		innermax := findMax(row[indexMax+1:], limit-1)
		// fmt.Printf("max: %v, sum: %v, inn: %v, limit: %v\n", max, sum, innermax, limit)
		return sum + innermax
	} else {
		// fmt.Printf("max: %v\n", max)
		return max
	}
}
