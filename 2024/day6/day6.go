package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strings"
)

func main() {

	file, err := os.Open("day6/day6.input")
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)
	guardMap := [][]string{}
	for scanner.Scan() {
		byteArr := scanner.Bytes()
		stringArr := []string{}
		for _, v := range byteArr {
			stringArr = append(stringArr, string(v))
		}
		guardMap = append(guardMap, stringArr)
	}
	loopMap := Clone(guardMap)
	position, incrementer := InitGuard(guardMap)

	println()
	for CanNavigate(&guardMap, position, incrementer) {
		loop := Navigate(&guardMap, &position, &incrementer)
		if loop {
			println("Siamo in un loop")
			break
		}
	}

	sum := 1
	for _, line := range guardMap {
		sum += strings.Count(strings.Join(line, ""), "-")
		sum += strings.Count(strings.Join(line, ""), "|")
		sum += strings.Count(strings.Join(line, ""), "+")
	}
	fmt.Printf("totale: %v\n", sum)

	position, incrementer = InitGuard(loopMap)
	fmt.Printf("position %v, incrementer: %v\n", position, incrementer)

	loopCount := 0
	for i := 0; i < len(loopMap); i++ {
		for j := 0; j < len(loopMap[0]); j++ {
			if loopMap[i][j] == "#" {
				continue
			}
			if guardMap[i][j] == "." {
				continue
			}
			if loopMap[i][j] == "^" {
				println("found guard")
				continue
			}
			tempMap := Clone(loopMap)
			tempMap[i][j] = "O"
			tempPosition := [2]int{position[0], position[1]}
			tempIncrementer := [2]int{incrementer[0], incrementer[1]}
			for CanNavigate(&tempMap, tempPosition, tempIncrementer) {
				loop := Navigate(&tempMap, &tempPosition, &tempIncrementer)
				if loop {
					loopCount++
					break
				}
			}
		}
	}
	fmt.Printf("totale Loop: %v\n", loopCount)
}

func Navigate(guardMap *[][]string, position *[2]int, incrementer *[2]int) bool {
	newPostion := [2]int{position[0] + incrementer[0], position[1] + incrementer[1]}

	if (*guardMap)[newPostion[1]][newPostion[0]] == "#" || (*guardMap)[newPostion[1]][newPostion[0]] == "O" {
		if (*guardMap)[(*position)[1]][(*position)[0]] == "+" {
			return true
		}
		Rotate(incrementer)
		newPostion = [2]int{position[0] + incrementer[0], position[1] + incrementer[1]}
		if (*guardMap)[newPostion[1]][newPostion[0]] == "#" || (*guardMap)[newPostion[1]][newPostion[0]] == "O" {
			Rotate(incrementer)
			newPostion = [2]int{position[0] + incrementer[0], position[1] + incrementer[1]}
		}

		(*guardMap)[(*position)[1]][(*position)[0]] = "+"
		(*position) = newPostion
	} else {

		if (*guardMap)[(*position)[1]][(*position)[0]] != "+" {
			if (*incrementer)[0] != 0 {
				(*guardMap)[(*position)[1]][(*position)[0]] = "-"
			} else {
				(*guardMap)[(*position)[1]][(*position)[0]] = "|"
			}
		}
		(*position) = newPostion
	}

	return false
}

func Rotate(incrementer *[2]int) {
	var newGuard string
	if (*incrementer)[0] == 0 && (*incrementer)[1] == -1 {
		newGuard = ">"
	} else if (*incrementer)[0] == 1 && (*incrementer)[1] == 0 {
		newGuard = "v"
	} else if (*incrementer)[0] == -1 && (*incrementer)[1] == 0 {
		newGuard = "^"
	} else if (*incrementer)[0] == 0 && (*incrementer)[1] == 1 {
		newGuard = "<"
	}
	(*incrementer) = GetIncrementer(newGuard)
}

func InitGuard(guardMap [][]string) (position [2]int, incrementer [2]int) {
	position = [2]int{}
	incrementer = [2]int{0, 0}
	for i, line := range guardMap {
		// fmt.Printf("%v\n", line)
		if guard := HasGuard(line); guard != "" {
			position = [2]int{slices.Index(line, guard), i}

			incrementer = GetIncrementer(guard)
		}
	}
	return position, incrementer
}

func GetIncrementer(guard string) [2]int {
	incrementer := [2]int{0, 0}
	if guard == "v" {
		incrementer = [2]int{0, 1}
	}
	if guard == ">" {
		incrementer = [2]int{1, 0}
	}
	if guard == "<" {
		incrementer = [2]int{-1, 0}
	}
	if guard == "^" {
		incrementer = [2]int{0, -1}
	}
	return incrementer
}

func CanNavigate(guardMap *[][]string, position [2]int, incrementer [2]int) bool {
	newPostion := [2]int{position[0] + incrementer[0], position[1] + incrementer[1]}
	if newPostion[0] < 0 || newPostion[0] >= len((*guardMap)[0]) || newPostion[1] < 0 || newPostion[1] >= len((*guardMap)) {
		(*guardMap)[position[1]][position[0]] = "X"
		return false
	}
	return true
}

func HasGuard(line []string) string {

	if slices.Contains(line, "^") {
		return "^"
	}
	if slices.Contains(line, ">") {
		return ">"
	}
	if slices.Contains(line, "<") {
		return "<"
	}
	if slices.Contains(line, "v") {
		return "v"
	}
	return ""
}

func Clone(arr [][]string) [][]string {
	newArr := [][]string{}
	for _, line := range arr {
		newLine := []string{}
		newLine = append(newLine, line...)
		newArr = append(newArr, newLine)
	}
	return newArr
}
