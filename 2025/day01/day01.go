package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("2025/day01/day01.input")
	if err != nil {
		println(err)
		panic(err)
	}
	defer file.Close()
	// partOne(file)
	partTwo(file)
}

func partOne(file *os.File) {

	start := 10050
	count := 0
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		row := scanner.Text()
		sign := 1
		if row[:1] == "L" {
			sign = -1
		}
		value, err := strconv.Atoi(row[1:])
		if err != nil {
			println(err)
			panic(err)
		}
		start = start + (value * sign)
		if start%100 == 0 {
			count++
		}
	}

	fmt.Printf("partOne: %v\n", count)
}

func partTwo(file *os.File) {

	start := 100050
	count := 0
	atZero := false
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		row := scanner.Text()
		sign := 1
		if row[:1] == "L" {
			sign = -1
		}
		value, err := strconv.Atoi(row[1:])
		if err != nil {
			println(err)
			panic(err)
		}
		temp := 0
		if sign > 0 {
			temp = (start % 100) + value
		} else {
			temp = 100 - (start % 100) + value
		}
		// fmt.Printf("start: %v, sign: %v, temp: %v, value:%v, count: %v\n", start, sign, temp, value, count)
		start = start + (value * sign)
		if start%100 == 0 {
			// fmt.Printf("At zero\n")
			count++
			atZero = true
		} else {
			delta := temp / 100
			if atZero && delta > 0 && sign < 0 {
				fmt.Printf("start: %v, sign: %v, temp: %v, value:%v, count: %v, delta: %v\n", start, sign, temp, value, count, delta)
				delta += -1
			}
			count += delta
			atZero = false
		}
	}

	fmt.Printf("partTwo: %v\n", count)
}
