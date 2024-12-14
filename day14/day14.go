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

	file, err := os.Open("day14/day14.input")
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(file)
	args := os.Args[1:]
	seconds, err := strconv.Atoi(args[0])
	if err != nil {
		panic(err)
	}
	areaWidth := 101
	areaHeight := 103
	// areaWidth := 11
	// areaHeight := 7
	robotFinals := [][2]int{}
	for scanner.Scan() {
		line := scanner.Text()
		startX, err := strconv.Atoi(line[strings.Index(line, "p=")+2 : strings.Index(line, ",")])
		if err != nil {
			panic(err)
		}
		startY, err := strconv.Atoi(line[strings.Index(line, ",")+1 : strings.Index(line, " ")])
		if err != nil {
			panic(err)
		}
		deltaX, err := strconv.Atoi(line[len(line)/2:][strings.Index(line[len(line)/2:], "v=")+2 : strings.Index(line[len(line)/2:], ",")])
		if err != nil {
			panic(err)
		}
		deltaY, err := strconv.Atoi(line[len(line)/2:][strings.Index(line[len(line)/2:], ",")+1:])
		if err != nil {
			panic(err)
		}

		totalX := (startX + (deltaX * seconds))
		totalY := (startY + (deltaY * seconds))

		finalX := totalX
		finalY := totalY

		if totalX/areaWidth != 0 {
			finalX = totalX % ((totalX / areaWidth) * areaWidth)

		}
		if totalY/areaHeight != 0 {
			finalY = totalY % ((totalY / areaHeight) * areaHeight)
		}
		if finalX < 0 {
			finalX += areaWidth
		}
		if finalY < 0 {
			finalY += areaHeight
		}
		robotFinals = append(robotFinals, [2]int{finalX, finalY})
	}

	quad1 := 0
	quad2 := 0
	quad3 := 0
	quad4 := 0
	for _, r := range robotFinals {
		if r[0] < areaWidth/2 && r[1] < areaHeight/2 {
			quad1++
		}
		if r[0] > areaWidth/2 && r[1] < areaHeight/2 {
			quad2++
		}
		if r[0] < areaWidth/2 && r[1] > areaHeight/2 {
			quad3++
		}
		if r[0] > areaWidth/2 && r[1] > areaHeight/2 {
			quad4++
		}
		// fmt.Printf("%v\n", r)
	}
	fmt.Printf("Total: %v, 1: %v, 2: %v, 3: %v, 4: %v\n", quad1*quad2*quad3*quad4, quad1, quad2, quad3, quad4)

	compact := slices.Compact(robotFinals)
	mappa := []string{}
	for i := 0; i < areaHeight; i++ {
		// mappa = append(mappa, "...........")
		mappa = append(mappa, ".....................................................................................................")
	}
	for _, r := range compact {
		line := mappa[r[1]]
		mappa[r[1]] = line[:r[0]] + "O" + line[r[0]+1:]
	}
	for _, v := range mappa {
		fmt.Printf("%v\n", v)
	}

	// christmas tree at 7569
}
