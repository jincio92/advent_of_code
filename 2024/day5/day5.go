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

	file, err := os.Open("day5/day5.input")
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)

	rules := [][]string{}
	updates := []string{}
	isrules := true
	for scanner.Scan() {
		if scanner.Text() == "" {
			isrules = false
		} else if isrules {
			rules = append(rules, strings.Split(scanner.Text(), "|"))
		} else {
			updates = append(updates, scanner.Text())
		}
	}

	correctUpdates := [][]string{}
	wrongUpdates := [][]string{}
	for _, up := range updates {
		broken := false
		for _, rul := range rules {
			firstIndex := strings.Index(up, rul[0])
			secondIndex := strings.Index(up, rul[1])
			if firstIndex >= 0 && secondIndex >= 0 && firstIndex > secondIndex {
				// fmt.Printf("broken: %v, rule: %v", up, rul)
				// println()
				broken = true
				wrongUpdates = append(wrongUpdates, strings.Split(up, ","))
				break
			}
		}
		if !broken {
			correctUpdates = append(correctUpdates, strings.Split(up, ","))
		}
	}
	// fmt.Printf("correctUpdates: %v\nwrongUpdates: %v", correctUpdates, wrongUpdates)
	// println()

	sum := 0
	for _, corr := range correctUpdates {
		value, err := strconv.Atoi(corr[len(corr)/2])
		if err != nil {
			panic(err)
		}
		sum += value
	}
	fmt.Printf("correctUpdates sum: %v", sum)
	println()

	for _, wro := range wrongUpdates {
		for reorder(rules, &wro) {
		}

	}
	// fmt.Printf("wrongUpdates: %v", wrongUpdates)
	// println()

	sum = 0
	for _, corr := range wrongUpdates {
		value, err := strconv.Atoi(corr[len(corr)/2])
		if err != nil {
			panic(err)
		}
		sum += value
	}
	fmt.Printf("wrongUpdates sum: %v", sum)
	println()
}

func reorder(rules [][]string, upd *[]string) bool {
	isWrong := false
	for _, rul := range rules {
		firstIndex := slices.Index(*upd, rul[0])
		secondIndex := slices.Index(*upd, rul[1])
		if firstIndex >= 0 && secondIndex >= 0 && firstIndex > secondIndex {
			temp := (*upd)[firstIndex]
			(*upd)[firstIndex] = (*upd)[secondIndex]
			(*upd)[secondIndex] = temp
			isWrong = true
		}
	}
	return isWrong
}
