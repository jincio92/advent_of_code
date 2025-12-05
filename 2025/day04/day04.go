package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	file, err := os.Open("2025/day04/day04.input")
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

	lines := []string{}
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	total := 0
	run := 1
	for run != 0 {
		run = 0
		for row := 0; row < len(lines); row++ {
			for col := 0; col < len(lines[row]); col++ {
				if string(lines[row][col]) == "." {
					continue
				}
				surround := ""
				if col > 0 {
					if row > 0 {
						surround += string(lines[row-1][col-1])
					}
					surround += string(lines[row][col-1])
					if row < len(lines)-1 {
						surround += string(lines[row+1][col-1])
					}
				}
				if row > 0 {
					surround += string(lines[row-1][col])
				}

				if row < len(lines)-1 {
					surround += string(lines[row+1][col])
				}
				if col < len(lines[row])-1 {
					if row > 0 {
						surround += string(lines[row-1][col+1])
					}
					surround += string(lines[row][col+1])
					if row < len(lines)-1 {
						surround += string(lines[row+1][col+1])
					}
				}
				if strings.Count(surround, "@") < 4 {
					// fmt.Printf("fount at row: %v, col: %v\n", row, col)
					run++
					total++
					lines[row] = lines[row][:col] + "." + lines[row][col+1:]
				}
			}
		}
	}
	fmt.Printf("partOne: %v\n", total)
}
