package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {

	file, err := os.Open("day11/day11.input")
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(file)
	scanner.Scan()
	line := scanner.Text()
	split := strings.Split(line, " ")
	group := []int{}
	for _, piece := range split {
		value, err := strconv.Atoi(piece)
		if err != nil {
			panic(err)
		}
		group = append(group, value)
	}
	sum := 0
	cacheRes := map[string]int{}
	for i, stone := range group {
		start := time.Now()
		sum += CalculateBlink([]int{stone}, 75, &cacheRes)
		fmt.Printf("Stone %v, time: %v, cached: %v\n", i+1, time.Since(start), len(cacheRes))
	}
	fmt.Printf("TotalStone: %v\n", sum)
}

func CalculateBlink(list []int, blink int, cacheRes *map[string]int) int {
	if blink == 0 {
		return len(list)
	}
	if val, ok := (*cacheRes)[fmt.Sprintf("%v%v", list, blink)]; ok {
		return val
	}
	// fmt.Printf("\rCurrent Blink: %v", blink)

	sum := 0
	for _, v := range list {
		var arr []int
		if v == 0 {
			arr = []int{1}
		} else {
			numLength := int(math.Ceil(math.Log10(float64(v + 1))))
			if numLength%2 == 0 {
				arr = []int{v / int(math.Pow10(numLength/2)), v % int(math.Pow10(numLength/2))}
			} else {
				arr = []int{v * 2024}
			}
		}
		value := CalculateBlink(arr, blink-1, cacheRes)
		(*cacheRes)[fmt.Sprintf("%v%v", arr, blink-1)] = value
		sum += value
	}
	return sum
}
