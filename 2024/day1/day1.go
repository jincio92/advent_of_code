package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func main() {

	file, err := os.Open("day1/day1.input")
	if err != nil {
		println(err)
		panic(err)
	}
	defer file.Close()

	partOne(file)
	partTwo(file)
}

func partTwo(file *os.File) {
	scanner := bufio.NewScanner(file)
	firstList := make([]int, 0)
	secondList := make([]int, 0)

	for scanner.Scan() {

		first, err := strconv.Atoi(scanner.Text()[:5])
		if err != nil {
			panic(err)
		}
		second, err := strconv.Atoi(scanner.Text()[8:])
		if err != nil {
			panic(err)
		}
		firstList = append(firstList, first)
		secondList = append(secondList, second)
	}
	similarity := 0
	for _, v := range firstList {
		multiplier := 0
		for _, s := range secondList {
			if s == v {
				multiplier++
			}
		}
		similarity += (v * multiplier)
	}

	fmt.Printf("similarity: %d", similarity)
	println()
}

func partOne(file *os.File) {

	scanner := bufio.NewScanner(file)
	firstList := make([]string, 0)
	secondList := make([]string, 0)

	for scanner.Scan() {
		// fmt.Printf("primo %s, secondo %s", scanner.Text()[:5], scanner.Text()[8:])

		firstList = append(firstList, (scanner.Text()[:5]))
		secondList = append(secondList, (scanner.Text()[8:]))
	}
	sort.Strings(firstList)
	sort.Strings(secondList)
	totalDistance := 0

	for i := 0; i < len(firstList); i++ {
		firstIndex, err := strconv.Atoi(firstList[i])
		if err != nil {
			panic(err)
		}
		secondIndex, err := strconv.Atoi(secondList[i])
		if err != nil {
			panic(err)
		}
		distance := secondIndex - firstIndex
		if distance < 0 {
			distance = -distance
		}
		totalDistance += distance
	}
	fmt.Printf("totalDistance: %d", totalDistance)
	println()
}
