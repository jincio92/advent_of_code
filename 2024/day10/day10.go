package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
)

func main() {
	file, err := os.Open("day10/day10.input")
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(file)
	topo := [][]int{}
	for scanner.Scan() {
		line := scanner.Text()
		intLine := []int{}
		for i := 0; i < len(line); i++ {
			value, err := strconv.Atoi(string(line[i]))
			if err != nil {
				panic(err)
			}
			intLine = append(intLine, value)

		}
		topo = append(topo, intLine)
	}

	for _, v := range topo {
		fmt.Printf("%v\n", v)
	}
	totalScore := 0
	totalRating := 0
	for y, riga := range topo {
		for x := 0; x < len(riga); x++ {
			if riga[x] == 0 {
				fmt.Printf("Trovato 0 a x: %v, y: %v\n", x, y)
				result := [][2]int{}
				result = append(result, [2]int{x, y})
				result = Navigate(&topo, result, 0)
				currentValue := 1
				for currentValue < 9 {
					result = Navigate(&topo, result, currentValue)
					currentValue++
				}
				slices.SortFunc(result, func(a [2]int, b [2]int) int {
					if a[0] == b[0] {
						return a[1] - b[1]
					}
					return a[0] - b[0]
				})
				compatted := slices.Compact(result)
				fmt.Printf("result: %v, nr: %v, rating: %v\n\n", compatted, len(compatted), len(result))
				totalRating += len(result)
				totalScore += len(compatted)

			}
		}
	}
	fmt.Printf("Total: %v, Rating: %v\n", totalScore, totalRating)
}

func Navigate(topo *[][]int, currentPositions [][2]int, currentValue int) [][2]int {
	result := [][2]int{}
	for _, current := range currentPositions {
		if current[1] > 0 && (*topo)[current[1]-1][current[0]] == currentValue+1 {
			result = append(result, [2]int{current[0], current[1] - 1})
		}
		if current[1] < len(*topo)-1 && (*topo)[current[1]+1][current[0]] == currentValue+1 {
			result = append(result, [2]int{current[0], current[1] + 1})
		}
		if current[0] > 0 && (*topo)[current[1]][current[0]-1] == currentValue+1 {
			result = append(result, [2]int{current[0] - 1, current[1]})
		}
		if current[0] < len((*topo)[0])-1 && (*topo)[current[1]][current[0]+1] == currentValue+1 {
			result = append(result, [2]int{current[0] + 1, current[1]})
		}
	}
	return result
}
