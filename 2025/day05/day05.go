package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	file, err := os.Open("2025/day05/day05.input")
	if err != nil {
		println(err)
		panic(err)
	}
	defer file.Close()

	ranges := [][]int{}
	ids := []int{}
	scanner := bufio.NewScanner(file)
	isIds := false
	for scanner.Scan() {
		row := scanner.Text()
		if len(row) == 0 {
			isIds = true
			continue
		}

		if !isIds {
			rs := strings.Split(row, "-")
			first, err := strconv.Atoi(rs[0])
			if err != nil {
				panic(err)
			}
			second, err := strconv.Atoi(rs[1])
			if err != nil {
				panic(err)
			}
			ranges = append(ranges, []int{first, second})
		} else {
			id, err := strconv.Atoi(row)
			if err != nil {
				panic(err)
			}
			ids = append(ids, id)
		}
	}

	newRanges := checkOverlap(ranges)
	partOne(newRanges, ids)
	partTwo(newRanges)
}

func partTwo(ranges [][]int) {
	sum := 0
	for _, r := range ranges {
		sum += r[1] - r[0] + 1
	}
	fmt.Printf("partTwo: %v\n", sum)
}

func partOne(ranges [][]int, ids []int) {
	count := 0
	// idLoop:
	for _, id := range ids {
		for _, r := range ranges {
			if id >= r[0] && id <= r[1] {
				count++
				// continue idLoop
			}
		}
	}
	fmt.Printf("partOne: %v\n", count)
}

func checkOverlap(ranges [][]int) [][]int {
	newRanges := [][]int{}
	for i := 0; i < len(ranges); i++ {
		current := ranges[i]
		merged := false
		for j := i + 1; j < len(ranges); j++ {
			temp := ranges[j]
			if current[0] >= temp[0] && current[0] <= temp[1] {
				// fmt.Printf("start is inside: %v, %v\n", current, temp)
				var newRange []int
				if current[1] < temp[1] {
					newRange = []int{temp[0], temp[1]}
				} else {
					newRange = []int{temp[0], current[1]}
				}
				// fmt.Printf("newRange is: %v\n", newRange)
				ranges[j] = newRange
				merged = true
				continue
			} else if current[1] >= temp[0] && current[1] <= temp[1] {
				// fmt.Printf("end is inside: %v, %v\n", current, temp)
				var newRange []int
				if current[0] < temp[0] {
					newRange = []int{current[0], temp[1]}
				} else {
					newRange = []int{temp[0], temp[1]}
				}
				// fmt.Printf("newRange is: %v\n", newRange)
				ranges[j] = newRange
				merged = true
				continue
			} else if current[0] <= temp[0] && current[1] >= temp[1] {
				ranges[j] = current
				merged = true
				continue
			}
		}
		if !merged {
			newRanges = append(newRanges, current)
		}
	}
	return newRanges
}
