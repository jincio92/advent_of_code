package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"time"
)

func main() {

	file, err := os.Open("day12/day12.input")
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(file)
	mappa := []string{}
	for scanner.Scan() {
		mappa = append(mappa, scanner.Text())
	}
	perimeterMap := map[string]int{}
	areaMap := map[string]int{}
	sideMap := map[string]int{}
	start := time.Now()
	visited := [][2]int{}
	for y, riga := range mappa {
		for x := 0; x < len(riga); x++ {
			letter := string(mappa[y][x])
			if slices.Contains(visited, [2]int{x, y}) || (x > 0 && riga[x] == riga[x-1]) {
				continue
			}
			count, peri, side := FindNear(&mappa, [2]int{x, y}, &visited)
			if count > 0 {
				areaMap[letter+strconv.Itoa(x)+strconv.Itoa(y)] = count
				perimeterMap[letter+strconv.Itoa(x)+strconv.Itoa(y)] = peri
				sideMap[letter+strconv.Itoa(x)+strconv.Itoa(y)] = side
			}
		}
	}

	sumPeri := 0
	sumSide := 0
	for key := range areaMap {
		sumPeri += areaMap[key] * perimeterMap[key]
		sumSide += areaMap[key] * sideMap[key]
	}

	fmt.Printf("TotalPerimeter: %v, TotalSide: %v, tempo: %v\n", sumPeri, sumSide, time.Since(start))
}

func FindNear(mappa *[]string, position [2]int, visited *[][2]int) (int, int, int) {

	letter := (*mappa)[position[1]][position[0]]

	if slices.Contains((*visited), position) {
		return 0, 0, 0
	}
	(*visited) = append((*visited), position)
	near := 0
	side := 0
	top := false
	bot := false
	right := false
	left := false
	nearPositions := [][2]int{}
	if position[0] > 0 && (*mappa)[position[1]][position[0]-1] == letter {
		near++
		nearPositions = append(nearPositions, [2]int{position[0] - 1, position[1]})
		left = true
	}
	if position[0] < len((*mappa)[0])-1 && (*mappa)[position[1]][position[0]+1] == letter {
		near++
		nearPositions = append(nearPositions, [2]int{position[0] + 1, position[1]})
		right = true
	}
	if position[1] > 0 && (*mappa)[position[1]-1][position[0]] == letter {
		near++
		nearPositions = append(nearPositions, [2]int{position[0], position[1] - 1})
		top = true
	}
	if position[1] < len(*mappa)-1 && (*mappa)[position[1]+1][position[0]] == letter {
		near++
		nearPositions = append(nearPositions, [2]int{position[0], position[1] + 1})
		bot = true
	}
	//convess corner
	if !top && !right {
		side++
	}
	if !top && !left {
		side++
	}
	if !bot && !left {
		side++
	}
	if !bot && !right {
		side++
	}

	//concave corner
	//top-right
	if top && right && (*mappa)[position[1]-1][position[0]+1] != letter {
		side++
	}
	//top-left
	if top && left && (*mappa)[position[1]-1][position[0]-1] != letter {
		side++
	}
	//bot-right
	if bot && right && (*mappa)[position[1]+1][position[0]+1] != letter {
		side++
	}
	//bot-left
	if bot && left && (*mappa)[position[1]+1][position[0]-1] != letter {
		side++
	}

	newCount, newNear := 0, 0
	for _, pos := range nearPositions {
		tempCount, tempNear, tempSide := FindNear(mappa, pos, visited)
		newCount += tempCount
		newNear += tempNear
		side += tempSide
	}
	return newCount + 1, newNear + 4 - near, side
}
