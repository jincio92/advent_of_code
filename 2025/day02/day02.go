package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("2025/day02/day02.input")
	if err != nil {
		println(err)
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {

		row := scanner.Text()
		partOne(row)
		partTwo(row)
	}
}

func partOne(row string) {

	sum := 0
	pair := strings.Split(row, ",")

	// couples := [][]int{}

	for _, v := range pair {
		cop := strings.Split(v, "-")
		first, err := strconv.Atoi(cop[0])
		if err != nil {
			println(err)
			panic(err)
		}
		second, err := strconv.Atoi(cop[1])
		if err != nil {
			println(err)
			panic(err)
		}
		// couples = append(couples, []int{first, second})
		for i := first; i <= second; i++ {
			current := strconv.Itoa(i)
			if current[:len(current)/2] == current[len(current)/2:] {
				sum += i
			}

		}
	}

	fmt.Printf("partOne: %v\n", sum)
}

func partTwo(row string) {

	sum := 0
	breakpoint := map[int][]int{}
	pair := strings.Split(row, ",")

	// couples := [][]int{}

	for _, v := range pair {
		cop := strings.Split(v, "-")
		first, err := strconv.Atoi(cop[0])
		if err != nil {
			println(err)
			panic(err)
		}
		second, err := strconv.Atoi(cop[1])
		if err != nil {
			println(err)
			panic(err)
		}
		// couples = append(couples, []int{first, second})
		for i := first; i <= second; i++ {
			value := isInvalid(i, breakpoint)
			sum += value
		}
	}

	fmt.Printf("partTwo: %v\n", sum)
}

func isInvalid(code int, breakpoint map[int][]int) int {
	current := strconv.Itoa(code)
	if breakpoint[len(current)] == nil {
		populateBrakpoint(breakpoint, len(current))
	}
out:
	for _, v := range breakpoint[len(current)] {

		for i := v; i < len(current); i += v {
			if current[i-v:i] != current[i:i+v] {
				continue out
			}
		}
		// fmt.Printf("isInvalid code: %v, value: %v\n", code, v)
		return code
	}
	return 0
}

func populateBrakpoint(breakpoint map[int][]int, length int) {
	values := []int{}
	for i := 1; i <= length/2; i++ {
		if length%i == 0 {
			values = append(values, i)
		}
	}
	breakpoint[length] = values

}
