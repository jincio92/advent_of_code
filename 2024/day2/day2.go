package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {

	file, err := os.Open("day2/day2.input")
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)

	safeReports := 0
	line := 0
	for scanner.Scan() {
		levels := strings.Split(scanner.Text(), " ")

		result := checkList(levels)
		if result {
			safeReports++
		} else {
			fmt.Printf("START*****")
			println()
			fmt.Printf("levels %v", levels)
			println()
			for i := 0; i < len(levels); i++ {
				newARR := slices.Delete(slices.Clone(levels), i, i+1)
				fmt.Printf("newARR %v", newARR)
				println()
				result = checkList(newARR)
				if result {
					safeReports++
					break
				}
			}
		}
		line++
	}

	fmt.Printf("Safe Reports: %d", safeReports)
	println()
}

func checkList(levels []string) bool {

	order := 0
	safe := true
	for i := 0; i < (len(levels) - 1); i++ {

		aint, err := strconv.Atoi(levels[i])
		if err != nil {
			panic(err)
		}
		bint, err := strconv.Atoi(levels[i+1])
		if err != nil {
			panic(err)
		}
		if order == 0 {
			if aint < bint {
				order = -1
			} else if aint > bint {
				order = 1
			} else {
				safe = false
				break
			}
		}
		if (aint < bint && order > 0) || (aint > bint && order < 0) || aint == bint {
			safe = false
			break
		}
		difference := bint - aint
		if difference < 0 {
			difference = -difference
		}
		if difference > 3 {
			safe = false
			break
		}
	}
	return safe
}
